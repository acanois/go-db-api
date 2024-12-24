package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/acanois/dbapi/api/auth"
	"github.com/lpernett/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	authConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8000/auth/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	app := auth.App{Config: authConfig}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("POST /login", app.Login)
	mux.HandleFunc("POST /auth", app.Auth)
	mux.HandleFunc("POST /callback", app.AuthCallback)
	mux.HandleFunc("GET /addUser", addUser)

	fmt.Printf("Server listening on port: %s", port)
	http.ListenAndServe(port, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	// This is temporary just so I can see something at the home page
	fmt.Fprint(w, "Home")
}
