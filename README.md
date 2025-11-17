# 🎬 Mini YT-DLP in Go

A simplified Go implementation of **YT-DLP** that downloads YouTube videos and playlists and automatically converts them to **MP3 audio**, with clean and user-friendly logs.

---

## 🚀 Features

- Download individual videos or entire playlists from YouTube.  
- Extract audio and convert to MP3 automatically.  
- Custom, clean logs instead of raw command-line output.  
- Organizes downloaded files in a dedicated folder (`./downloads`).  
- Built in Go, using external tools like **FFmpeg** for audio conversion.  

---

## ⚙️ Requirements

- [Go](https://golang.org/dl/) v1.20+  
- [FFmpeg](https://ffmpeg.org/) installed and accessible in your PATH  
- [yt-dlp](https://github.com/yt-dlp/yt-dlp) installed and accessible in your PATH  

---

## 📦 Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/mini-yt-dlp-go.git
cd mini-yt-dlp-go
