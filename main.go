package main

import (
	"log"
	"net/http"
	"time"
	"github.com/rs/cors"
	// hnd  "github.com/rishi-org-stack/pz/handler"
	ser "github.com/rishi-org-stack/pz/server"
)

func main() {
	router := ser.Route()
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8081"},
        AllowCredentials: true,
    })
	srv := &http.Server{
		Handler: c.Handler(router),
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
