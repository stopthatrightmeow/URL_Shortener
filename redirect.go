package main

import (
	"log"
	"net/http"
)

var Url string = "https://github.com"

func main(url) {
	http.Handle("/", http.RedirectHandler(Url, 301))
	log.Fatal(http.ListenAndServe(":9000", nil))
}
