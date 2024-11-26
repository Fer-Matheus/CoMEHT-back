package controllers

import (
	"commitinder/database"
	"commitinder/models"
	"commitinder/seeds"
	"commitinder/utils"
	"commitinder/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func authorize(r *http.Request) (int, error) {
	token := r.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		return 0, err
	}
	intUserId := int(userId)
	return intUserId, nil
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	A route to register a new user
//	@Tags			User
//	@Accept			json
//	@Param			userRequest	body	views.UserRequest	true	"userRequest"
//	@Success 		200 		{object} 	views.LoginResponse
//	@Router			/register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var userRequest views.UserRequest

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := userRequest.ToModel()

	err = utils.HashPassword(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	err = database.Db.SaveUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	seeds.SeedDuels(user.Id)

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Authorization", token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(views.LoginResponse{Token: token})
}

// Login godoc
//
//	@Summary		Login for a user
//	@Description	A route for user login
//	@Tags			User
//	@Accept			json
//	@Param			userRequest	body		views.UserRequest	true	"userRequest"
//	@Success 		200 		{object} 	views.LoginResponse
//	@Router			/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var userRequest views.UserRequest
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.Db.GetUser(userRequest.Username, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !utils.CheckHashPassword(userRequest.Password, user.HashedPassword) {
		http.Error(w, "password doesn't match", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(views.LoginResponse{Token: token})
}

// Logout godoc
//
//	@Summary		Logout a user
//	@Description	A route to logout a user by cookie
//	@Tags			User
//	@Security Bearer
//	@Router			/logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	_, err := authorize(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintln(w, "Logged out successfully!")
}
