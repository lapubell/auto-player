package main

import (
	_ "embed"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//go:embed index.html
var html string

//go:embed kiddo.svg
var kiddo string

//go:embed parents.img
var parents string

func handleIndex(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 100) // short pause because concurrancy and I'm lazy

	// start with the embeded HTML
	output := html
	output = strings.ReplaceAll(output, "%%VOLUME%%", strconv.Itoa(volume))
	output = strings.ReplaceAll(output, "%%MAX_VOLUME%%", strconv.Itoa(maxVolume))

	output = strings.ReplaceAll(output, "%%KIDDO%%", kiddo)
	output = strings.ReplaceAll(output, "%%PARENTS%%", parents)
	if !isPlaying {
		output = strings.ReplaceAll(output, "%%CURRENT%%", "")
		output = strings.ReplaceAll(output, "%%START_DISPLAY%%", "flex")
		output = strings.ReplaceAll(output, "%%STOP_DISPLAY%%", "none")
	} else {
		output = strings.ReplaceAll(output, "%%CURRENT%%", songTitle(currentSong))
		output = strings.ReplaceAll(output, "%%START_DISPLAY%%", "none")
		output = strings.ReplaceAll(output, "%%STOP_DISPLAY%%", "flex")
	}

	if isBabySafe {
		output = strings.ReplaceAll(output, "%%BABY_DISPAY%%", "block")
		output = strings.ReplaceAll(output, "%%PARENT_DISPAY%%", "none")
	} else {
		output = strings.ReplaceAll(output, "%%BABY_DISPAY%%", "none")
		output = strings.ReplaceAll(output, "%%PARENT_DISPAY%%", "block")
	}

	w.Write([]byte(output))
}
