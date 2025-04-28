package main

import (
	"log"
	"net/http"
	"time"

	"execute/internal"
	"execute/internal/handlers/auth"
	"execute/internal/handlers/group"
	"execute/internal/handlers/task"
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

	// AUTH
	mux.Handle("/register", middleware.ApplyMiddlewares(http.HandlerFunc(auth.RegisterHandler)))
	mux.Handle("/login", middleware.ApplyMiddlewares(http.HandlerFunc(auth.LoginHandler)))
	mux.Handle("/validate", middleware.ApplyAuthMiddlewares(http.HandlerFunc(auth.ValidateHandler)))

	// USER
	mux.Handle("/user", middleware.ApplyAuthMiddlewares(middleware.Router(map[string]http.HandlerFunc{
		"GET": user.UsersHandler,
		"PUT": user.EditUserHandler,
	})))
	mux.Handle("/avatar", middleware.ApplyAuthMiddlewares(http.HandlerFunc(user.ServeAvatarHandler)))
	mux.Handle("/user/current", middleware.ApplyAuthMiddlewares(http.HandlerFunc(user.UserProfileHandler)))

	// GROUP
	mux.Handle("/group", middleware.ApplyAuthMiddlewares(middleware.Router(map[string]http.HandlerFunc{
		"GET":  user.GroupUsersHanlder,
		"POST": group.CreateGroupHandler,
		"PUT":  group.UpdateGroupHandler,
	})))
	mux.Handle("/group/join", middleware.ApplyAuthMiddlewares(http.HandlerFunc(group.JoinGroupHandler)))

	// TASK
	mux.Handle("/task", middleware.ApplyAuthMiddlewares(middleware.Router(map[string]http.HandlerFunc{
		"GET":  task.ListTasksHandler,
		"POST": task.CreateTaskHandler,
		"PUT":  task.UpdateTaskHandler,
	})))

	// v1
	muxWithPrefix := http.StripPrefix("/api/v1", mux)

	utils.PrintIPs(addr)

	srv := &http.Server{
		Handler:      middleware.CorsMiddleware(muxWithPrefix),
		Addr:         addr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

