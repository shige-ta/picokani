#!/usr/bin/env python3
"""
Discord Task Status Script
Discord Activity APIを利用してステータスメッセージを設定・更新します
"""

import argparse
import json
import os
import sys
from typing import Optional

try:
    import discord
    from discord.ext import commands
except ImportError:
    print("Error: discord.py is not installed.")
    print("Install with: pip install discord.py")
    sys.exit(1)


class DiscordStatus(commands.Bot):
    """Discordステータス管理ボット"""

    def __init__(self):
        super().__init__(
            intents=discord.Intents.default(),
            command_prefix="!",
        )
        self.config = self._load_config()
        self._setup_hook()

    def _load_config(self) -> dict:
        """設定ファイルを読み込む"""
        config_path = os.path.expanduser("~/.picoclaw/config.json")

        if not os.path.exists(config_path):
            print(f"Error: Config file not found at {config_path}")
            print("Please add your Discord configuration to config.json")
            sys.exit(1)

        try:
            with open(config_path, 'r') as f:
                return json.load(f)
        except json.JSONDecodeError as e:
            print(f"Error: Invalid JSON in config file: {e}")
            sys.exit(1)

    def _setup_hook(self):
        """ボットのセットアップ"""
        if 'discord' not in self.config:
            print("Error: No Discord configuration found in config.json")
            print("Please add 'discord' section with 'token' and 'client_id'")
            sys.exit(1)

        token = self.config['discord'].get('token')
        client_id = self.config['discord'].get('client_id')

        if not token:
            print("Error: Discord token not found in config")
            sys.exit(1)

        if not client_id:
            print("Warning: Client ID not found in config")
            print("Status updates may not work correctly")

        self.token = token
        self.client_id = client_id

    async def setup_hook(self):
        """ボットのセットアップフック"""
        pass

    async def on_ready(self):
        """ボットが準備完了になったとき"""
        print(f"Logged in as {self.user.name} (ID: {self.user.id})")
        print(f"Ready to update status")

        # アクティビティを設定
        if self.activity:
            print(f"Current activity: {self.activity.name} ({self.activity.type})")

    async def update_status(self, message: str, activity_type: str = "PLAYING", url: Optional[str] = None):
        """ステータスを更新する"""
        try:
            # アクティビティタイプを変換
            activity_type_map = {
                "PLAYING": discord.ActivityType.playing,
                "LISTENING": discord.ActivityType.listening,
                "WATCHING": discord.ActivityType.watching,
                "STREAMING": discord.ActivityType.streaming,
                "COMPETING": discord.ActivityType.competing,
                "IDLE": discord.ActivityType.idle,
                "DND": discord.ActivityType.do_not_disturb,
            }

            activity_type_enum = activity_type_map.get(activity_type.upper(), discord.ActivityType.playing)

            # アクティビティを作成
            if activity_type_enum == discord.ActivityType.streaming and not url:
                print("Error: URL is required for STREAMING activity type")
                return False

            activity = discord.Activity(
                name=message,
                type=activity_type_enum,
                url=url if url else None
            )

            # ステータスを更新
            await self.change_presence(activity=activity)
            print(f"✓ Status updated: {message} ({activity_type})")
            return True

        except Exception as e:
            print(f"Error: Failed to update status: {e}")
            return False

    async def clear_status(self):
        """ステータスをクリアする"""
        try:
            await self.change_presence(activity=None)
            print("✓ Status cleared")
            return True
        except Exception as e:
            print(f"Error: Failed to clear status: {e}")
            return False

    async def show_status(self):
        """現在のステータスを表示する"""
        try:
            print("\nCurrent Status:")
            if self.user:
                if self.activity:
                    print(f"  Activity: {self.activity.name}")
                    print(f"  Type: {self.activity.type}")
                    if self.activity.url:
                        print(f"  URL: {self.activity.url}")
                else:
                    print("  No activity set")
            print(f"  User: {self.user.name} (ID: {self.user.id})")
            print()
            return True
        except Exception as e:
            print(f"Error: Failed to show status: {e}")
            return False


def main():
    """メイン関数"""
    parser = argparse.ArgumentParser(
        description="Discord Task Status - Discordのステータスメッセージを設定・更新します",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
使用例:
  %(prog)s set "作業中" --type PLAYING
  %(prog)s set "データ分析中... 50%" --type PLAYING
  %(prog)s set "音楽を聴いています" --type LISTENING
  %(prog)s set "動画を見ています" --type WATCHING
  %(prog)s set "ライブ配信" --type STREAMING --url "https://example.com/stream"
  %(prog)s show
  %(prog)s clear
        """
    )

    subparsers = parser.add_subparsers(dest="command", help="コマンド")

    # setコマンド
    set_parser = subparsers.add_parser("set", help="ステータスメッセージを設定")
    set_parser.add_argument("message", help="ステータスメッセージ（最大128文字）")
    set_parser.add_argument("--type", choices=["PLAYING", "LISTENING", "WATCHING", "STREAMING", "COMPETING", "IDLE", "DND"],
                          default="PLAYING", help="アクティビティタイプ")
    set_parser.add_argument("--url", help="STREAMING用のURL")

    # showコマンド
    subparsers.add_parser("show", help="現在のステータスを表示")

    # clearコマンド
    subparsers.add_parser("clear", help="ステータスをクリア")

    args = parser.parse_args()

    if not args.command:
        parser.print_help()
        sys.exit(1)

    # ボットを作成
    bot = DiscordStatus()

    # コマンドを実行
    if args.command == "set":
        success = bot.loop.run_until_complete(
            bot.update_status(args.message, args.type, args.url)
        )
        sys.exit(0 if success else 1)

    elif args.command == "show":
        success = bot.loop.run_until_complete(bot.show_status())
        sys.exit(0 if success else 1)

    elif args.command == "clear":
        success = bot.loop.run_until_complete(bot.clear_status())
        sys.exit(0 if success else 1)


if __name__ == "__main__":
    main()
