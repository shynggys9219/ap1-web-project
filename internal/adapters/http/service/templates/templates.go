package templates

import (
	"github.com/shynggys9219/ap1-web-project/internal/model"
	"net/url"
)

type TemplateData struct {
	FormData   url.Values
	FormErrors map[string]string
	Snippet    *model.Snippet
	Snippets   []*model.Snippet
	Flash      string
}
