package main

import (
	"log"
	"math/rand"
	"os/exec"
	"time"
)

func playASong(songs []string) {
	rand.Seed(time.Now().Unix())

	for {
		randomIndex := rand.Intn(len(songs))
		currentSong = songs[randomIndex]
		log.Println("Now playing: " + currentSong)

		mpgCmd = exec.Command("mpg123", musicDir+"/"+currentSong)
		mpgCmd.Run()
	}
}
