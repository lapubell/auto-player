package main

import (
	"net/http"
	"strings"
)

var html string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>

	<style>
	body {background: black;}
	* {color: white;}
	</style>
</head>
<body>

%%CURRENT%%
<a href='/skip'>Next Track</a>

</body>
</html>
`

func handleIndex(w http.ResponseWriter, r *http.Request) {
	output := strings.ReplaceAll(html, "%%CURRENT%%", "<h1>Now Playing: "+currentSong+"</h1>")
	w.Write([]byte(output))
}
