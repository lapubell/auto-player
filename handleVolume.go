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
		currentState.Volume = 0
	}
	currentState.Volume = newVolumeInt
	updatePlayerVolume()

	getState(w)
}
