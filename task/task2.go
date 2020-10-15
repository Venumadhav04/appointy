
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)






func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Articles)
}

// Get single article using id
func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Looping through articles and find one with the id from the params
	for _, item := range Articles {
		if item.Name == params["Id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Article Article
	_ = json.NewDecoder(r.Body).Decode(&Article)
	Articles = append(Articles, Article)
	json.NewEncoder(w).Encode(Article)
}


func main() {
	// Init router
	r := mux.NewRouter()
	currentTime := time.Now()

	// Hardcoded data - @todo: add database
	Articles = append(Articles, Article{Title: "Titl1_1", Id = "1", subtitle = "jsdbfjksd", time = currentTime.String() ,content = "jfhvkjdjlkjfgklj"})
	Articles = append(Articles, Article{Title: "Titl1_2", Id = "2", subtitle = "jsdbfjksd", time = currentTime.String() ,content = "jfhvkjdjlkjfglj"})
	Articles = append(Articles, Article{Title: "Titl1_3", Id = "3", subtitle = "jsdbfjksd", time = currentTime.String() ,content = "jfhvkjdjlkjfgkj"})
	
	// Route handles & endpoints
	r.HandleFunc("/Articles", getArticles).Methods("GET")
	r.HandleFunc("/Articles/{Id}", getArticle).Methods("GET")
	r.HandleFunc("/Articles", createContact).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))
}