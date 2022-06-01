package main

import (
	_ "embed"
	"net/http"
	"strings"
)

//go:embed index.html
var html string

//go:embed kiddo.svg
var kiddo string

//go:embed parents.img
var parents string

//go:embed vue.js
var vuejs string

//go:embed style.css
var stylecss string

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// start with the embeded HTML
	output := html

	output = strings.ReplaceAll(output, "%%KIDDO%%", kiddo)
	output = strings.ReplaceAll(output, "%%PARENTS%%", parents)
	output = strings.ReplaceAll(output, "/*STYLECSS*/", stylecss)
	output = strings.ReplaceAll(output, "/*VUEJS*/", vuejs)

	w.Write([]byte(output))
}
