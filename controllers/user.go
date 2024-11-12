package controllers

import (
	"commitinder/database"
	"commitinder/seeds"
	"commitinder/views"
	"encoding/json"
	"fmt"
	"net/http"
)

// Register godoc
// @Summary Register a new user
// @Description A route to register a new user
// @Tags User
// @Accept json
// @Param userRegistration body views.UserRegistration true "userRegistration"
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var userRegistration views.UserRegistration

	err := json.NewDecoder(r.Body).Decode(&userRegistration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := userRegistration.ToModel()

	err = database.Db.SaveUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	seeds.SeedDuels(user.Id)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "User registered successfully!")
}
