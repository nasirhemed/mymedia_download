package helper

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ExtractVideoId(url string) string {
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	splitted := strings.Split(url, "/")
	return splitted[len(splitted)-1]
}

func downloadChunk(videoId string, chunkFile string) {

	respBody := GetChunk(videoId, chunkFile)

	file, err := os.Create(chunkFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	file.Write(respBody)

}

func DownloadFiles(videoId string, chunkList []string) {
	done := make(chan bool)
	for i := 0; i < len(chunkList); i++ {

		go func(index int) {
			fmt.Printf("Downloading file %s\n", chunkList[index])
			downloadChunk(videoId, chunkList[index])

			done <- true
		}(i)

	}

	for i := 0; i < len(chunkList); i++ {
		<-done
	}
}

func WriteChunkList(filename string, chunkList []string) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	for _, chunk := range chunkList {
		chunkLine := fmt.Sprintf("file %s \n", chunk)
		file.Write([]byte(chunkLine))
	}
}

func CleanUp(chunkFiles []string, tempFile string) {

	os.Remove(tempFile)

	for _, file := range chunkFiles {
		err := os.Remove(file)

		if err != nil {
			log.Fatal(err)
		}
	}

	return
}
