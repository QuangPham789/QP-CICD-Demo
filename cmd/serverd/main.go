package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/handler/rest"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/config"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/db"
	"github.com/s3corp-github/SP_FriendManagementAPI_QuangPham/internal/pkg/utils"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbConn, err := db.ConnectDB(config.DBSource)
	if err != nil {
		log.Println("Connection database fail with error: ", err)
	}

	router := initRouter(dbConn)

	// Start server with port 8080
	log.Println("Server start at port 8088")
	if err := http.ListenAndServe(":8088", router); err != nil {
		log.Println("Error start server with port 8088", err)
	}
}

func initRouter(dbConn *sql.DB) *chi.Mux {
	handler := rest.NewHandler(dbConn)
	router := chi.NewRouter()
	router.Use(utils.LogRequest)

	router.Route("/", func(r chi.Router) {
		r.Route("/users", usersRouter(handler))
		r.Route("/friends", friendsRouter(handler))
	})

	return router
}

func usersRouter(h rest.Handler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Get("/", h.GetUsers)
	}
}

func friendsRouter(h rest.Handler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/create", h.CreateFriend)
		r.Post("/subscription", h.CreateSubscription)
		r.Post("/block", h.CreateBlock)
		r.Post("/list", h.GetFriends)
		r.Post("/commonfriends", h.GetCommonFriends)
		r.Post("/emailreceive", h.GetEmailReceive)
	}
}
