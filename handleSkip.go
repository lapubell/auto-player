package main

import (
	"net/http"
	"time"
)

func handleSkip(w http.ResponseWriter, r *http.Request) {
	mpgCmd.Process.Kill()

	time.Sleep(500 * time.Millisecond)
	getState(w)
}
