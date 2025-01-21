package app

import (
	"github.com/shynggys9219/ap1-web-project/internal/adapters/http"
)

const (
	addr = ":9000"
)

type App struct {
	SimpleServer *http.SimpleServer
}

func New(server *http.SimpleServer) App {
	return App{
		SimpleServer: server,
	}
}
