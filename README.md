# Auto random music player

Howdy! Here's a silly little program that I wrote to build an mp3 player for my upcoming nursery. This is intended to be put on a headless little linux computer. Mine is running on a C.H.I.P. It's kinda like a network conrtolled iPod shuffle. Anyone remember those?

## What does it do?

Not much! On boot it expects you to have `mpg123` installed, and a folder full of mp3s. This is currently hardcoded to `/music`, but I'll likely make that configurable in the future.

So yeah, load up some tunes, and turn this program on.

## How to run it?

Well, if you build this for arm (check out the `build.sh` script) you can put it on a C.H.I.P. or raspberry pi. Then you can have the program run automatically. The easiest way to do that is to have it auto restart via cron. Something like this:

`@reboot /root/auto-player-chip`

## Wow, I wish I could skip this track

You can! Included in this software is a webserver. It's ugly and utilitarian, but if you start this software up, and your device is configured to get an IP address on your local network, then pop open a browser and visit `http://{YOUR.DEVICE.IP.ADDRESS}:8899`. You should see a big ol `<h1>` with the filename of the current mp3 playing, and a link to skip and go to the next song.

## Roadmap

- System check to help you configure the device if it's not finding everything it needs to run
- Better Web UI with more controls (previous track, volume up/down, etc)
- A web page to upload new mp3s
- A web page to show statistics (hard drive usage, uptime, etc)
- A link to skip and delete the current mp3 in case you start to get sick of a certain song
- Configurable web port and music folder locations

## Possible features (if anyone cares?)

- Swap out mpg123 and .mp3s for something similar using open formats (cvlc and .ogm?)
- Another cool idea? Open an issue and let's chat about it!
