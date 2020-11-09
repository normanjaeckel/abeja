package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":9000", http.HandlerFunc(healthHandler)))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {}
