package http

import (
	"github.com/shynggys9219/ap1-web-project/internal/adapters/http/service"
)

type SnippetUsecase interface {
	service.SnippetUsecase
}

type AuthUsecase interface {
	service.AuthUsecase
}
