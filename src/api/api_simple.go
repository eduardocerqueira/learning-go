package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
steps to build and run:
	cd bin
	go build ../src/api/api_simple.go
	 ./api_simple
	 http://localhost:10000/

	 --end-points available--

	 localhost:10000/
	 localhost:10000/articles

*/

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// article array for temp storage
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
	w.Header().Set("Content-Type", "application/json")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint: getAllArticles")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(generateArticles())
}

func generateArticles() []Article {
	Articles := []Article{
		Article{Title: "Hello 1", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	return Articles
}

func handleRequests() {
	// route home
	http.HandleFunc("/", homePage)
	// route for articles
	http.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	log.Println("server is running at: http://localhost:10000/")
	handleRequests()
}
