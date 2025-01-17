package app

import (
	"github.com/shynggys9219/ap1-web-project/internal/adapters/service"
)

const (
	addr = ":9000"
)

type App struct {
	SimpleServer *service.SimpleServer
}

func New(server *service.SimpleServer) App {
	return App{
		SimpleServer: server,
	}
}
