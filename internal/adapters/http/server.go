package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/shynggys9219/ap1-web-project/internal/adapters/http/service"
	"net/http"
)

type SimpleServer struct {
	srv            http.Handler
	snippetUsecase *service.Snippet
}

func NewSimpleServer(usecase SnippetUsecase) *SimpleServer {
	stdMiddleware := alice.New(recoverPanic, logRequest, secureHeaders)
	uc := service.NewSnippet(usecase)

	s := mux.NewRouter()
	s.HandleFunc("/", uc.Home).Methods(http.MethodGet)
	// 1. {id:[0-9]+} guarantee of integer
	// 2. gorilla evaluates routes in specific order, this was the problem in the lecture
	s.HandleFunc("/snippet/{id:[0-9]+}", uc.GetSnippet).Methods(http.MethodGet)
	s.HandleFunc("/snippet/create", uc.CreateSnippet).Methods(http.MethodPost)
	s.HandleFunc("/snippet/create", uc.CreateSnippetForm).Methods(http.MethodGet)
	s.HandleFunc("/snippet/update", service.UpdateSnippet).Methods(http.MethodPut)
	s.HandleFunc("/snippet/delete", service.DeleteSnippet).Methods(http.MethodDelete)

	s.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return &SimpleServer{
		srv:            stdMiddleware.Then(s),
		snippetUsecase: uc,
	}
}

func (s *SimpleServer) Run(ctx context.Context) {
	err := http.ListenAndServe(":9000", s.srv)
	panic(err)
}
