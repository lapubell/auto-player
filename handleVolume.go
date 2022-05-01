package main

import (
	"net/http"
	"strconv"
)

func handleVolume(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newVolumeString := r.FormValue("value")

	newVolumeInt, err := strconv.Atoi(newVolumeString)
	if err != nil {
		volume = 0
	}
	volume = newVolumeInt
	updatePlayerVolume()

	http.Redirect(w, r, "/", http.StatusFound)
}
