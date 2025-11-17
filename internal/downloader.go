package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func DownloadVideo(url string) error {
	fmt.Println("[DOWNLOAD] Checking playlist/video:", url)

	outputDir := "./downloads/"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	fmt.Println("[DOWNLOAD] Creating downloads folder at", outputDir)

	cmd := exec.Command(
		"yt-dlp",
		"-x",
		"--audio-format", "mp3",
		"-o", filepath.Join(outputDir, "%(title)s.%(ext)s"),
		url,
	)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		return err
	}

	processPipe := func(scanner *bufio.Scanner) {
		for scanner.Scan() {
			line := scanner.Text()
			if containsDownloadInfo(line) {
				fmt.Println(Green + "[DOWNLOAD] " + line + Reset)
			}
		}
	}

	go processPipe(bufio.NewScanner(stdoutPipe))
	go processPipe(bufio.NewScanner(stderrPipe))

	if err := cmd.Wait(); err != nil {
		return err
	}

	fmt.Println(Green + "[DOWNLOAD] Completed:", url + Reset)
	return nil
}

func containsDownloadInfo(line string) bool {
	keywords := []string{"[download]", "[ExtractAudio]", "[ffmpeg]"}
	for _, k := range keywords {
		if contains(line, k) {
			return true
		}
	}
	return false
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(s) > len(substr) && (s[0:len(substr)] == substr || contains(s[1:], substr))))
}
