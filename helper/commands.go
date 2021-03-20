package helper

import (
	"fmt"
	"os/exec"
	"strings"
)

func FfmpegCommand(chunkListFile string, videoName string) {
	cmd := exec.Command("ffmpeg", "-y", "-f", "concat", "-i", chunkListFile, "-c", "copy", "-bsf:a", "aac_adtstoasc", videoName)
	fmt.Printf("cmd: \"%s\n\"", strings.Join(cmd.Args, " "))
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
