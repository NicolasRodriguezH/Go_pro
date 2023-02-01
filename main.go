package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	hanlders "pruebago.com/go/rest-ws/handlers"
	"pruebago.com/go/rest-ws/middleware"
	"pruebago.com/go/rest-ws/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loagind .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/", hanlders.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", hanlders.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", hanlders.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", hanlders.MeHandler(s)).Methods(http.MethodGet)
}
