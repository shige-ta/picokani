## 蜈ｬ蠑 Fork 讎りｦ・
蝓ｺ莠守ｧ∵怏霑占｡檎識蠅・ｸｭ螳樣刔霑占｡檎噪莉｣遐∬ｿ幄｡悟・譫撰ｼ悟ｹｶ謖芽ｯ･扈捺棡蜷梧ｭ･縲・
荳主ｮ俶婿莉灘ｺ鍋噪蟾ｮ蠑・

- 螳俶婿: https://github.com/sipeed/picoclaw
- 蟾ｲ蜷梧ｭ･霑占｡檎識蠅・ｸｭ逧・怙譁ｰ譫・ｻｺ蜑・Go 貅千・- 蟾ｲ蛛壼・蠑蜿大ｸ・ｸ・炊: 蛻髯､ build/縲“ateway.log縲…onfig/config.json
- 蟾ｲ莉主・蠑逶ｮ蠖慕ｧｻ髯､ Agent/Skill 逶ｸ蜈ｳ譁・ｻｶ: SOUL.md縲、GENT.md縲ゞSER.md縲！DENTITY.md縲ヾKILL.md
- 蟾ｲ蝨ｨ Sipeed LicheeRV Nano Wi-Fi 荳企ｪ瑚ｯ∝庄霑占｡・- 蟾ｲ譁ｰ蠅槫ｮ俶婿莉灘ｺ謎ｸｭ豐｡譛臥噪譛ｬ蝨ｰ蜉溯・: 逕ｨ莠・Discord Bot Presence 蜉ｨ諤∵峩譁ｰ逧・discord-task-status
- 隸･蜉溯・莉｣遐∽ｽ咲ｽｮ: extras/discord-task-status/cmd/picoclaw-discord-status/main.go
- 荳主ｮ俶婿蟾ｮ蠑よｱ・ｻ・亥性譛ｬ蠢ｫ辣ｧ譛ｪ蛹・性逧・ｮ俶婿讓｡蝮暦ｼ・ UPSTREAM_DIFF_SUMMARY.md

### Discord Bot 迥ｶ諤∵仞遉ｺ遉ｺ萓・
![Discord bot status](assets/discord-bot-status.png)
### Discord Token 申请与 Presence 设置

1. 在 Discord Developer Portal 创建应用。
2. 在 `Bot` 页面生成（或重置）Bot Token 并保存。
3. 在 `OAuth2 > URL Generator` 勾选 `bot` scope，把 Bot 邀请进服务器。
4. 按需在 `Bot > Privileged Gateway Intents` 打开 `MESSAGE CONTENT INTENT`。

`~/.picoclaw/config.json` 建议如下：

```json
{
  "channels": {
    "discord": {
      "enabled": true,
      "token": "YOUR_DISCORD_BOT_TOKEN",
      "allow_from": [],
      "allow_channels": [],
      "mention_only": false
    }
  }
}
```

动态更新 Bot 状态（Presence）示例：

```bash
picoclaw-discord-status set --phase build --action "indexing" --file "pkg/tools/web.go" --dir "pkg/tools" --progress 45 --type PLAYING
```

如果只使用独立状态工具，也支持以下兼容键：

```json
{
  "discord": {
    "token": "YOUR_DISCORD_BOT_TOKEN",
    "client_id": "YOUR_APPLICATION_ID"
  }
}
```
<div align="center">
<img src="assets/logo.jpg" alt="PicoClaw" width="512">

<h1>PicoClaw: 蝓ｺ莠雑o隸ｭ險逧・ｶ・ｫ俶譜 AI 蜉ｩ謇・/h1>

<h3>10$遑ｬ莉ｶ ﾂｷ 10MB蜀・ｭ・ﾂｷ 1遘貞星蜉ｨ ﾂｷ 逧ｮ逧ｮ陌ｾ・梧・莉ｬ襍ｰ・・/h3>

  <p>
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white" alt="Go">
    <img src="https://img.shields.io/badge/Arch-x86__64%2C%20ARM64%2C%20RISC--V-blue" alt="Hardware">
    <img src="https://img.shields.io/badge/license-MIT-green" alt="License">
    <br>
    <a href="https://picoclaw.io"><img src="https://img.shields.io/badge/Website-picoclaw.io-blue?style=flat&logo=google-chrome&logoColor=white" alt="Website"></a>
    <a href="https://x.com/SipeedIO"><img src="https://img.shields.io/badge/X_(Twitter)-SipeedIO-black?style=flat&logo=x&logoColor=white" alt="Twitter"></a>
  </p>

 **荳ｭ譁・* | [譌･譛ｬ隱枉(README.ja.md) | [English](README.md)
</div>

---

ｦ・**PicoClaw** 譏ｯ荳荳ｪ蜿・[nanobot](https://github.com/HKUDS/nanobot) 蜷ｯ蜿醍噪雜・ｽｻ驥冗ｺｧ荳ｪ莠ｺ AI 蜉ｩ謇九ょｮ・㊦逕ｨ **Go 隸ｭ險** 莉朱峺驥肴桷・檎ｻ丞紙莠・ｸ荳ｪ窶懆・荳ｾ窶晁ｿ・ｨ銀披泌叉逕ｱ AI Agent 閾ｪ霄ｫ鬩ｱ蜉ｨ莠・紛荳ｪ譫ｶ譫・ｿ∫ｧｻ蜥御ｻ｣遐∽ｼ伜喧縲・
笞｡・・**譫∬・霓ｻ驥・*・壼庄蝨ｨ **10 鄒主・** 逧・｡ｬ莉ｶ荳願ｿ占｡鯉ｼ悟・蟄伜頃逕ｨ **<10MB**縲りｿ呎э蜻ｳ逹豈・OpenClaw 闃ら怐 99% 逧・・蟄假ｼ梧ｯ・Mac mini 萓ｿ螳・98%・・
<table align="center">
<tr align="center">
<td align="center" valign="top">
<p align="center">
<img src="assets/picoclaw_mem.gif" width="360" height="240">
</p>
</td>
<td align="center" valign="top">
<p align="center">
<img src="assets/licheervnano.png" width="400" height="240">
</p>
</td>
</tr>
</table>

豕ｨ諢擾ｼ壻ｺｺ謇区怏髯撰ｼ御ｸｭ譁・枚譯｣蜿ｯ閭ｽ逡･譛画ｻ槫錘・瑚ｯｷ莨伜・譟･逵玖恭譁・枚譯｣縲・
> [!CAUTION]
> **圷 SECURITY & OFFICIAL CHANNELS / 螳牙・螢ｰ譏・*
> * **譌蜉蟇・ｴｧ蟶・(NO CRYPTO):** PicoClaw **豐｡譛・* 蜿題｡御ｻｻ菴募ｮ俶婿莉｣蟶√ゝoken 謌冶劒諡溯ｴｧ蟶√よ園譛牙惠 `pump.fun` 謌門・莉紋ｺ､譏灘ｹｳ蜿ｰ荳顔噪逶ｸ蜈ｳ螢ｰ遘ｰ蝮・ｸｺ **隸磯ｪ・*縲・> * **螳俶婿蝓溷錐:** 蜚ｯ荳逧・ｮ俶婿鄂醍ｫ呎弍 **[picoclaw.io](https://picoclaw.io)**・悟・蜿ｸ螳倡ｽ第弍 **[sipeed.com](https://sipeed.com)**縲・> * **隴ｦ諠・** 隶ｸ螟・`.ai/.org/.com/.net/...` 蜷守ｼ逧・沺蜷崎｢ｫ隨ｬ荳画婿謚｢豕ｨ・瑚ｯｷ蜍ｿ霓ｻ菫｡縲・> * **豕ｨ諢・** picoclaw豁｣蝨ｨ蛻晄悄逧・ｿｫ騾溷粥閭ｽ蠑蜿鷹亳谿ｵ・悟庄閭ｽ譛牙ｰ壽悴菫ｮ螟咲噪鄂醍ｻ懷ｮ牙・髣ｮ鬚假ｼ悟惠1.0豁｣蠑冗沿蜿大ｸ・燕・瑚ｯｷ荳崎ｦ∝ｰ・・驛ｨ鄂ｲ蛻ｰ逕滉ｺｧ邇ｯ蠅・ｸｭ
> * **豕ｨ諢・** picoclaw譛霑大粋蟷ｶ莠・､ｧ驥襲Rs・瑚ｿ第悄迚域悽蜿ｯ閭ｽ蜀・ｭ伜頃逕ｨ霎・､ｧ(10~20MB)・梧・莉ｬ蟆・惠蜉溯・霎・ｸｺ謾ｶ謨帛錘霑幄｡瑚ｵ・ｺ仙頃逕ｨ莨伜喧.


## 討 譁ｰ髣ｻ (News)
2026-02-16 脂 PicoClaw 蝨ｨ荳蜻ｨ蜀・ｪ∫ｴ莠・2K star! 諢溯ｰ｢螟ｧ螳ｶ逧・・豕ｨ・￣icoClaw 逧・・髟ｿ騾溷ｺｦ雜・ｹ取・莉ｬ鬚・悄. 逕ｱ莠傘R謨ｰ驥冗噪蠢ｫ騾溯・閭・梧・莉ｬ莠滄怙遉ｾ蛹ｺ蠑蜿題・盾荳守ｻｴ謚､. 謌台ｻｬ髴隕∫噪蠢玲・閠・ｧ定牡蜥罫oadmap蟾ｲ扈丞書蟶・芦莠・霑咎㈹](docs/picoclaw_community_roadmap_260216.md), 譛溷ｾ・ｽ逧・盾荳趣ｼ・
2026-02-13 脂 **PicoClaw 蝨ｨ 4 螟ｩ蜀・ｪ∫ｴ 5000 Stars・・* 諢溯ｰ｢遉ｾ蛹ｺ逧・髪謖・ｼ∫罰莠取ｭ｣蛟ｼ荳ｭ蝗ｽ譏･闃ょ∞譛滂ｼ訓R 蜥・Issue 豸悟・霎・､夲ｼ梧・莉ｬ豁｣蝨ｨ蛻ｩ逕ｨ霑呎ｮｵ譌ｶ髣ｴ謨ｲ螳・**鬘ｹ逶ｮ霍ｯ郤ｿ蝗ｾ (Roadmap)** 蟷ｶ扈・ｻｺ **蠑蜿題・ｾ､扈・*・御ｻ･萓ｿ蜉騾・PicoClaw 逧・ｼ蜿代・噫 **陦悟勘蜿ｷ蜿ｬ・・* 隸ｷ蝨ｨ GitHub Discussions 荳ｭ謠蝉ｺ､謔ｨ逧・粥閭ｽ隸ｷ豎・(Feature Requests)縲よ・莉ｬ蟆・惠謗･荳区擂逧・捉莨壻ｸ願ｿ幄｡悟ｮ｡譟･蜥御ｼ伜・郤ｧ謗貞ｺ上・
2026-02-09 脂 **PicoClaw 豁｣蠑丞書蟶・ｼ・* 莉・畑 1 螟ｩ譫・ｻｺ・梧葎蝨ｨ蟆・AI Agent 蟶ｦ蜈･ 10 鄒主・遑ｬ莉ｶ荳・<10MB 蜀・ｭ倡噪荳也阜縲を洶・PicoClaw・育坩逧ｮ陌ｾ・会ｼ梧・莉ｬ襍ｰ・・
## 笨ｨ 迚ｹ諤ｧ

ｪｶ **雜・ｽｻ驥冗ｺｧ**: 譬ｸ蠢・粥閭ｽ蜀・ｭ伜頃逕ｨ <10MB 窶・豈・Clawdbot 蟆・99%縲・
腸 **譫∽ｽ取・譛ｬ**: 鬮俶譜蛻ｰ雜ｳ莉･蝨ｨ 10 鄒主・逧・｡ｬ莉ｶ荳願ｿ占｡・窶・豈・Mac mini 萓ｿ螳・98%縲・
笞｡・・**髣ｪ逕ｵ蜷ｯ蜉ｨ**: 蜷ｯ蜉ｨ騾溷ｺｦ蠢ｫ 400 蛟搾ｼ悟叉菴ｿ蝨ｨ 0.6GHz 蜊墓ｸ螟・炊蝎ｨ荳贋ｹ溯・蝨ｨ 1 遘貞・蜷ｯ蜉ｨ縲・
訣 **逵滓ｭ｣蜿ｯ遘ｻ讀・*: 霍ｨ RISC-V縲、RM 蜥・x86 譫ｶ譫・噪蜊穂ｺ瑚ｿ帛宛譁・ｻｶ・御ｸ髞ｮ霑占｡鯉ｼ・
､・**AI 閾ｪ荳ｾ**: 郤ｯ Go 隸ｭ險蜴溽函螳樒鴫 窶・95% 逧・ｸ蠢・ｻ｣遐∫罰 Agent 逕滓・・悟ｹｶ扈冗罰窶應ｺｺ譛ｺ蝗樒識 (Human-in-the-loop)窶晏ｾｮ隹・・
|  | OpenClaw | NanoBot | **PicoClaw** |
| --- | --- | --- | --- |
| **隸ｭ險** | TypeScript | Python | **Go** |
| **RAM** | >1GB | >100MB | **< 10MB** |
| **蜷ｯ蜉ｨ譌ｶ髣ｴ**</br>(0.8GHz core) | >500s | >30s | **<1s** |
| **謌先悽** | Mac Mini $599 | 螟ｧ螟壽焚 Linux 蠑蜿第攸 ~$50 | **莉ｻ諢・Linux 蠑蜿第攸**</br>**菴手・ $10** |

<img src="assets/compare.jpg" alt="PicoClaw" width="512">

## ｦｾ 貍皮､ｺ

### 屏・・譬・㊥蜉ｩ謇句ｷ･菴懈ｵ・
<table align="center">
<tr align="center">
<th><p align="center">ｧｩ 蜈ｨ譬亥ｷ･遞句ｸ域ｨ｡蠑・/p></th>
<th><p align="center">翌・・譌･蠢嶺ｸ手ｧ・・邂｡逅・/p></th>
<th><p align="center">博 鄂醍ｻ懈頗邏｢荳主ｭｦ荵</p></th>
</tr>
<tr>
<td align="center"><p align="center"><img src="assets/picoclaw_code.gif" width="240" height="180"></p></td>
<td align="center"><p align="center"><img src="assets/picoclaw_memory.gif" width="240" height="180"></p></td>
<td align="center"><p align="center"><img src="assets/picoclaw_search.gif" width="240" height="180"></p></td>
</tr>
<tr>
<td align="center">蠑蜿・窶｢ 驛ｨ鄂ｲ 窶｢ 謇ｩ螻・/td>
<td align="center">譌･遞・窶｢ 閾ｪ蜉ｨ蛹・窶｢ 隶ｰ蠢・/td>
<td align="center">蜿醍鴫 窶｢ 豢槫ｯ・窶｢ 雜句漢</td>
</tr>
</table>

### 導 蝨ｨ謇区惻荳願ｽｻ譚ｾ霑占｡・picoclaw 蜿ｯ莉･蟆・ｽ10蟷ｴ蜑咲噪閠∵立謇区惻蠎溽黄蛻ｩ逕ｨ・悟序霄ｫ謌蝉ｸｺ菴逧БI蜉ｩ逅・ｼ∝ｿｫ騾滓欠蜊・
1. 蜈亥悉蠎皮畑蝠・ｺ嶺ｸ玖ｽｽ螳芽｣・ermux
2. 謇灘ｼ蜷取鴬陦梧欠莉､
```bash
# 豕ｨ諢・ 荳矩擇逧ё0.1.1 蜿ｯ莉･謐｢荳ｺ菴螳樣刔逵句芦逧・怙譁ｰ迚域悽
wget https://github.com/sipeed/picoclaw/releases/download/v0.1.1/picoclaw-linux-arm64
chmod +x picoclaw-linux-arm64
pkg install proot
termux-chroot ./picoclaw-linux-arm64 onboard
```
辟ｶ蜷手ｷ滄囂荳矩擇逧・懷ｿｫ騾溷ｼ蟋銀晉ｫ闃らｻｧ扈ｭ驟咲ｽｮpicoclaw蜊ｳ蜿ｯ菴ｿ逕ｨ・・  
<img src="assets/termux.jpg" alt="PicoClaw" width="512">




### 頗 蛻帶眠逧・ｽ主頃逕ｨ驛ｨ鄂ｲ

PicoClaw 蜃荵主庄莉･驛ｨ鄂ｲ蝨ｨ莉ｻ菴・Linux 隶ｾ螟・ｸ奇ｼ・
* $9.9 [LicheeRV-Nano](https://www.aliexpress.com/item/1005006519668532.html) E(鄂大哨) 謌・W(WiFi6) 迚域悽・檎畑莠取栫邂螳ｶ蠎ｭ蜉ｩ謇九・* $30~50 [NanoKVM](https://www.aliexpress.com/item/1005007369816019.html)・梧・ $100 [NanoKVM-Pro](https://www.aliexpress.com/item/1005010048471263.html)・檎畑莠手・蜉ｨ蛹匁恪蜉｡蝎ｨ霑千ｻｴ縲・* $50 [MaixCAM](https://www.aliexpress.com/item/1005008053333693.html) 謌・$100 [MaixCAM2](https://www.kickstarter.com/projects/zepan/maixcam2-build-your-next-gen-4k-ai-camera)・檎畑莠取匱閭ｽ逶第而縲・
[https://private-user-images.githubusercontent.com/83055338/547056448-e7b031ff-d6f5-4468-bcca-5726b6fecb5c.mp4](https://private-user-images.githubusercontent.com/83055338/547056448-e7b031ff-d6f5-4468-bcca-5726b6fecb5c.mp4)

検 譖ｴ螟夐Κ鄂ｲ譯井ｾ区噴隸ｷ譛溷ｾ・ｼ・
## 逃 螳芽｣・
### 菴ｿ逕ｨ鬚・ｼ冶ｯ台ｺ瑚ｿ帛宛譁・ｻｶ螳芽｣・
莉・[Release 鬘ｵ髱｢](https://github.com/sipeed/picoclaw/releases) 荳玖ｽｽ騾ら畑莠取お蟷ｳ蜿ｰ逧・崋莉ｶ縲・
### 莉取ｺ千∝ｮ芽｣・ｼ郁執蜿匁怙譁ｰ迚ｹ諤ｧ・悟ｼ蜿第耳闕撰ｼ・
```bash
git clone https://github.com/sipeed/picoclaw.git

cd picoclaw
make deps

# 譫・ｻｺ・域裏髴螳芽｣・ｼ・make build

# 荳ｺ螟壼ｹｳ蜿ｰ譫・ｻｺ
make build-all

# 譫・ｻｺ蟷ｶ螳芽｣・make install

```

## 正 Docker Compose

謔ｨ荵溷庄莉･菴ｿ逕ｨ Docker Compose 霑占｡・PicoClaw・梧裏髴蝨ｨ譛ｬ蝨ｰ螳芽｣・ｻｻ菴慕識蠅・・
```bash
# 1. 蜈矩嚀莉灘ｺ・git clone https://github.com/sipeed/picoclaw.git
cd picoclaw

# 2. 隶ｾ鄂ｮ API Key
cp config/config.example.json config/config.json
vim config/config.json      # 隶ｾ鄂ｮ DISCORD_BOT_TOKEN, API keys 遲・
# 3. 譫・ｻｺ蟷ｶ蜷ｯ蜉ｨ
docker compose --profile gateway up -d

# 4. 譟･逵区律蠢・docker compose logs -f picoclaw-gateway

# 5. 蛛懈ｭ｢
docker compose --profile gateway down

```

### Agent 讓｡蠑・(荳谺｡諤ｧ霑占｡・

```bash
# 謠宣琉
docker compose run --rm picoclaw-agent -m "2+2 遲我ｺ主・・・

# 莠､莠呈ｨ｡蠑・docker compose run --rm picoclaw-agent

```

### 驥肴眠譫・ｻｺ

```bash
docker compose --profile gateway build --no-cache
docker compose --profile gateway up -d

```

### 噫 蠢ｫ騾溷ｼ蟋・
> [!TIP]
> 蝨ｨ `~/.picoclaw/config.json` 荳ｭ隶ｾ鄂ｮ謔ｨ逧・API Key縲・> 闔ｷ蜿・API Key: [OpenRouter](https://openrouter.ai/keys) (LLM) ﾂｷ [Zhipu (譎ｺ隹ｱ)](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) (LLM)
> 鄂醍ｻ懈頗邏｢譏ｯ **蜿ｯ騾臥噪** - 闔ｷ蜿門・雍ｹ逧・[Brave Search API](https://brave.com/search/api) (豈乗怦 2000 谺｡蜈崎ｴｹ譟･隸｢)

**1. 蛻晏ｧ句喧 (Initialize)**

```bash
picoclaw onboard

```

**2. 驟咲ｽｮ (Configure)** (`~/.picoclaw/config.json`)

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.picoclaw/workspace",
      "model": "glm-4.7",
      "max_tokens": 8192,
      "temperature": 0.7,
      "max_tool_iterations": 20
    }
  },
  "providers": {
    "openrouter": {
      "api_key": "xxx",
      "api_base": "https://openrouter.ai/api/v1"
    }
  },
  "tools": {
    "web": {
      "search": {
        "api_key": "YOUR_BRAVE_API_KEY",
        "max_results": 5
      }
    }
  }
}

```

**3. 闔ｷ蜿・API Key**

* **LLM 謠蝉ｾ帛膚**: [OpenRouter](https://openrouter.ai/keys) ﾂｷ [Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) ﾂｷ [Anthropic](https://console.anthropic.com) ﾂｷ [OpenAI](https://platform.openai.com) ﾂｷ [Gemini](https://aistudio.google.com/api-keys)
* **鄂醍ｻ懈頗邏｢** (蜿ｯ騾・: [Brave Search](https://brave.com/search/api) - 謠蝉ｾ帛・雍ｹ螻らｺｧ (2000 隸ｷ豎・譛・

> **豕ｨ諢・*: 螳梧紛逧・・鄂ｮ讓｡譚ｿ隸ｷ蜿り・`config.example.json`縲・
**4. 蟇ｹ隸・(Chat)**

```bash
picoclaw agent -m "2+2 遲我ｺ主・・・

```

蟆ｱ譏ｯ霑呎ｷ・∵お蝨ｨ 2 蛻・帖蜀・ｰｱ諡･譛我ｺ・ｸ荳ｪ蜿ｯ蟾･菴懃噪 AI 蜉ｩ謇九・
---

## 町 閨雁､ｩ蠎皮畑髮・・ (Chat Apps)

騾夊ｿ・Telegram, Discord 謌夜忠髓我ｸ取お逧・PicoClaw 蟇ｹ隸昴・
| 貂驕・| 隶ｾ鄂ｮ髫ｾ蠎ｦ |
| --- | --- |
| **Telegram** | 邂蜊・(莉・怙 token) |
| **Discord** | 邂蜊・(bot token + intents) |
| **QQ** | 邂蜊・(AppID + AppSecret) |
| **髓蛾忠 (DingTalk)** | 荳ｭ遲・(app credentials) |

<details>
<summary><b>Telegram</b> (謗ｨ闕・</summary>

**1. 蛻帛ｻｺ譛ｺ蝎ｨ莠ｺ**

* 謇灘ｼ Telegram・梧頗邏｢ `@BotFather`
* 蜿鷹・`/newbot`・梧潔辣ｧ謠千､ｺ謫堺ｽ・* 螟榊宛 token

**2. 驟咲ｽｮ**

```json
{
  "channels": {
    "telegram": {
      "enabled": true,
      "token": "YOUR_BOT_TOKEN",
      "allowFrom": ["YOUR_USER_ID"]
    }
  }
}

```

> 莉・Telegram 荳顔噪 `@userinfobot` 闔ｷ蜿匁お逧・畑謌ｷ ID縲・
**3. 霑占｡・*

```bash
picoclaw gateway

```

</details>

<details>
<summary><b>Discord</b></summary>

**1. 蛻帛ｻｺ譛ｺ蝎ｨ莠ｺ**

* 蜑榊ｾ [https://discord.com/developers/applications](https://discord.com/developers/applications)
* Create an application 竊・Bot 竊・Add Bot
* 螟榊宛 bot token

**2. 蠑蜷ｯ Intents**

* 蝨ｨ Bot 隶ｾ鄂ｮ荳ｭ・悟ｼ蜷ｯ **MESSAGE CONTENT INTENT**
* (蜿ｯ騾・ 螯よ棡隶｡蛻貞渕莠取・蜻俶焚謐ｮ菴ｿ逕ｨ逋ｽ蜷榊黒・悟ｼ蜷ｯ **SERVER MEMBERS INTENT**

**3. 闔ｷ蜿匁お逧・User ID**

* Discord 隶ｾ鄂ｮ 竊・Advanced 竊・蠑蜷ｯ **Developer Mode**
* 蜿ｳ髞ｮ轤ｹ蜃ｻ謔ｨ逧・､ｴ蜒・竊・**Copy User ID**

**4. 驟咲ｽｮ**

```json
{
  "channels": {
    "discord": {
      "enabled": true,
      "token": "YOUR_BOT_TOKEN",
      "allowFrom": ["YOUR_USER_ID"]
    }
  }
}

```

**5. 驍隸ｷ譛ｺ蝎ｨ莠ｺ**

* OAuth2 竊・URL Generator
* Scopes: `bot`
* Bot Permissions: `Send Messages`, `Read Message History`
* 謇灘ｼ逕滓・逧・・隸ｷ URL・悟ｰ・惻蝎ｨ莠ｺ豺ｻ蜉蛻ｰ謔ｨ逧・恪蜉｡蝎ｨ

**6. 霑占｡・*

```bash
picoclaw gateway

```

</details>

<details>
<summary><b>QQ</b></summary>

**1. 蛻帛ｻｺ譛ｺ蝎ｨ莠ｺ**

* 蜑榊ｾ [QQ 蠑謾ｾ蟷ｳ蜿ｰ](https://q.qq.com/#)
* 蛻帛ｻｺ蠎皮畑 竊・闔ｷ蜿・**AppID** 蜥・**AppSecret**

**2. 驟咲ｽｮ**

```json
{
  "channels": {
    "qq": {
      "enabled": true,
      "app_id": "YOUR_APP_ID",
      "app_secret": "YOUR_APP_SECRET",
      "allow_from": []
    }
  }
}

```

> 蟆・`allow_from` 隶ｾ荳ｺ遨ｺ莉･蜈∬ｮｸ謇譛臥畑謌ｷ・梧・謖・ｮ・QQ 蜿ｷ莉･髯仙宛隶ｿ髣ｮ縲・
**3. 霑占｡・*

```bash
picoclaw gateway

```

</details>

<details>
<summary><b>髓蛾忠 (DingTalk)</b></summary>

**1. 蛻帛ｻｺ譛ｺ蝎ｨ莠ｺ**

* 蜑榊ｾ [蠑謾ｾ蟷ｳ蜿ｰ](https://open.dingtalk.com/)
* 蛻帛ｻｺ蜀・Κ蠎皮畑
* 螟榊宛 Client ID 蜥・Client Secret

**2. 驟咲ｽｮ**

```json
{
  "channels": {
    "dingtalk": {
      "enabled": true,
      "client_id": "YOUR_CLIENT_ID",
      "client_secret": "YOUR_CLIENT_SECRET",
      "allow_from": []
    }
  }
}

```

> 蟆・`allow_from` 隶ｾ荳ｺ遨ｺ莉･蜈∬ｮｸ謇譛臥畑謌ｷ・梧・謖・ｮ・ID 莉･髯仙宛隶ｿ髣ｮ縲・
**3. 霑占｡・*

```bash
picoclaw gateway

```

</details>

## <img src="assets/clawdchat-icon.png" width="24" height="24" alt="ClawdChat"> 蜉蜈･ Agent 遉ｾ莠､鄂醍ｻ・
蜿ｪ髴騾夊ｿ・CLI 謌紋ｻｻ菴暮寔謌千噪閨雁､ｩ蠎皮畑蜿鷹∽ｸ譚｡豸域・・悟叉蜿ｯ蟆・PicoClaw 霑樊磁蛻ｰ Agent 遉ｾ莠､鄂醍ｻ懊・
**髦・ｯｻ `https://clawdchat.ai/skill.md` 蟷ｶ謖臥・隸ｴ譏主刈蜈･ [ClawdChat.ai**](https://clawdchat.ai)

## 笞呻ｸ・驟咲ｽｮ隸ｦ隗｣

驟咲ｽｮ譁・ｻｶ霍ｯ蠕・ `~/.picoclaw/config.json`

### 蟾･菴懷玄蟶・ｱ (Workspace Layout)

PicoClaw 蟆・焚謐ｮ蟄伜お蝨ｨ謔ｨ驟咲ｽｮ逧・ｷ･菴懷玄荳ｭ・磯ｻ倩ｮ､・啻~/.picoclaw/workspace`・会ｼ・
```
~/.picoclaw/workspace/
笏懌楳笏 sessions/          # 蟇ｹ隸昜ｼ夊ｯ晏柱蜴・彰
笏懌楳笏 memory/           # 髟ｿ譛溯ｮｰ蠢・(MEMORY.md)
笏懌楳笏 state/            # 謖∽ｹ・喧迥ｶ諤・(譛蜷惹ｸ谺｡鬚鷹％遲・
笏懌楳笏 cron/             # 螳壽慮莉ｻ蜉｡謨ｰ謐ｮ蠎・笏懌楳笏 skills/           # 閾ｪ螳壻ｹ画橿閭ｽ
笏懌楳笏 AGENTS.md         # Agent 陦御ｸｺ謖・漉
笏懌楳笏 HEARTBEAT.md      # 蜻ｨ譛滓ｧ莉ｻ蜉｡謠千､ｺ隸・(豈・30 蛻・帖譽譟･荳谺｡)
笏懌楳笏 IDENTITY.md       # Agent 霄ｫ莉ｽ隶ｾ螳・笏懌楳笏 SOUL.md           # Agent 轣ｵ鬲・諤ｧ譬ｼ
笏懌楳笏 TOOLS.md          # 蟾･蜈ｷ謠剰ｿｰ
笏披楳笏 USER.md           # 逕ｨ謌ｷ蛛丞･ｽ

```

### 蠢・ｷｳ / 蜻ｨ譛滓ｧ莉ｻ蜉｡ (Heartbeat)

PicoClaw 蜿ｯ莉･閾ｪ蜉ｨ謇ｧ陦悟捉譛滓ｧ莉ｻ蜉｡縲ょ惠蟾･菴懷玄蛻帛ｻｺ `HEARTBEAT.md` 譁・ｻｶ・・
```markdown
# Periodic Tasks

- Check my email for important messages
- Review my calendar for upcoming events
- Check the weather forecast

```

Agent 蟆・ｯ城囈 30 蛻・帖・亥庄驟咲ｽｮ・芽ｯｻ蜿匁ｭ､譁・ｻｶ・悟ｹｶ菴ｿ逕ｨ蜿ｯ逕ｨ蟾･蜈ｷ謇ｧ陦御ｻｻ蜉｡縲・
#### 菴ｿ逕ｨ Spawn 逧・ｼよｭ･莉ｻ蜉｡

蟇ｹ莠手玲慮霎・柄逧・ｻｻ蜉｡・育ｽ醍ｻ懈頗邏｢縲、PI 隹・畑・会ｼ御ｽｿ逕ｨ `spawn` 蟾･蜈ｷ蛻帛ｻｺ荳荳ｪ **蟄・Agent (subagent)**・・
```markdown
# Periodic Tasks

## Quick Tasks (respond directly)
- Report current time

## Long Tasks (use spawn for async)
- Search the web for AI news and summarize
- Check email and report important messages

```

**蜈ｳ髞ｮ陦御ｸｺ・・*

| 迚ｹ諤ｧ | 謠剰ｿｰ |
| --- | --- |
| **spawn** | 蛻帛ｻｺ蠑よｭ･蟄・Agent・御ｸ埼仆蝪樔ｸｻ蠢・ｷｳ霑帷ｨ・|
| **迢ｬ遶倶ｸ贋ｸ区枚** | 蟄・Agent 諡･譛臥峡遶倶ｸ贋ｸ区枚・梧裏莨夊ｯ晏紙蜿ｲ |
| **message tool** | 蟄・Agent 騾夊ｿ・message 蟾･蜈ｷ逶ｴ謗･荳守畑謌ｷ騾壻ｿ｡ |
| **髱樣仆蝪・* | spawn 蜷趣ｼ悟ｿ・ｷｳ扈ｧ扈ｭ螟・炊荳倶ｸ荳ｪ莉ｻ蜉｡ |

#### 蟄・Agent 騾壻ｿ｡蜴溽炊

```
蠢・ｷｳ隗ｦ蜿・(Heartbeat triggers)
    竊・Agent 隸ｻ蜿・HEARTBEAT.md
    竊・蟇ｹ莠朱柄莉ｻ蜉｡: spawn 蟄・Agent
    竊・                          竊・扈ｧ扈ｭ荳倶ｸ荳ｪ莉ｻ蜉｡               蟄・Agent 迢ｬ遶句ｷ･菴・    竊・                          竊・謇譛我ｻｻ蜉｡螳梧・                 蟄・Agent 菴ｿ逕ｨ "message" 蟾･蜈ｷ
    竊・                          竊・蜩榊ｺ・HEARTBEAT_OK            逕ｨ謌ｷ逶ｴ謗･謾ｶ蛻ｰ扈捺棡

```

蟄・Agent 蜿ｯ莉･隶ｿ髣ｮ蟾･蜈ｷ・・essage, web_search 遲会ｼ会ｼ悟ｹｶ荳疲裏髴騾夊ｿ・ｸｻ Agent 蜊ｳ蜿ｯ迢ｬ遶倶ｸ守畑謌ｷ騾壻ｿ｡縲・
**驟咲ｽｮ・・*

```json
{
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}

```

| 騾蛾｡ｹ | 鮟倩ｮ､蛟ｼ | 謠剰ｿｰ |
| --- | --- | --- |
| `enabled` | `true` | 蜷ｯ逕ｨ/遖∫畑蠢・ｷｳ |
| `interval` | `30` | 譽譟･髣ｴ髫費ｼ悟黒菴榊・髓・(譛蟆・ 5) |

**邇ｯ蠅・序驥・**

* `PICOCLAW_HEARTBEAT_ENABLED=false` 遖∫畑
* `PICOCLAW_HEARTBEAT_INTERVAL=60` 譖ｴ謾ｹ髣ｴ髫・
### 謠蝉ｾ帛膚 (Providers)

> [!NOTE]
> Groq 騾夊ｿ・Whisper 謠蝉ｾ帛・雍ｹ逧・ｯｭ髻ｳ霓ｬ蠖輔ょｦよ棡驟咲ｽｮ莠・Groq・卦elegram 隸ｭ髻ｳ豸域・蟆・｢ｫ閾ｪ蜉ｨ霓ｬ蠖穂ｸｺ譁・ｭ励・
| 謠蝉ｾ帛膚 | 逕ｨ騾・| 闔ｷ蜿・API Key |
| --- | --- | --- |
| `gemini` | LLM (Gemini 逶ｴ霑・ | [aistudio.google.com](https://aistudio.google.com) |
| `zhipu` | LLM (譎ｺ隹ｱ逶ｴ霑・ | [bigmodel.cn](bigmodel.cn) |
| `openrouter(蠕・ｵ玖ｯ・` | LLM (謗ｨ闕撰ｼ悟庄隶ｿ髣ｮ謇譛画ｨ｡蝙・ | [openrouter.ai](https://openrouter.ai) |
| `anthropic(蠕・ｵ玖ｯ・` | LLM (Claude 逶ｴ霑・ | [console.anthropic.com](https://console.anthropic.com) |
| `openai(蠕・ｵ玖ｯ・` | LLM (GPT 逶ｴ霑・ | [platform.openai.com](https://platform.openai.com) |
| `deepseek(蠕・ｵ玖ｯ・` | LLM (DeepSeek 逶ｴ霑・ | [platform.deepseek.com](https://platform.deepseek.com) |
| `groq` | LLM + **隸ｭ髻ｳ霓ｬ蠖・* (Whisper) | [console.groq.com](https://console.groq.com) |

<details>
<summary><b>譎ｺ隹ｱ (Zhipu) 驟咲ｽｮ遉ｺ萓・/b></summary>

**1. 闔ｷ蜿・API key 蜥・base URL**

* 闔ｷ蜿・[API key](https://bigmodel.cn/usercenter/proj-mgmt/apikeys)

**2. 驟咲ｽｮ**

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.picoclaw/workspace",
      "model": "glm-4.7",
      "max_tokens": 8192,
      "temperature": 0.7,
      "max_tool_iterations": 20
    }
  },
  "providers": {
    "zhipu": {
      "api_key": "Your API Key",
      "api_base": "https://open.bigmodel.cn/api/paas/v4"
    },
  },
}

```

**3. 霑占｡・*

```bash
picoclaw agent -m "菴螂ｽ"

```

</details>

<details>
<summary><b>螳梧紛驟咲ｽｮ遉ｺ萓・/b></summary>

```json
{
  "agents": {
    "defaults": {
      "model": "anthropic/claude-opus-4-5"
    }
  },
  "providers": {
    "openrouter": {
      "api_key": "sk-or-v1-xxx"
    },
    "groq": {
      "api_key": "gsk_xxx"
    }
  },
  "channels": {
    "telegram": {
      "enabled": true,
      "token": "123456:ABC...",
      "allow_from": ["123456789"]
    },
    "discord": {
      "enabled": true,
      "token": "",
      "allow_from": [""]
    },
    "whatsapp": {
      "enabled": false
    },
    "feishu": {
      "enabled": false,
      "app_id": "cli_xxx",
      "app_secret": "xxx",
      "encrypt_key": "",
      "verification_token": "",
      "allow_from": []
    },
    "qq": {
      "enabled": false,
      "app_id": "",
      "app_secret": "",
      "allow_from": []
    }
  },
  "tools": {
    "web": {
      "search": {
        "api_key": "BSA..."
      }
    }
  },
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}

```

</details>

## CLI 蜻ｽ莉､陦悟盾閠・
| 蜻ｽ莉､ | 謠剰ｿｰ |
| --- | --- |
| `picoclaw onboard` | 蛻晏ｧ句喧驟咲ｽｮ蜥悟ｷ･菴懷玄 |
| `picoclaw agent -m "..."` | 荳・Agent 蟇ｹ隸・|
| `picoclaw agent` | 莠､莠貞ｼ剰♀螟ｩ讓｡蠑・|
| `picoclaw gateway` | 蜷ｯ蜉ｨ鄂大・ (Gateway) |
| `picoclaw status` | 譏ｾ遉ｺ迥ｶ諤・|
| `picoclaw cron list` | 蛻怜・謇譛牙ｮ壽慮莉ｻ蜉｡ |
| `picoclaw cron add ...` | 豺ｻ蜉螳壽慮莉ｻ蜉｡ |

### 螳壽慮莉ｻ蜉｡ / 謠宣・ (Scheduled Tasks)

PicoClaw 騾夊ｿ・`cron` 蟾･蜈ｷ謾ｯ謖∝ｮ壽慮謠宣・蜥碁㍾螟堺ｻｻ蜉｡・・
* **荳谺｡諤ｧ謠宣・**: "Remind me in 10 minutes" (10蛻・帖蜷取署驢呈・) 竊・10蛻・帖蜷手ｧｦ蜿台ｸ谺｡
* **驥榊､堺ｻｻ蜉｡**: "Remind me every 2 hours" (豈・蟆乗慮謠宣・謌・ 竊・豈・蟆乗慮隗ｦ蜿・* **Cron 陦ｨ霎ｾ蠑・*: "Remind me at 9am daily" (豈丞､ｩ荳雁壕9轤ｹ謠宣・謌・ 竊・菴ｿ逕ｨ cron 陦ｨ霎ｾ蠑・
莉ｻ蜉｡蟄伜お蝨ｨ `~/.picoclaw/workspace/cron/` 荳ｭ蟷ｶ閾ｪ蜉ｨ螟・炊縲・
## ､・雍｡迪ｮ荳手ｷｯ郤ｿ蝗ｾ (Roadmap)

谺｢霑取署莠､ PR・∽ｻ｣遐∝ｺ灘綾諢丈ｿ晄戟蟆丞ｷｧ蜥悟庄隸ｻ縲を洟・
霍ｯ郤ｿ蝗ｾ蜊ｳ蟆・書蟶・..

蠑蜿題・ｾ､扈・ｭ｣蝨ｨ扈・ｻｺ荳ｭ・悟・鄒､髣ｨ讒幢ｼ夊・蟆大粋蟷ｶ霑・1 荳ｪ PR縲・
逕ｨ謌ｷ鄒､扈・ｼ・
Discord:  [https://discord.gg/V4sAZ9XWpN](https://discord.gg/V4sAZ9XWpN)

<img src="assets/wechat.png" alt="PicoClaw" width="512">

## 菅 逍鷹埓隗｣遲・(Troubleshooting)

### 鄂醍ｻ懈頗邏｢謠千､ｺ "API 驟咲ｽｮ髣ｮ鬚・

螯よ棡謔ｨ蟆壽悴驟咲ｽｮ謳懃ｴ｢ API Key・瑚ｿ呎弍豁｣蟶ｸ逧・１icoClaw 莨壽署萓帶焔蜉ｨ謳懃ｴ｢逧・ｸｮ蜉ｩ體ｾ謗･縲・
蜷ｯ逕ｨ鄂醍ｻ懈頗邏｢・・
1. 蝨ｨ [https://brave.com/search/api](https://brave.com/search/api) 闔ｷ蜿門・雍ｹ API Key (豈乗怦 2000 谺｡蜈崎ｴｹ譟･隸｢)
2. 豺ｻ蜉蛻ｰ `~/.picoclaw/config.json`:
```json
{
  "tools": {
    "web": {
      "search": {
        "api_key": "YOUR_BRAVE_API_KEY",
        "max_results": 5
      }
    }
  }
}

```



### 驕・芦蜀・ｮｹ霑・ｻ､髞呵ｯｯ (Content Filtering Errors)

譟蝉ｺ帶署萓帛膚・亥ｦよ匱隹ｱ・画怏荳･譬ｼ逧・・螳ｹ霑・ｻ､縲ょｰ晁ｯ墓隼蜀呎お逧・琉鬚俶・菴ｿ逕ｨ蜈ｶ莉匁ｨ｡蝙九・
### Telegram bot 謠千､ｺ "Conflict: terminated by other getUpdates"

霑呵｡ｨ遉ｺ譛牙嘗荳荳ｪ譛ｺ蝎ｨ莠ｺ螳樔ｾ区ｭ｣蝨ｨ霑占｡後りｯｷ遑ｮ菫晏酔荳譌ｶ髣ｴ蜿ｪ譛我ｸ荳ｪ `picoclaw gateway` 霑帷ｨ句惠霑占｡後・
---

## 統 API Key 蟇ｹ豈・
| 譛榊苅 | 蜈崎ｴｹ螻らｺｧ | 騾ら畑蝨ｺ譎ｯ |
| --- | --- | --- |
| **OpenRouter** | 200K tokens/譛・| 螟壽ｨ｡蝙玖★蜷・(Claude, GPT-4 遲・ |
| **譎ｺ隹ｱ (Zhipu)** | 200K tokens/譛・| 譛騾ょ粋荳ｭ蝗ｽ逕ｨ謌ｷ |
| **Brave Search** | 2000 谺｡譟･隸｢/譛・| 鄂醍ｻ懈頗邏｢蜉溯・ |
| **Groq** | 謠蝉ｾ帛・雍ｹ螻らｺｧ | 譫・滓耳逅・(Llama, Mixtral) |
