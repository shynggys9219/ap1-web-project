package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

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
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/create" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("creating code snippet"))
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

func GetSnippet(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "obtaining code snippet with ID=%v", id)
}
