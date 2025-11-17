package internal

import (
	"io"
	"net/http"
	"regexp"
)

func ExtractPlaylistVideos(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	html := string(body)

	re := regexp.MustCompile(`"videoId":"([a-zA-Z0-9_-]{11})"`)
	matches := re.FindAllStringSubmatch(html, -1)

	videoIDs := []string{}
	seen := map[string]bool{}
	for _, m := range matches {
		if !seen[m[1]] {
			videoIDs = append(videoIDs, "https://www.youtube.com/watch?v="+m[1])
			seen[m[1]] = true
		}
	}
	return videoIDs, nil
}
