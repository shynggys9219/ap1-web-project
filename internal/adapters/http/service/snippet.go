package service

import (
	"fmt"
	"github.com/gorilla/mux"
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

func (s *Snippet) CreateSnippetForm(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateSnippetForm handler triggered")
	files := []string{
		"./ui/html/create.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (s *Snippet) CreateSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Use the r.PostForm.Get() method to retrieve the relevant data fields
	// from the r.PostForm map.
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	id, err := s.uc.Create(title, content, expires)
	if err != nil {
		http.Error(w, fmt.Sprintf("internal error"), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func UpdateSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updating code snippet"))
}

func DeleteSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleting code snippet"))
}

func (s *Snippet) GetSnippet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"]) // this pattern {id:[0-9]+} from routes guarantees id to be an integer
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
