package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/Natouche68/poisson-rouge/api"
	"github.com/Natouche68/poisson-rouge/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db.InitDB()

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/", fs)

	r.HandleFunc("/api/notes", api.GetNotes)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
