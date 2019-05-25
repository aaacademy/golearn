package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Judul", Desc: "Lorem ipsum lorem ipsum, Lorem ipsum...", Content: "Ini Isi seluruh dari lorem ipsum lorem ipsum, dan lorem ipsum"},
		Article{Title: "Judul", Desc: "Lorem ipsum lorem ipsum, Lorem ipsum...", Content: "Ini Isi seluruh dari lorem ipsum lorem ipsum, dan lorem ipsum"},
		Article{Title: "Judul", Desc: "Lorem ipsum lorem ipsum, Lorem ipsum...", Content: "Ini Isi seluruh dari lorem ipsum lorem ipsum, dan lorem ipsum"},
		Article{Title: "Judul", Desc: "Lorem ipsum lorem ipsum, Lorem ipsum...", Content: "Ini Isi seluruh dari lorem ipsum lorem ipsum, dan lorem ipsum"},
		Article{Title: "Judul", Desc: "Lorem ipsum lorem ipsum, Lorem ipsum...", Content: "Ini Isi seluruh dari lorem ipsum lorem ipsum, dan lorem ipsum"},
	}

	fmt.Println("End point Hit")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage ")
}

func homePage2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage 2")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", getAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
