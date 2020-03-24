package main

import "database/sql"

// Blog - Структура блога
type Blog struct {
	ID          int
	Name        string
	Description string
	Posts       []Post
}

//Post - Структура постов в блоге
type Post struct {
	ID     int
	Header string
	Text   string
}

// Server -
type Server struct {
	db *sql.DB
}

// Conf - Конфигурация
type Conf struct {
	DBLink  string `json:"db_link" check:"required"`
	LogFile string `json:"log_file" check:"required"`
	Port    string `json:"port" check:"required"`
}
