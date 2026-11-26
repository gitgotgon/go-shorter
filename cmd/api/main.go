package main

import (
	"log"
	"net/http"

	apphttp "github.com/gitgotgon/go-shorter/internal/http"
)

func main() {
	router := apphttp.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Printf("Starting server on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
