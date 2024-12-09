package main

import (
	"log"
	"net/http"
	"time"

	"gcommons/handler"
)

func main() {
	http.HandleFunc("/healthz", handler.HealthCheck(time.Now(), "User"))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
