package main

import (
	"log"
	"net/http"
	"html/template"
	"path/filepath"
	"github.com/HunterAC/blog/handlers"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", "index.html")

	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/article", handlers.ArticleHandler)
	
	log.Print("listening on port :8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
