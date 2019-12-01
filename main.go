package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vn-ki/tapchief-chall/handlers"
)

func main() {
	r := mux.NewRouter()
	indexTmpl := template.Must(template.ParseFiles("templates/index.gohtml", "templates/base.gohtml"))
	searchTmpl := template.Must(template.ParseFiles("templates/search.gohtml", "templates/base.gohtml"))

	r.HandleFunc("/", handlers.RootHandler)
	r.HandleFunc("/api/clear", handlers.ClearAPIHandler())
	r.HandleFunc("/api/index", handlers.IndexAPIHandler())
	r.HandleFunc("/api/search", handlers.SearchAPIHandler())
	r.HandleFunc("/clear", handlers.ClearHandler(indexTmpl))
	r.HandleFunc("/index", handlers.IndexHandler(indexTmpl))
	r.HandleFunc("/search", handlers.SearchHandler(searchTmpl))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
