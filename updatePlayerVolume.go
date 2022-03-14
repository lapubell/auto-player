package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

// read the global variable and send a system command to update amixer
func updatePlayerVolume() {
	volumeCommand := exec.Command("amixer", "sset", amixerID, strconv.Itoa(volume))
	err := volumeCommand.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
