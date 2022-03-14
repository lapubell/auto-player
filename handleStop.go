package main

import "net/http"

func handleStop(w http.ResponseWriter, r *http.Request) {
	isPlaying = false
	mpgCmd.Process.Kill()

	http.Redirect(w, r, "/", 302)
}
