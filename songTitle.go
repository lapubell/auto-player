package main

import (
	"strings"
)

func songTitle(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		return path
	}
	return parts[2]
}
