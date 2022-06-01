package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func handleCurrentState(w http.ResponseWriter, r *http.Request) {
	getState(w)
}

func getState(w http.ResponseWriter) {
	if currentState.Mode == "waves" && currentState.IsPlaying {
		currentState.NowPlaying = "Waves Sound Effect"
	} else {
		if currentState.IsPlaying {
			currentState.NowPlaying = currentState.PlayList[currentState.currentSongIndex]
		} else {
			currentState.NowPlaying = ""
		}
	}

	var output []byte
	output, _ = json.Marshal(currentState)

	time.Sleep(time.Millisecond * 100)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
