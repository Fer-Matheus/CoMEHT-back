package main

import (
	"commitinder/controllers"
	_ "commitinder/docs"
	"encoding/json"
	"fmt"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

var port = ":8080"

// @title	Commitinder Server
// @version 1.0
// @description	This is the commitinder server api.
func main() {

	http.Handle("/swagger/", http.StripPrefix("/swagger/", httpSwagger.WrapHandler))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("success")
	})

	http.HandleFunc("/diffs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.SaveDiff(w, r)
		case http.MethodPost:
			controllers.SaveDiff(w, r)
		}
	})

	fmt.Println("Server listen at the port " + port)
	http.ListenAndServe(port, nil)
}
