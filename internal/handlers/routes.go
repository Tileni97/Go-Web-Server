package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"web-server/internal/models"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", loggingMiddleware(http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/submit", submitHandler)
	mux.HandleFunc("/greet", greetHandler)

	return mux
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(filepath.Join("templates", "greet.html"))
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := models.GreetingData{
		Name:    r.FormValue("name"),
		Message: r.FormValue("message"),
	}

	tmpl.Execute(w, data)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r,
		"/greet?name="+r.FormValue("name")+"&message="+r.FormValue("message"),
		http.StatusSeeOther)
}
