package main

import "net/http"

func handlePlay(w http.ResponseWriter, r *http.Request) {
	currentState.IsPlaying = true
	mpgCmd.Process.Kill()

	getState(w)
}
