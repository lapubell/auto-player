package main

import (
	_ "embed"
	"net/http"
	"strconv"
	"strings"
)

//go:embed index.html
var html string

//go:embed kiddo.svg
var kiddo string

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// start with the embeded HTML
	output := html
	output = strings.ReplaceAll(output, "%%VOLUME%%", strconv.Itoa(volume))

	output = strings.ReplaceAll(output, "%%KIDDO%%", kiddo)
	if !isPlaying {
		output = strings.ReplaceAll(output, "%%CURRENT%%", "")
		output = strings.ReplaceAll(output, "%%START_DISPLAY%%", "flex")
		output = strings.ReplaceAll(output, "%%STOP_DISPLAY%%", "none")
	} else {
		output = strings.ReplaceAll(output, "%%CURRENT%%", currentSong)
		output = strings.ReplaceAll(output, "%%START_DISPLAY%%", "none")
		output = strings.ReplaceAll(output, "%%STOP_DISPLAY%%", "flex")
	}

	w.Write([]byte(output))
}
