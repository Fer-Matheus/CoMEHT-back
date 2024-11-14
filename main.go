package main

import (
	"commitinder/controllers"
	_ "commitinder/docs"
	"commitinder/seeds"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@title			Commitinder Server
//	@version		1.0
//	@description	This is the commitinder server api.
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @Security Bearer
func main() {

	seeds.SeedDiffs()
	seeds.SeedModels()
	seeds.SeedCommitMessages()

	http.Handle("/swagger/", http.StripPrefix("/swagger/", httpSwagger.WrapHandler))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("success")
	})

	// Diff
	http.HandleFunc("/diffs/{id}", controllers.GetDiff)
	http.HandleFunc("/diffs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllDiffs(w, r)
		case http.MethodPost:
			controllers.SaveDiff(w, r)
		}
	})

	//Model
	http.HandleFunc("/models/{id}", controllers.GetModel)
	http.HandleFunc("/models", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllModels(w, r)
		case http.MethodPost:
			controllers.SaveModel(w, r)
		}
	})

	// Commit message
	http.HandleFunc("/commit_messages/{id}", controllers.GetCommitMessage)
	http.HandleFunc("/commit_messages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetAllCommitMessages(w, r)
		case http.MethodPost:
			controllers.SaveCommitMessage(w, r)
		}
	})

	// User
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)

	http.HandleFunc("/duels", controllers.GetDuel)
	http.HandleFunc("/results", controllers.SaveResults)

	port := os.Getenv("PORT")
	fmt.Println("Server listen at the port: " + port)
	http.ListenAndServe(":"+port, nil)
}
