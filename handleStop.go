package main

import (
	"net/http"
)

func handleStop(w http.ResponseWriter, r *http.Request) {
	currentState.IsPlaying = false
	mpgCmd.Process.Kill()

	getState(w)
}
