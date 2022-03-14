package main

import "net/http"

func handlePlay(w http.ResponseWriter, r *http.Request) {
	isPlaying = true
	mpgCmd.Process.Kill()

	http.Redirect(w, r, "/", 302)
}
