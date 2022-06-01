#! /usr/bin/bash

# kill any running processes
killall mpg123
killall auto-player
killall play

# build for CHIP
GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o auto-player-chip

# build for local and gimmie the terminal back
go build
./auto-player &
