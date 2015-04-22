package main

import (
	"log"
	"net/http"
)

func main() {
	var handler Handler

	ReadConfigs(&handler)

	s := &http.Server{
		Addr:    handler.Config.Listen,
		Handler: handler,
	}

	log.Println("listening on", handler.Config.Listen)
	log.Fatal(s.ListenAndServe())
}
