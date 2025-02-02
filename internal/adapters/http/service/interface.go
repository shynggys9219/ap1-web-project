package service

import "github.com/shynggys9219/ap1-web-project/internal/model"

type SnippetUsecase interface {
	Create(title, content, expires string) (int, error)
	Get(id int) (*model.Snippet, error)
	Update(id int) (*model.Snippet, error)
	Delete(id int) (*model.Snippet, error)
	Latest() ([]*model.Snippet, error)
}

type AuthUsecase interface {
	Create(user model.User) (int, error)
	Get(id int) (*model.User, error)
	Update(id int) (*model.User, error)
	Delete(id int) error
}
