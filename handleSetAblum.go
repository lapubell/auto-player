package main

import (
	"net/http"
)

func handleSetAlbum(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	currentState.CurrentAlbum = r.FormValue("value")
	currentState.IsPlaying = false
	currentState.setMusicList()

	mpgCmd.Process.Kill()
	getState(w)
}
