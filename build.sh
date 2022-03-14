#! /usr/bin/bash

# kill any running processes
killall mpg123
killall auto-player

# build for CHIP
GOOS=linux GOARCH=arm go build -o auto-player-chip

# build for local and gimmie the terminal back
go build
./auto-player &
