package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title: "Test title",Desc: "Test desc",Content: "Hello world"},
	}
	fmt.Println("EndPoint Hit: all Articles endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"HomePage Endpoint,Hit")
}

func handleRequests(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

// func main(){
// 	handleRequests()
// }