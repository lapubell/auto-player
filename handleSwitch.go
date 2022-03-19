package main

import "net/http"

func handleSwitch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newTypeString := r.FormValue("type")

	if newTypeString == "baby" {
		isBabySafe = true
		musicDir = "/music"
	}
	if newTypeString == "grown-ups" {
		isBabySafe = false
		musicDir = "/grown-ups"
	}

	isPlaying = false
	mpgCmd.Process.Kill()

	http.Redirect(w, r, "/", 302)
}
