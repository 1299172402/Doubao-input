# Doubao Input – Cross-Device Input Tool

> [中文](./README.md) | **English**

Open-source cross-device input tool that works with the Doubao mobile app – no extra software required. Text from voice or handwriting on your phone can be sent directly to any input field on your computer, smoothly bridging daily work and creative flows.

### ✨ Highlights

- High-accuracy voice recognition – works in noisy environments, supports punctuation and mixed Chinese/English input.
- Real‑time sync – automatically writes to system clipboard; paste with `Ctrl+V` whenever needed.
- **Auto‑type mode** – text flows directly into the focused input field without manual pasting.
- System tray – runs quietly in the background.
- Web‑based configuration – no need to edit config files manually.
- Silent startup & auto-launch on boot (Windows / Linux / macOS).

### 🎯 Use Cases

- **Coding assistance** – feed voice input into CLI tools like OpenCode or Claude Code.
- **Wireless handwriting** – use your phone as a handwriting pad without extra drivers.
- **Elderly‑friendly input** – no need for Pinyin or mouse-based handwriting; just speak or write.

### 📥 Download

Get the latest executable from [Releases](https://github.com/1299172402/Doubao-input/releases) or [Lanzou Cloud](https://zhiyuyu.lanzout.com/b01tqc738d?pwd=57rw) (password: `57rw`).

**Supported platforms:**
- Windows (amd64 / i386)
- Linux (amd64)
- macOS (amd64 / arm64)

### 🚀 Quick Start

1.  **Get configuration data:**
    - Open [Doubao Web](https://www.doubao.com), log in and start a conversation.
    - Press `F12` → **Network** tab.
    - Send a message and find the latest `single` request (to `https://www.doubao.com/im/chain/single`).
    - Right‑click it → **Copy** → **Copy as cURL (Bash)**.

2.  **Launch and configure:**
    - Run `doubao-input.exe` – your browser will open the configuration page.
    - Paste the copied cURL command and click **💾 Save**.
    - Click **🚀 Fetch Messages** to test.

3.  **Start using:**
    - On your phone, speak or handwrite into the **same Doubao conversation**.
    - Back on your computer, open any text field and press `Ctrl+V` – the text appears!
    - With **auto‑type** enabled, text is directly inserted where your cursor blinks.

### 💴 Donate

If this tool helps you, consider buying me a coffee ☕️.  
For custom development or cooperation, feel free to contact me.

| WeChat Pay | Alipay |
|------------|--------|
| <img src="./assets/static/donate/wechat.png" height="300"> | <img src="./assets/static/donate/alipay.jpg" height="300"> |

### 🚀 Acknowledgements

Special thanks to **ByteDance** and **Doubao** for their excellent speech recognition technology and product experience. Doubao’s voice input outperforms many alternatives (Microsoft Speech, OpenAI Whisper, etc.) in speed, accuracy, mixed language, punctuation, low‑volume speech, noisy environments, far‑field recognition, and the cocktail party effect.

> **Disclaimer:** This is an independent open‑source tool with no official affiliation with ByteDance or Doubao. It is for research and personal use only. The developer assumes no legal responsibility for any issues or consequences arising from its use. All trademarks and product names are the property of their respective owners.
