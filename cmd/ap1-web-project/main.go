package main

import (
	"context"
	"github.com/shynggys9219/ap1-web-project/internal/adapters/http"
	"github.com/shynggys9219/ap1-web-project/internal/adapters/postgres"
	"github.com/shynggys9219/ap1-web-project/internal/app"
	"github.com/shynggys9219/ap1-web-project/internal/usecase"
	"github.com/shynggys9219/ap1-web-project/pkg"
	"log"
)

func main() {
	ctx := context.Background()
	db, err := pkg.NewDB(ctx, pkg.Config{
		Database: "snippetbox",
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
	})
	defer db.Conn.Close(ctx)
	if err != nil {
		panic(err)
	}
	snippetRepo := postgres.NewSnippet(db.Conn)
	snippetUsecase := usecase.NewSnippet(snippetRepo)

	server := http.NewSimpleServer(snippetUsecase)
	application := app.New(server)

	log.Println("starting the server on :9000")

	application.SimpleServer.Run(ctx)
	log.Fatalf("error occured during the server start: %v", err)
}
