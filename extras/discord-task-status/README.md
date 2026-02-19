# Discord Task Status Skill

Discord Activity APIを利用して、ボットのステータスメッセージを自動更新するスキルです。

## 機能

- Discordステータスメッセージの設定・更新
- タスク進行中のステータス表示
- アクティビティタイプの指定（PLAYING, LISTENING, WATCHING, IDLEなど）
- 現在のステータス表示

## 使用方法

### ステータスメッセージを設定

```bash
# 基本的なステータス
picoclaw-discord-status.py set "Working on task..."

# アクティビティタイプを指定
picoclaw-discord-status.py set "Processing data" --type PLAYING
picoclaw-discord-status.py set "Idle" --type IDLE
picoclaw-discord-status.py set "Listening to music" --type LISTENING
picoclaw-discord-status.py set "Watching videos" --type WATCHING
```

### 現在のステータスを表示

```bash
picoclaw-discord-status.py show
```

## 設定

`~/.picoclaw/config.json`に以下の設定を追加する必要があります：

```json
{
  "discord": {
    "token": "your-discord-bot-token",
    "client_id": "your-client-id"
  }
}
```

## タスク実行時の使用例

```bash
# タスク開始
picoclaw-discord-status.py set "Starting task..."

# 処理中（10〜30秒ごとに更新）
picoclaw-discord-status.py set "Processing... 25%"
picoclaw-discord-status.py set "Processing... 50%"
picoclaw-discord-status.py set "Processing... 75%"

# 完了
picoclaw-discord-status.py set "Task completed!"
```

## 依存関係

- Python 3.6+
- discord.py

インストール：
```bash
pip install discord.py
```
