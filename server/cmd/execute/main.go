package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"execute/internal"
	"execute/internal/handlers/auth"
	"execute/internal/handlers/user"
	"execute/internal/middleware"
)

func main() {
	internal.InitDB()
	go auth.CleanupExpiredSessions(10 * time.Minute)

	mux := http.NewServeMux()

	// Wrap handlers with ApplyMiddlewares or ApplyAuthMidlewares!
	mux.Handle("/register", middleware.ApplyMiddlewares(http.HandlerFunc(auth.RegisterHandler)))
	mux.Handle("/login", middleware.ApplyMiddlewares(http.HandlerFunc(auth.LoginHandler)))
	mux.Handle("/validate", middleware.ApplyAuthMiddlewares(http.HandlerFunc(auth.ValidateHandler)))
	mux.Handle("/users", middleware.ApplyAuthMiddlewares(http.HandlerFunc(user.UsersHandler)))

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
