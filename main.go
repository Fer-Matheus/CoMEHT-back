package main

import (
	"comeht/controllers"
	_ "comeht/docs"
	"comeht/seeds"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func alreadySeeded() bool {
	_, err := os.ReadFile("alreadySeeded")
	return err == nil
}
func setAlreadySeeded() {
	os.WriteFile("alreadySeeded", nil, 0600)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@title			comeht Server
//	@version		1.0
//	@description	This is the comeht server api.
//
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @Security Bearer
func main() {

	if !alreadySeeded() {
		seeds.SeedDiffs()
		seeds.SeedModels()
		seeds.SeedCommitMessages()
		seeds.SeedUsers()
		setAlreadySeeded()
	}

	mux := http.NewServeMux()

	mux.Handle("/swagger/", http.StripPrefix("/swagger/", httpSwagger.WrapHandler))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("success")
	})

	// Diff
	mux.HandleFunc("/diffs/{id}", controllers.GetDiff)
	mux.HandleFunc("/diffs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllDiffs(w, r)
		case http.MethodPost:
			controllers.SaveDiff(w, r)
		}
	})

	//Model
	mux.HandleFunc("/models/{id}", controllers.GetModel)
	mux.HandleFunc("/models", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllModels(w, r)
		case http.MethodPost:
			controllers.SaveModel(w, r)
		}
	})

	// Commit message
	mux.HandleFunc("/commit_messages/{id}", controllers.GetCommitMessage)
	mux.HandleFunc("/commit_messages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllCommitMessages(w, r)
		case http.MethodPost:
			controllers.SaveCommitMessage(w, r)
		}
	})

	// User
	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/logout", controllers.Logout)

	mux.HandleFunc("/duels", controllers.GetDuel)
	mux.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			controllers.GetAllResults(w, r)
		case http.MethodPost:
			controllers.SaveResults(w, r)
		}
	})

	handler := cors.AllowAll().Handler(mux)

	port := os.Getenv("PORT")
	fmt.Println("Server listen at the port: " + port)
	http.ListenAndServe(":"+port, handler)
}
