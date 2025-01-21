package http

import (
	"context"
	"github.com/shynggys9219/ap1-web-project/internal/adapters/http/service"
	"net/http"
)

type SimpleServer struct {
	srv            *http.ServeMux
	snippetUsecase *service.Snippet
}

func NewSimpleServer(usecase SnippetUsecase) *SimpleServer {
	s := http.NewServeMux()
	uc := service.NewSnippet(usecase)
	s.HandleFunc("/", uc.Home)
	s.HandleFunc("/snippet", uc.GetSnippet)
	s.HandleFunc("/snippet/create", uc.CreateSnippet)
	s.HandleFunc("/snippet/update", service.UpdateSnippet)
	s.HandleFunc("/snippet/delete", service.DeleteSnippet)

	s.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))
	return &SimpleServer{
		srv:            s,
		snippetUsecase: uc,
	}
}

func (s *SimpleServer) Run(ctx context.Context) {
	err := http.ListenAndServe(":9000", s.srv)
	panic(err)
}
