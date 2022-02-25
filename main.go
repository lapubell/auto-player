package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

var musicDir string
var currentSong string
var mpgCmd *exec.Cmd

func main() {
	musicDir = "/music"

	files, err := os.ReadDir(musicDir)
	if err != nil {
		log.Fatal(err)
	}

	songs := []string{}
	for _, f := range files {
		songs = append(songs, f.Name())
	}

	// go routine for the audio player (command launcher)
	go playASong(songs)

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/skip", handleSkip)
	log.Fatal(http.ListenAndServe(":8899", nil))
}
