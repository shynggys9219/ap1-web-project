package service

import (
	"context"
	handler "github.com/shynggys9219/ap1-web-project/internal/adapters/service/http"
	"net/http"
)

type SimpleServer struct {
	srv            *http.ServeMux
	snippetUsecase SnippetUsecase
}

func NewSimpleServer(usecase SnippetUsecase) *SimpleServer {
	s := http.NewServeMux()
	s.HandleFunc("/", handler.Home)
	s.HandleFunc("/snippet", handler.GetSnippet)
	s.HandleFunc("/snippet/create", handler.CreateSnippet)
	s.HandleFunc("/snippet/update", handler.UpdateSnippet)
	s.HandleFunc("/snippet/delete", handler.DeleteSnippet)

	s.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))
	return &SimpleServer{
		srv:            s,
		snippetUsecase: usecase,
	}
}

func (s *SimpleServer) Run(ctx context.Context) {
	err := http.ListenAndServe(":9000", s.srv)
	panic(err)
}
