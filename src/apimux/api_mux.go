package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
steps to build and run:
	cd bin
	go build ../src/api/api_mux.go
 	./api_mux
	 http://localhost:10000/

	 --end-points available--

	 localhost:10000/
	 localhost:10000/articles
	 localhost:10000/article/1 { GET, POST and DELETE }

*/

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// article array for temp storage
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("endpoint: homePage")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpointt: getAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func getArticleById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpointt: getArticlebyId")
	vars := mux.Vars(r)
	article_id := vars["id"]
	w.WriteHeader(http.StatusOK)

	// fmt.Fprintf(w, "article id: "+article_id)
	// search article by id ( very unefficient way, just for tutorial )
	for _, article := range Articles {
		if article.Id == article_id {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpointt: createArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticleById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpointt: deleteArticleById")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func handleRequests() {
	// mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", getAllArticles)
	router.HandleFunc("/article/{id}", getArticleById)
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article/{id}", deleteArticleById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	log.Println("API running MUX Routers")
	log.Println("server is running at: http://localhost:10000/")

	Articles = []Article{
		Article{Id: "1", Title: "Title 1", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Title 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
