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
		log.Fatal(err)
	}

	fmt.Println("Vídeos encontrados:", len(videos))
	for _, v := range videos {
		fmt.Println("Vídeo:", v)
	}

	for _, v := range videos {
		err := internal.DownloadVideo(v)
		if err != nil {
			log.Println("Erro ao baixar:", v, err)
		}
	}
}
