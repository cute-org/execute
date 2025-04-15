package main

import (
	"log"
	"net/http"
	"time"

	"execute/internal"
	"execute/internal/handlers/auth"
	"execute/internal/handlers/user"
	"execute/internal/middleware"
	"execute/internal/utils"
)

const addr = ":8437"

func main() {
	internal.InitDB()
	go auth.CleanupExpiredSessions(10 * time.Minute)

	mux := http.NewServeMux()

	// Wrap handlers with ApplyMiddlewares or ApplyAuthMidlewares!
	mux.Handle("/register", middleware.ApplyMiddlewares(http.HandlerFunc(auth.RegisterHandler)))
	mux.Handle("/login", middleware.ApplyMiddlewares(http.HandlerFunc(auth.LoginHandler)))
	mux.Handle("/validate", middleware.ApplyAuthMiddlewares(http.HandlerFunc(auth.ValidateHandler)))
	mux.Handle("/users", middleware.ApplyAuthMiddlewares(http.HandlerFunc(user.UsersHandler)))
	mux.Handle("/user-edit", middleware.ApplyAuthMiddlewares(http.HandlerFunc(user.EditUserHandler)))
	mux.Handle("/avatar", middleware.ApplyImageMiddlewares(http.HandlerFunc(user.ServeAvatarHandler)))

	// v1
	muxWithPrefix := http.StripPrefix("/api/v1", mux)

	utils.PrintIPs(addr)

	srv := &http.Server{
		Handler:      muxWithPrefix,
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
