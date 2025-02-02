package service

import (
	"fmt"
	"github.com/golangcollege/sessions"
	"github.com/gorilla/mux"
	"github.com/shynggys9219/ap1-web-project/internal/adapters/http/service/templates"
	"github.com/shynggys9219/ap1-web-project/internal/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Snippet struct {
	uc      SnippetUsecase
	session *sessions.Session
}

func NewSnippet(uc SnippetUsecase, session *sessions.Session) *Snippet {
	return &Snippet{
		uc:      uc,
		session: session,
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

	errorsMap := make(map[string]string)

	expiry, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}
	snippet := model.Snippet{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Created: time.Time{},
		Expires: time.Now().UTC().Add(time.Duration(24*expiry) * time.Hour),
	}

	snippet.Validate(errorsMap)
	if len(errorsMap) > 0 {
		files := []string{
			"./ui/html/create.page.tmpl",
			"./ui/html/base.layout.tmpl",
		}

		ts, templateError := template.ParseFiles(files...)
		if templateError != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(
			w,
			&templates.TemplateData{
				FormErrors: errorsMap,
				FormData:   r.PostForm,
			})
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}

		return
	}

	id, err := s.uc.Create(snippet.Title, snippet.Content, r.PostForm.Get("expires"))
	if err != nil {
		http.Error(w, fmt.Sprintf("internal error"), http.StatusInternalServerError)

		return
	}

	s.session.Put(r, "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
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
	flash := s.session.PopString(r, "flash")
	data := templates.TemplateData{
		Snippet: snippet,
		Flash:   flash,
	}
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
