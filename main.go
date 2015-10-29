package main

import (
	"io"
	"net/http"
	"github.com/fvbock/endless"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	endless.ListenAndServe(":80", nil)
}
