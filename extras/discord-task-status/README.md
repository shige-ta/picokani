# Discord Task Status Skill (Go)

Discord Bot のプレゼンスを動的に更新する Go 製ツールです。

## 実装

- ソース: `cmd/picoclaw-discord-status/main.go`
- コマンド: `picoclaw-discord-status`
- 対応コマンド: `set` / `show` / `clear`

## ビルド

```bash
# リポジトリルートで実行
GOOS=linux GOARCH=riscv64 go build -o extras/discord-task-status/bin/picoclaw-discord-status ./extras/discord-task-status/cmd/picoclaw-discord-status
```

## 使用方法

```bash
# ステータス設定
picoclaw-discord-status set "Working on task..." --type PLAYING

# 詳細情報付き（処理内容 / ファイル / ディレクトリ / 進捗）
picoclaw-discord-status set --phase build --action "indexing" --file pkg/tools/web.go --dir pkg/tools --progress 45 --type PLAYING

# 進捗更新
picoclaw-discord-status set "Processing... 50%" --type PLAYING

# 表示（ローカルキャッシュ）
picoclaw-discord-status show

# クリア
picoclaw-discord-status clear
```

## 設定ファイル

`~/.picoclaw/config.json` の以下のどちらかを参照します。

- `discord.token`
- `channels.discord.token`

## 補足

- `show` は直近に設定した内容をローカル状態ファイルから表示します。
- 状態ファイル: `~/.picoclaw/state/discord-task-status.json`
