package main

import (
	handler "github.com/shynggys9219/ap1-web-project/internal/adapters/http"
	"github.com/shynggys9219/ap1-web-project/internal/app"
	"log"
	"net/http"
)

func main() {
	application := app.New()
	application.HttpServer.HandleFunc("/", handler.Home)
	application.HttpServer.HandleFunc("/snippet", handler.GetSnippet)
	application.HttpServer.HandleFunc("/snippet/create", handler.CreateSnippet)
	application.HttpServer.HandleFunc("/snippet/update", handler.UpdateSnippet)
	application.HttpServer.HandleFunc("/snippet/delete", handler.DeleteSnippet)

	application.HttpServer.Handle("/static/", http.StripPrefix("/static", application.FileServer))

	log.Println("starting the server on :9000")

	err := http.ListenAndServe(":9000", application.HttpServer)
	log.Fatalf("error occured during the server start: %v", err)
}
