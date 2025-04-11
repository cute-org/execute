package main

import (
	"fmt"
	"log"
	"net/http"

	"execute/internal"
)

func main() {
	internal.InitDB()

	mux := http.NewServeMux()

	// Wrap handlers with ApplyMiddlewares or ApplyAuthMidlewares!
	mux.Handle("/register", internal.ApplyMiddlewares(http.HandlerFunc(internal.RegisterHandler)))
	mux.Handle("/login", internal.ApplyMiddlewares(http.HandlerFunc(internal.LoginHandler)))
	mux.Handle("/validate", internal.ApplyAuthMiddlewares(http.HandlerFunc(internal.ValidateHandler)))

	// v1
	muxWithPrefix := http.StripPrefix("/api/v1", mux)

	addr := ":8437"
	fmt.Printf("Server running at %s\n", addr)
	srv := &http.Server{
		Handler:      muxWithPrefix,
		Addr:         addr,
		ReadTimeout:  5e9,
		WriteTimeout: 10e9,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
