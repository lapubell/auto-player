package main

import "net/http"

func handleSkip(w http.ResponseWriter, r *http.Request) {
	mpgCmd.Process.Kill()

	http.Redirect(w, r, "/", http.StatusFound)
}
