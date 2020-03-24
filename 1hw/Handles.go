package main

import (
	"log"
	"net/http"
	"strings"
)

// Цикломатическая сложность - 3
func (database *Server) showBlog(w http.ResponseWriter, r *http.Request) {
	blogs, err := getBlogs(database.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmplBlog.ExecuteTemplate(w, "blog", blogs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Цикломатическая сложность - 3
func (database *Server) showPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	posts, err := getPosts(database.db, (path[len(path)-1]))
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	if err := tmplPost.ExecuteTemplate(w, "post", posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Цикломатическая сложность - 4
func (database *Server) newPost(w http.ResponseWriter, r *http.Request) {
	newpost := Post{
		Header: r.FormValue("header"),
		Text:   r.FormValue("text"),
	}

	if len(newpost.Header) != 0 && len(newpost.Text) != 0 {
		err := newPost(database.db, newpost)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := tmplNewPost.ExecuteTemplate(w, "newpost", newpost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Цикломатическая сложность - 7
func (database *Server) editPost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	p := (path[len(path)-1])
	posts, err := getPosts(database.db, p)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	delPost := r.URL.Query()
	if delPost["delete"] != nil {
		err := deletePost(database.db, strings.Join(delPost["id"], ""))
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	editpost := Post{
		Header: r.FormValue("header"),
		Text:   r.FormValue("text"),
	}
	if len(editpost.Header) != 0 && len(editpost.Text) != 0 {
		err := editPost(database.db, editpost, p)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := tmplEditPost.ExecuteTemplate(w, "editpost", posts); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
