package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var musicDir string
var songs []string
var currentSong string
var mpgCmd *exec.Cmd
var amixerID string
var isPlaying bool
var isBabySafe bool = true
var volume int

func main() {
	godotenv.Load("/etc/server-run.env")

	// setup
	musicDir = os.Getenv("MUSIC_DIRECTORY")
	if musicDir == "" {
		musicDir = "/music"
	}

	amixerID = os.Getenv("AMIXER_ID")
	if amixerID == "" {
		amixerID = "Power Amplifier"
	}

	webPort := os.Getenv("PORT")
	if webPort == "" {
		webPort = "80"
	}

	volumeString := os.Getenv("VOLUME")
	if volumeString == "" {
		volumeString = "10"
	}
	volumeInt, err := strconv.Atoi(volumeString)
	if err != nil {
		volume = 0
	}
	volume = volumeInt

	autoPlayingString := os.Getenv("AUTO_PLAY")
	if autoPlayingString == "true" {
		isPlaying = true
	} else {
		isPlaying = false
	}

	songs, err = setMusicList()
	if err != nil {
		log.Fatal(err)
	}

	// set the volume on boot
	updatePlayerVolume()

	// go routine for the audio player (command launcher)
	go playASong()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/skip", handleSkip)
	http.HandleFunc("/stop", handleStop)
	http.HandleFunc("/play", handlePlay)
	http.HandleFunc("/volume", handleVolume)
	http.HandleFunc("/switch", handleSwitch)
	log.Fatal(http.ListenAndServe(":"+webPort, nil))
}
