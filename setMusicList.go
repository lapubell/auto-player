package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func (cs *CurrentState) setMusicList() {
	cs.Albums = []string{"All"}
	cs.PlayList = []string{}
	cs.currentSongIndex = 0

	if cs.IsBabySafe {
		cs.musicDir = "/music"
	} else {
		cs.musicDir = "/grown-ups"
	}

	files, err := os.ReadDir(cs.musicDir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		album, track := parseFilename(f.Name())
		albumInListAlready := false

		for _, existingAlbum := range cs.Albums {
			if album == existingAlbum {
				albumInListAlready = true
				break
			}
		}
		if !albumInListAlready {
			cs.Albums = append(cs.Albums, album)
		}

		if cs.CurrentAlbum == "All" || cs.CurrentAlbum == album {
			cs.PlayList = append(cs.PlayList, track)
		}
	}

	if cs.IsRandomMode {
		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(cs.PlayList), func(i, j int) {
			cs.PlayList[i], cs.PlayList[j] = cs.PlayList[j], cs.PlayList[i]
		})
	}
	fmt.Println(cs.Albums)
	fmt.Println(cs.PlayList)
}

func parseFilename(filename string) (string, string) {
	parts := strings.Split(filename, "|")
	if len(parts) < 2 {
		return "Other", filename
	}
	return parts[0], filename
}
