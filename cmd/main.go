package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/acanois/dbapi/internal/auth"
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
	mux.HandleFunc("GET /fetch", fetchData)

	fmt.Printf("Server listening on port: %s", port)
	http.ListenAndServe(port, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}

func fetchData(w http.ResponseWriter, r *http.Request) {
	DB_USER := os.Getenv("DB_USER")
	DB_PW := os.Getenv("DB_PW")
	DB_NAME := os.Getenv("DB_NAME")
	conn := fmt.Sprintf(
		"postgres://%s:%s@localhost:5432/%s?sslmode=disable",
		DB_USER,
		DB_PW,
		DB_NAME,
	)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// var (
	// 	id   int
	// 	name string
	// )
	// rows, err := db.Query("select id, name from users where id = ?", 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(id, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
