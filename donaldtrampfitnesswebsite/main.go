package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", idx)
	http.Handle("/assets/",http.StripPrefix("/assets",http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
	var first string
	if req.Method == http.MethodPost {
		first=req.FormValue("new-word")
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", first)
	if err != nil {
		log.Fatal(err)
	}
}
//func ee(w http.ResponseWriter, req *http.Request) {
//}
