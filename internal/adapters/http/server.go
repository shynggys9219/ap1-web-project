package http

import (
	"context"
	"github.com/golangcollege/sessions"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/shynggys9219/ap1-web-project/internal/adapters/http/service"
	"net/http"
	"time"
)

type SimpleServer struct {
	srv            http.Handler
	snippetHandler *service.Snippet
	authHandler    *service.Auth
}

func NewSimpleServer(snippet SnippetUsecase, auth AuthUsecase) *SimpleServer {
	stdMiddleware := alice.New(recoverPanic, logRequest, secureHeaders)
	secretKey := "session_secret_key"
	session := sessions.New([]byte(secretKey))
	session.Lifetime = 12 * time.Hour

	snippetHandler := service.NewSnippet(snippet, session)
	authHandler := service.NewAuth(auth, session)

	dynamicMiddleware := alice.New(session.Enable)

	s := mux.NewRouter()
	// Snippet routes
	s.Handle("/", dynamicMiddleware.ThenFunc(snippetHandler.Home)).Methods(http.MethodGet)
	// 1. {id:[0-9]+} guarantee of integer
	// 2. gorilla evaluates routes in specific order, this was the problem in the lecture
	s.Handle("/snippet/{id:[0-9]+}", dynamicMiddleware.ThenFunc(snippetHandler.GetSnippet)).Methods(http.MethodGet)
	s.Handle("/snippet/create", dynamicMiddleware.ThenFunc(snippetHandler.CreateSnippet)).Methods(http.MethodPost)
	s.Handle("/snippet/create", dynamicMiddleware.ThenFunc(snippetHandler.CreateSnippetForm)).Methods(http.MethodGet)
	s.HandleFunc("/snippet/update", service.UpdateSnippet).Methods(http.MethodPut)
	s.HandleFunc("/snippet/delete", service.DeleteSnippet).Methods(http.MethodDelete)

	// Auth routes
	s.Handle("/user/signup", dynamicMiddleware.ThenFunc(authHandler.GetSignUpForm)).Methods(http.MethodGet)
	s.Handle("/user/signup", dynamicMiddleware.ThenFunc(authHandler.Create)).Methods(http.MethodPost)
	s.Handle("/user/login", dynamicMiddleware.ThenFunc(authHandler.GetLoginForm)).Methods(http.MethodGet)
	s.Handle("/user/login", dynamicMiddleware.ThenFunc(authHandler.Login)).Methods(http.MethodPost)
	s.Handle("/user/logout", dynamicMiddleware.ThenFunc(authHandler.Logout)).Methods(http.MethodPost)

	s.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return &SimpleServer{
		srv:            stdMiddleware.Then(s),
		snippetHandler: snippetHandler,
		authHandler:    authHandler,
	}
}

func (s *SimpleServer) Run(ctx context.Context) {
	err := http.ListenAndServe(":9000", s.srv)
	panic(err)
}
