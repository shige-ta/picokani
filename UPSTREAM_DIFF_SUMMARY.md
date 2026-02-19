## upstream 比較メモ

比較対象:

- 公式: https://github.com/sipeed/picoclaw (clone時点の最新)
- 公開用: このディレクトリ

### 公式にない独自追加（この公開版に含めたもの）

- `extras/discord-task-status/README.md`
- `extras/discord-task-status/scripts/picoclaw-discord-status.py`
  - Discord Bot のプレゼンスを動的更新する拡張スクリプト
  - 私有実行環境の skill 実装から取り込み

### 公式にはあるが、この公開版に未反映の主な機能群

- ルーティング関連:
  - `pkg/routing/*`
- 新しい provider 分割/共通化関連:
  - `pkg/providers/anthropic/*`
  - `pkg/providers/openai_compat/*`
  - `pkg/providers/protocoltypes/*`
  - `pkg/providers/factory.go`
  - `pkg/providers/fallback.go`
  - `pkg/providers/cooldown.go`
  - `pkg/providers/error_classifier.go`
  - `pkg/providers/model_ref.go`
- agent registry/instance 関連:
  - `pkg/agent/instance.go`
  - `pkg/agent/registry.go`
- その他:
  - `.golangci.yaml`
  - `docs/tools_configuration.md`
  - `README.pt-br.md`
  - `README.vi.md`

### 公開向けに意図的に除外したもの

- `build/`
- `gateway.log`
- `config/config.json`
- `SOUL.md`, `AGENT.md`, `USER.md`, `IDENTITY.md`, `SKILL.md`
