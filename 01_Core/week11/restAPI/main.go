package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Article is a way to define articles to show to the user
type Article struct {
	ID      string
	Title   string
	Desc    string
	Content string
}

//Articles is just an array of articles
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Enpoint Hit")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "returnAllArticle endpoint hit\n")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "returnSingleArticle endpoint hit\n")
	vars := mux.Vars(r)
	key := vars["id"]
	//Get article specified by route
	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for i, article := range Articles {
		if article.ID == key {
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	for i, el := range Articles {
		if el.ID == key {
			Articles[i] = article
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", returnAllArticles)
	router.HandleFunc("/articles/{id}", returnSingleArticle)
	router.HandleFunc("/article", createNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Rest API started")
	Articles = []Article{
		Article{ID: "1", Title: "Hello", Desc: "Hello from the hello  article", Content: "HI!"},
		Article{ID: "2", Title: "World", Desc: "World from the world article", Content: "The world is almost 10 feet long!"},
	}

	handleRequests()
}
