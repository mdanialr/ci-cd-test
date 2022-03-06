package main

import (
	"log"
	"net/http"

	"github.com/mdanialr/ci-cd-test/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/print/", handlers.PrintName)
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatalln("Failed to serve app:", err)
	}
}
