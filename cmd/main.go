package main

import (
	"PllDownloader/internal"
	"fmt"
	"log"
)

func main() {
	playlistURL := "https://www.youtube.com/playlist?list=PL2TLRFhnMM6yisZhf730yy4T08Q1qf_du"

	videos, err := internal.ExtractPlaylistVideos(playlistURL)
	if err != nil {
		log.Fatal("Failed to extract videos from playlist:", err)
	}

	fmt.Println("[INFO] Videos found:", len(videos))
	for _, v := range videos {
		fmt.Println("[INFO] Video URL:", v)
	}

	for _, v := range videos {
		err := internal.DownloadVideo(v)
		if err != nil {
			log.Println("[ERROR] Failed to download video:", v, "-", err)
		}
	}
}
