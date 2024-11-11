package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Natouche68/poisson-rouge/db"
)

var getNotesTmpl = template.Must(template.ParseFiles("templates/getNotes.html"))

func GetNotes(w http.ResponseWriter, r *http.Request) {
	database, err := db.GetDB()
	if err != nil {
		fmt.Println(err)
	}

	var notes []db.Note
	database.Order("created_at desc").Find(&notes)

	getNotesTmpl.Execute(w, struct{ Notes []db.Note }{Notes: notes})
}
