package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./html/base.html",
		"./html/hello.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "title", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not defined in the enivronment")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/hello", home)

	log.Printf("Starting server on :%v\n", portString)
	err := http.ListenAndServe(":"+portString, r)
	if err != nil {
		log.Fatal(err)
	}
}
