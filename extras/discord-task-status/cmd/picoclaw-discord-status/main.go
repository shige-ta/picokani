package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type discordConfig struct {
	Token    string `json:"token"`
	ClientID string `json:"client_id"`
}

type configFile struct {
	Discord  discordConfig `json:"discord"`
	Channels struct {
		Discord discordConfig `json:"discord"`
	} `json:"channels"`
}

type cachedStatus struct {
	Message      string `json:"message"`
	ActivityType string `json:"activity_type"`
	URL          string `json:"url,omitempty"`
	Status       string `json:"status"`
	UpdatedAt    string `json:"updated_at"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "set":
		if err := runSet(cfg, os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "show":
		if err := runShow(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "clear":
		if err := runClear(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Discord Task Status (Go)")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  picoclaw-discord-status set \"message\" [--type PLAYING|LISTENING|WATCHING|STREAMING|COMPETING|IDLE|DND] [--url URL]")
	fmt.Println("  picoclaw-discord-status set --phase build --action \"indexing\" --file pkg/tools/web.go --dir pkg/tools --progress 45")
	fmt.Println("  picoclaw-discord-status show")
	fmt.Println("  picoclaw-discord-status clear")
}

func runSet(cfg discordConfig, args []string) error {
	fs := flag.NewFlagSet("set", flag.ContinueOnError)
	kind := fs.String("type", "PLAYING", "activity type")
	url := fs.String("url", "", "streaming url")
	phase := fs.String("phase", "", "phase label (e.g. build/train/infer)")
	action := fs.String("action", "", "action label (e.g. indexing files)")
	file := fs.String("file", "", "file path to show")
	dir := fs.String("dir", "", "directory path to show")
	progress := fs.Int("progress", -1, "progress percent (0-100)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	msg := ""
	if fs.NArg() >= 1 {
		msg = fs.Arg(0)
	} else {
		msg = composeDetailMessage(*phase, *action, *file, *dir, *progress)
	}
	if msg == "" {
		return errors.New("message is required (or provide --action/--file/--dir/--phase)")
	}
	msg = clampMessage(msg, 128)
	if *progress > 100 || *progress < -1 {
		return errors.New("progress must be between 0 and 100")
	}

	activity, status, err := parseActivity(msg, *kind, *url)
	if err != nil {
		return err
	}

	s, err := newSession(cfg.Token)
	if err != nil {
		return err
	}
	defer s.Close()

	if err := s.Open(); err != nil {
		return fmt.Errorf("failed to open discord gateway: %w", err)
	}
	if err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status:     status,
		Activities: []*discordgo.Activity{activity},
	}); err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	_ = writeCache(cachedStatus{
		Message:      msg,
		ActivityType: strings.ToUpper(*kind),
		URL:          *url,
		Status:       status,
		UpdatedAt:    time.Now().Format(time.RFC3339),
	})
	fmt.Printf("OK: status updated (%s)\n", strings.ToUpper(*kind))
	return nil
}

func runClear(cfg discordConfig) error {
	s, err := newSession(cfg.Token)
	if err != nil {
		return err
	}
	defer s.Close()

	if err := s.Open(); err != nil {
		return fmt.Errorf("failed to open discord gateway: %w", err)
	}
	if err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status:     "online",
		Activities: []*discordgo.Activity{},
	}); err != nil {
		return fmt.Errorf("failed to clear status: %w", err)
	}

	_ = writeCache(cachedStatus{
		Message:      "",
		ActivityType: "NONE",
		Status:       "online",
		UpdatedAt:    time.Now().Format(time.RFC3339),
	})
	fmt.Println("OK: status cleared")
	return nil
}

func runShow(cfg discordConfig) error {
	s, err := newSession(cfg.Token)
	if err != nil {
		return err
	}
	defer s.Close()

	u, err := s.User("@me")
	if err != nil {
		return fmt.Errorf("failed to fetch bot user: %w", err)
	}
	fmt.Printf("Bot: %s (ID: %s)\n", u.Username, u.ID)

	c, err := readCache()
	if err != nil {
		fmt.Println("Status: unknown (no local cache)")
		return nil
	}
	if c.ActivityType == "NONE" || c.Message == "" {
		fmt.Println("Status: cleared")
		return nil
	}
	fmt.Printf("Status: %s\n", c.Status)
	fmt.Printf("Activity: %s - %s\n", c.ActivityType, c.Message)
	if c.URL != "" {
		fmt.Printf("URL: %s\n", c.URL)
	}
	fmt.Printf("UpdatedAt: %s\n", c.UpdatedAt)
	return nil
}

func parseActivity(message, kind, url string) (*discordgo.Activity, string, error) {
	k := strings.ToUpper(kind)
	status := "online"
	actType := discordgo.ActivityTypeGame

	switch k {
	case "PLAYING":
		actType = discordgo.ActivityTypeGame
	case "STREAMING":
		actType = discordgo.ActivityTypeStreaming
		if url == "" {
			return nil, "", errors.New("url is required for STREAMING")
		}
	case "LISTENING":
		actType = discordgo.ActivityTypeListening
	case "WATCHING":
		actType = discordgo.ActivityTypeWatching
	case "COMPETING":
		actType = discordgo.ActivityTypeCompeting
	case "IDLE":
		actType = discordgo.ActivityTypeGame
		status = "idle"
	case "DND":
		actType = discordgo.ActivityTypeGame
		status = "dnd"
	default:
		return nil, "", fmt.Errorf("unsupported activity type: %s", kind)
	}

	return &discordgo.Activity{
		Name: message,
		Type: actType,
		URL:  url,
	}, status, nil
}

func composeDetailMessage(phase, action, file, dir string, progress int) string {
	parts := make([]string, 0, 5)
	if phase != "" {
		parts = append(parts, phase)
	}
	if action != "" {
		parts = append(parts, action)
	}
	if file != "" {
		parts = append(parts, "file:"+shortenPath(normalizePath(file), 44))
	}
	if dir != "" {
		parts = append(parts, "dir:"+shortenPath(normalizePath(dir), 34))
	}
	if progress >= 0 {
		parts = append(parts, fmt.Sprintf("%d%%", progress))
	}
	return strings.Join(parts, " | ")
}

func normalizePath(p string) string {
	return strings.ReplaceAll(filepath.Clean(p), "\\", "/")
}

func clampMessage(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	if max <= 3 {
		return string(r[:max])
	}
	return string(r[:max-3]) + "..."
}

func shortenPath(p string, max int) string {
	r := []rune(p)
	if len(r) <= max {
		return p
	}
	if max < 7 {
		return string(r[:max])
	}
	head := (max - 3) / 2
	tail := max - 3 - head
	return string(r[:head]) + "..." + string(r[len(r)-tail:])
}

func newSession(token string) (*discordgo.Session, error) {
	if token == "" {
		return nil, errors.New("discord token not found in config")
	}
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("failed to create discord session: %w", err)
	}
	s.Identify.Intents = 0
	return s, nil
}

func loadConfig() (discordConfig, error) {
	path, err := userConfigPath()
	if err != nil {
		return discordConfig{}, err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return discordConfig{}, fmt.Errorf("failed to read config %s: %w", path, err)
	}
	var cfg configFile
	if err := json.Unmarshal(b, &cfg); err != nil {
		return discordConfig{}, fmt.Errorf("invalid JSON in %s: %w", path, err)
	}
	if cfg.Discord.Token != "" {
		return cfg.Discord, nil
	}
	if cfg.Channels.Discord.Token != "" {
		return cfg.Channels.Discord, nil
	}
	return discordConfig{}, errors.New("discord token not found (supported keys: discord.token or channels.discord.token)")
}

func userConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".picoclaw", "config.json"), nil
}

func cachePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".picoclaw", "state", "discord-task-status.json"), nil
}

func writeCache(s cachedStatus) error {
	p, err := cachePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0o644)
}

func readCache() (cachedStatus, error) {
	p, err := cachePath()
	if err != nil {
		return cachedStatus{}, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return cachedStatus{}, err
	}
	var s cachedStatus
	if err := json.Unmarshal(b, &s); err != nil {
		return cachedStatus{}, err
	}
	return s, nil
}
