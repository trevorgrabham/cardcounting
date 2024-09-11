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
	if err := index.Execute(w, nil); err != nil { panic(err) }
}

var cookieMiddleWare = handlers.NewMiddleWare(handlers.WithCookies)
var cookiesAndHandsMiddleWare = handlers.NewMiddleWare(handlers.WithCookies, handlers.WithHands)

func StartServer() {
	http.Handle("/", cookieMiddleWare(http.HandlerFunc(HandleRoot)))
	http.Handle("/training/start", cookiesAndHandsMiddleWare(http.HandlerFunc(handlers.HandleStartTraining)))
	http.Handle("/training/split", cookiesAndHandsMiddleWare(http.HandlerFunc(handlers.HandleSplit)))
	http.Handle("/training/stand", cookiesAndHandsMiddleWare(http.HandlerFunc(handlers.HandleStand)))
	http.Handle("/training/hit", cookiesAndHandsMiddleWare(http.HandlerFunc(handlers.HandleHit)))
	http.Handle("/training/double", cookiesAndHandsMiddleWare(http.HandlerFunc(handlers.HandleDouble)))

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}