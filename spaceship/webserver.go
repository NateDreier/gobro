package main

import (
	"code/gobro/spaceship/hlogger"
	"fmt"
	"net/http"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting mah spaceship")
	// http.HandlerFunc("/", sroot)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to mah spaceship")
	})
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprintf(w, "Welcome to mah spaceship")

	logger.Println("rcvd thing")
}
