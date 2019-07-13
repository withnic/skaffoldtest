package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "OK"); err != nil {
		log.Print("error")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println(http.ListenAndServe(":8080", nil))
}
