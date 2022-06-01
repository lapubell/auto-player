package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleSwitch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newTypeString := r.FormValue("type")

	// sound effects
	if newTypeString == "waves" {
		currentState.IsBabySafe = true
		currentState.Mode = "waves"
	}
	if newTypeString == "music" {
		currentState.IsBabySafe = true
		currentState.Mode = "music"
	}
	if newTypeString == "baby" {
		currentState.IsBabySafe = true
		currentState.Mode = "music"
	}
	if newTypeString == "grown-ups" {
		currentState.IsBabySafe = false
		currentState.Mode = "music"
	}

	currentState.IsPlaying = false
	currentState.setMusicList()

	mpgCmd.Process.Kill()
	time.Sleep(500 * time.Millisecond)

	fmt.Println(currentState)

	getState(w)
}
