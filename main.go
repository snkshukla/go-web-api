package main

import (
  "fmt"
	"github.com/gorilla/mux"
	"log"
  "net/http"
  "encoding/json"
)

type Req struct {
  URL string `json:"url"`
  Headers http.Header `json:"headers"`
}

func handleError(err error) {
  	if err != nil {
		log.Println("An error occurred: \n", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
  var req = new(Req)
  req.URL = fmt.Sprintf("%s%s", r.Host, r.URL)
  req.Headers = r.Header

  log.Printf("Recieved a request with following Parameters:")
  log.Printf("URL: %s%s", req.URL)
  log.Println("Headers: ", req.Headers)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(req)
  fmt.Println("\n")
}
