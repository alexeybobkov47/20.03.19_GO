package main

import (
	"database/sql"
	"strconv"
	"testing"
)

// Цикломатическая сложность - 3
func TestNewPost(t *testing.T) {
	db, err := sql.Open("mysql", config.DBLink)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	s := Server{
		db: db,
	}
	post := Post{
		Header: "Тестовый заголовок",
		Text:   "Тест",
	}
	if err := newPost(s.db, post); err != nil {
		t.Error(err)
	}
}

// Цикломатическая сложность - 4
func TestGetPost(t *testing.T) {
	db, err := sql.Open("mysql", config.DBLink)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	s := Server{
		db: db,
	}
	type iD struct {
		id int
	}
	postID := iD{}
	row := s.db.QueryRow(`select id from Site.Post WHERE header = 'Тестовый заголовок'`)
	err = row.Scan(&postID.id)

	post := Post{
		ID:     postID.id,
		Header: "Тестовый заголовок",
		Text:   "Тест",
	}
	receivedpost, err := getPosts(s.db, strconv.Itoa(postID.id))
	if err != nil {
		t.Error(err)
	}
	if post != receivedpost {
		t.Error("Посты не совпадают")
	}

}

// Цикломатическая сложность - 3
func TestEditPost(t *testing.T) {
	db, err := sql.Open("mysql", config.DBLink)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	s := Server{
		db: db,
	}
	type iD struct {
		id int
	}
	postID := iD{}
	row := s.db.QueryRow(`select id from Site.Post WHERE header = 'Тестовый заголовок'`)
	err = row.Scan(&postID.id)
	updatepost := Post{
		Header: "Обновленный заголовок",
		Text:   "Обновленный тест",
	}
	err = editPost(s.db, updatepost, strconv.Itoa(postID.id))
	if err != nil {
		t.Error(err)
	}

}

// Цикломатическая сложность - 3
func TestDeletePost(t *testing.T) {
	db, err := sql.Open("mysql", config.DBLink)
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()

	s := Server{
		db: db,
	}
	type iD struct {
		id int
	}
	postID := iD{}
	row := s.db.QueryRow(`select id from Site.Post WHERE header = 'Обновленный заголовок'`)
	err = row.Scan(&postID.id)
	if err := deletePost(s.db, strconv.Itoa(postID.id)); err != nil {
		t.Error(err)
	}
}
