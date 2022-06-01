package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	_ "github.com/joho/godotenv/autoload"
)

var currentState CurrentState
var mpgCmd *exec.Cmd
var amixerID string

func main() {
	currentState = initializeState()

	amixerID = os.Getenv("AMIXER_ID")
	if amixerID == "" {
		amixerID = "Power Amplifier"
	}

	webPort := os.Getenv("PORT")
	if webPort == "" {
		webPort = "80"
	}

	// go routine for the audio player (command launcher)
	go playASong()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/state", handleCurrentState)
	http.HandleFunc("/skip", handleSkip)
	http.HandleFunc("/stop", handleStop)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/random", handleRandom)
	http.HandleFunc("/volume", handleVolume)
	http.HandleFunc("/album", handleSetAlbum)
	http.HandleFunc("/switch", handleSwitch)
	log.Fatal(http.ListenAndServe(":"+webPort, nil))
}
