package service

import (
	"fmt"
	"github.com/shynggys9219/ap1-web-project/internal/adapters/http/service/templates"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Snippet struct {
	uc SnippetUsecase
}

func NewSnippet(uc SnippetUsecase) *Snippet {
	return &Snippet{
		uc: uc,
	}
}

func (s *Snippet) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	snippets, err := s.uc.Latest()
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	data := templates.TemplateData{Snippets: snippets}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (s *Snippet) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/create" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	id, err := s.uc.Create(title, content, expires)
	if err != nil {
		http.Error(w, fmt.Sprintf("internal error"), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func UpdateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/update" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("updating code snippet"))
}

func DeleteSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/delete" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("deleting code snippet"))
}

func (s *Snippet) GetSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	snippet, err := s.uc.Get(id)
	if err != nil {
		fmt.Fprintf(w, "%v", err.Error())
		return
	}
	data := templates.TemplateData{Snippet: snippet}
	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return

	}
	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
