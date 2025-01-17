package usecase

import "github.com/shynggys9219/ap1-web-project/internal/model"

type Snippet struct {
	snippetRepo SnippetRepo
}

func NewSnippet(snippetRepo SnippetRepo) *Snippet {
	return &Snippet{
		snippetRepo: snippetRepo,
	}
}

func (uc *Snippet) Create(title, content, expires string) (int, error) {
	return uc.snippetRepo.Create("test", "content", "1")
}

func (uc *Snippet) Get(id int) (*model.Snippet, error) {
	return nil, nil
}
func (uc *Snippet) Update(id int) (*model.Snippet, error) {
	return nil, nil
}
func (uc *Snippet) Delete(id int) (*model.Snippet, error) {
	return nil, nil
}
func (uc *Snippet) Latest() ([]*model.Snippet, error) {
	return nil, nil
}
