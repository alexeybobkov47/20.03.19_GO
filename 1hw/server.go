package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	tmplBlog     = template.Must(template.New("BlogTemplate").ParseFiles("index.html"))
	tmplPost     = template.Must(template.New("PostTemplate").ParseFiles("post.html"))
	tmplNewPost  = template.Must(template.New("PostTemplate").ParseFiles("newpost.html"))
	tmplEditPost = template.Must(template.New("PostTemplate").ParseFiles("editpost.html"))
	config       = new(Conf)
	// dsn  = "root:12345@tcp(192.168.0.39:3306)/Site?charset=utf8"
)

// Цикломатическая сложность - 3
func init() {
	bytes, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(bytes, config); err != nil {
		log.Fatal(err)
	}

}

// Цикломатическая сложность - 3
func main() {
	f, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Print("Тест")

	db, err := sql.Open("mysql", config.DBLink)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	s := Server{
		db: db,
	}

	router := http.NewServeMux()
	router.HandleFunc("/blog", s.showBlog)
	router.HandleFunc("/blog/newpost", s.newPost)
	router.HandleFunc("/blog/", s.showPost)
	router.HandleFunc("/blog/edit/", s.editPost)
	log.Printf("start listen on port %v", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))

}
