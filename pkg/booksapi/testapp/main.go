package main

// BFT for one of the API endpoints
import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/dedobbin/letterbooked/pkg/booksapi"
)

type PageData struct {
	Author booksapi.Author
	Books  []booksapi.Document
}

func main() {
	tmpl := template.Must(template.ParseFiles("authorInfo.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mPath := filepath.Base(r.URL.String())
		url.QueryUnescape(mPath)
		fmt.Println(mPath)
		res, err := booksapi.QueryAuthorInformation(mPath)
		if err != nil || len(res.Docs) == 0 {
			fmt.Fprintf(w, "<!DOCTYPE html><head><title>404</title></head><body><h1>No info found!</h1></body>")
			return
		}
		res2, err := booksapi.QueryBooksByAuthor(res.Docs[0].Name, 5)
		if err != nil {
			fmt.Fprintf(w, "<!DOCTYPE html><head><title>404</title></head><body><h1>No info found!</h1></body>")
			return
		}
		data := PageData{Author: res.Docs[0], Books: res2.Docs}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
