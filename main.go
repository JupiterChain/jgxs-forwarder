package main

import (
	"log"
	"net/http"
	"os"
)

const endpoint = "https://beta-enterprise.jedtrade.com"

func main() {
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(port, nil))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, endpoint+r.URL.Path, http.StatusMovedPermanently)
}
