package api

import (
	"html/template"
	"net/http"
)

var getNotesTmpl = template.Must(template.ParseFiles("templates/getNotes.html"))

func GetNotes(w http.ResponseWriter, r *http.Request) {
	getNotesTmpl.Execute(w, struct{ Notes []string }{Notes: []string{"Note 1", "Note 2", "Note 3", "Note 4", "Note 5", "Note 6"}})
}
