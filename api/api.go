package api

import (
	"html/template"
	"net/http"
)

var getNotesTmpl = template.Must(template.ParseFiles("templates/getNotes.html"))

func GetNotes(w http.ResponseWriter, r *http.Request) {
	getNotesTmpl.Execute(w, nil)
}
