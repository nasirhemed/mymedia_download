package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/nasirhemed/mymedia_download/helper"
)

func DownloadVideo(url string, output string) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	videoId := helper.ExtractVideoId(url)

	chunkId := helper.GetChunkId(videoId)
	fmt.Println(chunkId)

	chunkList := helper.GetChunkList(videoId, chunkId)
	helper.WriteChunkList("temp.txt", chunkList)

	helper.DownloadFiles(videoId, chunkList)

	helper.FfmpegCommand("temp.txt", output)

	helper.CleanUp(chunkList, "temp.txt")
}

func main() {
	var opts struct {
		Url    string `long:"url" description:"url of the video to download" required:"true"`
		Output string `short:"o" long:"output" description:"Video output file name" default:"out.mp4"`
	}

	parser := flags.NewParser(&opts, 0)

	_, err := parser.Parse()
	// _, err := flags.Parse(&opts)

	if err != nil {
		// log.Fatal(err)
		parser.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	DownloadVideo(opts.Url, opts.Output)

}
