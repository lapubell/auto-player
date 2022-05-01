package main

import (
	"math/rand"
	"os"
	"time"
)

func setMusicList() ([]string, error) {
	songs = []string{}

	files, err := os.ReadDir(musicDir)
	if err != nil {
		return songs, err
	}

	for _, f := range files {
		songs = append(songs, musicDir+"/"+f.Name())
	}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(songs), func(i, j int) {
		songs[i], songs[j] = songs[j], songs[i]
	})
	return songs, nil
}
