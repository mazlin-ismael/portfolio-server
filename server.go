package main

import (
	"fmt"
	"net/http"
	"time"
)

func newServer() *http.Server { 
	mux := http.NewServeMux()

	mux.HandleFunc("/contact", contactHandler)
	mux.HandleFunc("/register", registerHandler)

	srv := &http.Server {
		Addr: ":8080",
		Handler: mux,
		ReadHeaderTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout: 60 * time.Second,
	}
	return srv
}

func startServer(srv *http.Server) {
	fmt.Println("http://localhost:" + srv.Addr)
	panic(srv.ListenAndServe())
}