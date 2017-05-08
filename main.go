package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type ImageJson struct {
	URL       string `json:"url"`
	Snippet   string `json:"snippet"`
	Thumbnail string `json:"thumbnail"`
	Context   string `json:"context"`
}

type Config struct {
	API string
	Cx string
}

file, _ := os.Open("config.json")
decoder := json.NewDecoder(file)
config := Config{}
err := decoder.Decode(&config)
check(err)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	router.GET("/api/imagesearch/:queries", getQuery)
	router.GET("/api/latest/imagesearch", getLatest)
	router.GET("/", index)
	http.ListenAndServe(":"+port, router)
}

//API Url: "https://www.googleapis.com/customsearch/v1?key=" + config.API + "&cx=" + config.Cx "&q=" + query + ""

func getQuery(res http.ResponseWriter, req http.Request, ps httprouter.Params) {
  query := ps.ByName("queries")

	request, err := http.NewRequest("GET", url, nil)
	check(err)

	client := &http.Client{}

	resp, err := client.Do(request)
	check(err)

	defer resp.Body.Close()


}

func getLatest(res http.ResponseWriter, req http.Request, _ httprouter.Params) {

}

func index(res http.ResponseWriter, req http.Request, _ httprouter.Params) {
  http.ServeFile(res, req, "./static/index.html")
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func makeList(img) {
	return {
		"url": img.url,
		"snippet": img.title,
		"thumbnail": img.thumbnail.url,
		"context": img.sourceUrl
	}
}
