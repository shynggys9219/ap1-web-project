package usecase

import "github.com/shynggys9219/ap1-web-project/internal/model"

type SnippetRepo interface {
	Create(title, content, expires string) (int, error)
	Get(id int) (*model.Snippet, error)
	Update(id int) (*model.Snippet, error)
	Delete(id int) (*model.Snippet, error)
	Latest() ([]*model.Snippet, error)
}
