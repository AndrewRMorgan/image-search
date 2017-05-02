package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "os"

  "github.com/julienschmidt/httprouter"
)

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

func getQuery(res http.ResponseWriter, req http.Request, ps httprouter.Params) {

}

func getLatest(res http.ResponseWriter, req http.Request, ps httprouter.Params) {

}

func index(res http.ResponseWriter, req http.Request, ps httprouter.Params) {

}
