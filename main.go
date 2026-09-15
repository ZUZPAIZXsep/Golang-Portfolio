package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title      string
	ActivePage string
}

// homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/index.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "Kantharakorn | Full Stack Developer",
		ActivePage: "home",
	}

	err = tmpl.ExecuteTemplate(w, "base", data)

}

// aboutpage
func aboutHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/about.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "About | Kantharakorn",
		ActivePage: "about",
	}

	tmpl.ExecuteTemplate(w, "base", data)
}

// educationpage
func educationHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/education.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "Education | Kantharakorn",
		ActivePage: "education",
	}

	tmpl.ExecuteTemplate(w, "base", data)
}

// experiencepage
func experienceHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/experience.html",
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:      "experience | Kantharakorn",
		ActivePage: "experience",
	}

	tmpl.ExecuteTemplate(w, "base", data)
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/education", educationHandler)
	http.HandleFunc("/experience", experienceHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	println("Server running at http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
