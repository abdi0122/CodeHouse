package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}
func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processFormData)
	
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processFormData(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	
	name := r.FormValue("name")
	location := r.FormValue("location")
	school := r.FormValue("school")
	year := r.FormValue("year")
	courses := r.FormValue("courses")
	interests := r.FormValue("interests")
	goals := r.FormValue("goals")
	linkedin := r.FormValue("linkedin")

	d := struct{
		Name string
		Location string
		School string
		Year   string
		Courses string
		Interests string
		Goals string
		LinkedIn string
	}{
		Name: name,
		Location: location,
		School: school,
		Year: year,
		Courses: courses,
		Interests: interests,
		Goals: goals,
		LinkedIn: linkedin,
	}

	tpl.ExecuteTemplate(w, "profile-data.html", d)
}