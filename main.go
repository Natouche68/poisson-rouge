package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)

	http.ListenAndServe(":5173", nil)
}
