package main

import (
	"os"
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
	return songs, nil
}
