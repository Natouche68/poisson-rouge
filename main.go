package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Natouche68/poisson-rouge/api"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/", fs)

	r.HandleFunc("/api/notes", api.GetNotes)

	http.ListenAndServe(":5173", r)
}
