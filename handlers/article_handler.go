package handlers

import (
	"net/http"
	"path/filepath"
	"html/template"
	"github.com/HunterAC/blog/models"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	ap := filepath.Join("templates", "article.html")

	tmpl, err := template.ParseFiles(lp, ap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}

	tmpl.Execute(w, getArticle("Hunter's Article"))
}

func getArticle(title string) *models.Article {
	// Could get article from db or memory, etc.
	return &models.Article{
		Title: title,
		Content: "This is where the article content goes.",
		Author: "Hunter Allen",
		PubDay: 17,
		PubYear: 2023,
		PubMonth: 12,
	}
}