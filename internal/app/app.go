package app

import "net/http"

const (
	addr = ":9000"
)

type App struct {
	HttpServer *http.ServeMux
	FileServer http.Handler
}

func New() App {
	return App{
		HttpServer: http.NewServeMux(),
		FileServer: http.FileServer(http.Dir("./ui/static/")),
	}
}
