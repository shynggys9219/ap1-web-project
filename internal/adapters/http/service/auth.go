package service

import (
	"github.com/golangcollege/sessions"
	"net/http"
)

type Auth struct {
	uc      AuthUsecase
	session *sessions.Session
}

func NewAuth(uc AuthUsecase, session *sessions.Session) *Auth {
	return &Auth{
		uc:      uc,
		session: session,
	}
}

func (a *Auth) Create(w http.ResponseWriter, r *http.Request) {

}

func (a *Auth) Get(w http.ResponseWriter, r *http.Request) {}

func (a *Auth) GetSignUpForm(w http.ResponseWriter, r *http.Request) {

}

func (a *Auth) GetLoginForm(w http.ResponseWriter, r *http.Request) {}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {}

func (a *Auth) Logout(w http.ResponseWriter, r *http.Request) {}

func (a *Auth) Update(w http.ResponseWriter, r *http.Request) {}

func (a *Auth) Delete(w http.ResponseWriter, r *http.Request) {}
