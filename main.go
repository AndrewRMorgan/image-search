package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

//var db *sql.DB
var err error

type Item struct {
	URL       string `json:"link"`
	Snippet   string `json:"snippet"`
	Image struct {
		Context   string `json:"contextLink"`
		Thumbnail string `json:"thumbnailLink"`
	} `json:"image"`
}

type GoogleAPIResponse struct {
	Items []Image `json:"items"`
}

type Config struct {
	API string
	Cx string
}

type History struct {
	Term string `json:"term"`
	When string `json:"when"`
}

file, _ := os.Open("config.json")
decoder := json.NewDecoder(file)
config := Config{}
err := decoder.Decode(&config)
check(err)

func main() {
	//databaseURI := config.db

	//db, err = sqp.Open("mysql", databaseURI)
	//check(err)
	//defer db.Close()

	//err = db.Ping()
	//check(err)

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

	safeQuery := url.QueryEscape(query)

	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s", config.API, config.Cx, safeQuery)

	request, err := http.NewRequest("GET", url, nil)
	check(err)

	client := &http.Client{}

	resp, err := client.Do(request)
	check(err)

	defer resp.Body.Close()
}

func getLatest(res http.ResponseWriter, req http.Request, _ httprouter.Params) {
	// err = db.QueryRow("SELECT term, when FROM history ORDER BY when LIMIT 10").Scan()
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

func getImages (body []byte) (*) {

}
