package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/trevorgrabham/cardcounting/cardcounting/html"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/handlers"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.New("index").ParseFiles(html.IncludeFiles["index"]...))
	if err := index.Execute(w, html.NewIndexData()); err != nil { panic(err) }
}

func StartServer() {
	http.Handle("/", handlers.SetCookieContext(http.HandlerFunc(HandleRoot)))
	http.Handle("/startPractice", handlers.SetCookieContext(http.HandlerFunc(handlers.HandleStartPractice)))

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}