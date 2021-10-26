# MyMedia Downloader
Command line program to download videos from MyMedia (UofT)

## Requirements
- go
- ffmpeg

## Installation

You can use `go install github.com/nasirhemed/mymedia_download` to install the program

```bash
go get install github.com/nasirhemed/mymedia_download
```

## Usage
For some reason, MyMedia uses a certificate authority that may not be installed in your computer. To overcome this, use the `--insecure` flag. This will ignore validating the certificate from the server. For something like downloading videos this shouldn't be a big deal.

```bash
mymedia_download --url https://play.library.utoronto.ca https://play.library.utoronto.ca/play/${videoId}
mymedia_download --url https://play.library.utoronto.ca https://play.library.utoronto.ca/play/${videoId} --output lecture2.mp4
mymedia_download --url https://play.library.utoronto.ca https://play.library.utoronto.ca/play/${videoId} --insecure
```
