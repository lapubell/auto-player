package main

import (
	"os/exec"
	"strconv"
)

// read the global variable and send a system command to update amixer
func updatePlayerVolume() {
	volumeCommand := exec.Command("amixer", "sset", amixerID, strconv.Itoa(currentState.Volume))
	volumeCommand.Run()
}
