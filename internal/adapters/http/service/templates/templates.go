package templates

import "github.com/shynggys9219/ap1-web-project/internal/model"

type TemplateData struct {
	Snippet  *model.Snippet
	Snippets []*model.Snippet
}
