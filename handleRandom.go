package main

import (
	"net/http"
)

func handleRandom(w http.ResponseWriter, r *http.Request) {
	currentState.IsRandomMode = !currentState.IsRandomMode
	currentState.setMusicList()

	getState(w)
}
