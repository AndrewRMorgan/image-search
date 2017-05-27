package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var (
	//db *sql.DB
	err    error
	config Config
)

type GoogleAPIResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
	Image   struct {
		ContextLink   string `json:"contextLink"`
		ThumbnailLink string `json:"thumbnailLink"`
	} `json:"image"`
}

type ImageList struct {
	Images []Image
}

type Image struct {
	URL       string `json:"url"`
	Snippet   string `json:"snippet"`
	Thumbnail string `json:"thumbnail"`
	Context   string `json:"context"`
}

type Config struct {
	API string
	Cx  string
}

type History struct {
	Term string `json:"term"`
	When string `json:"when"`
}

func main() {
	file, err := os.Open("config.json")
	check(err)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	check(err)

	//TODO: History function.

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

func getQuery(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	query := ps.ByName("queries")

	queryValues := req.URL.Query()
	offset := queryValues.Get("offset")

	safeQuery := url.QueryEscape(query)

	var url string

	if offset != "" {
		url = fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s&start=%s&searchType=image", config.API, config.Cx, safeQuery, offset)
	} else {
		url = fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s&searchType=image", config.API, config.Cx, safeQuery)
	}

	request, err := http.NewRequest("GET", url, nil)
	check(err)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(request)
	check(err)

	defer resp.Body.Close()

	var apiResp GoogleAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	check(err)

	var imageList = addImages(apiResp.Items)

	js, err := json.Marshal(imageList.Images)
	check(err)
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func getLatest(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// err = db.QueryRow("SELECT term, when FROM history ORDER BY when LIMIT 10").Scan()
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(res, req, "./static/index.html")
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func addImages(dataArr []Item) ImageList {
	var imageList ImageList

	for _, elem := range dataArr {
		imageList.Images = append(imageList.Images, Image{
			URL:       elem.Link,
			Snippet:   elem.Snippet,
			Thumbnail: elem.Image.ThumbnailLink,
			Context:   elem.Image.ContextLink,
		})
	}

	return imageList
}
