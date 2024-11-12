package controllers

import (
	"commitinder/database"
	"commitinder/models"
	"commitinder/seeds"
	"commitinder/utils"
	"commitinder/views"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var errorAuth = errors.New("unauthorized")

func authorize(r *http.Request) error {
	token, err := r.Cookie("commitinder_session_token")
	if err != nil || token.Value == "" {
		return errorAuth
	}

	var user models.User

	getUserByCookie(&user, r)

	if token.Value != user.SessionToken {
		return errorAuth
	}
	// csrf := r.Header.Get("X-CSRF-Token")
	// if csrf != user.CsrfToken || csrf == "" {
	// 	return errorAuth
	// }

	return nil
}

func getUserByCookie(user *models.User, r *http.Request) error {
	username, _ := r.Cookie("user")

	username.Value = strings.Replace(username.Value, "'", "", 2)

	err := database.Db.GetUser(username.Value, user)
	if err != nil {
		return err
	}
	return nil
}

// Register godoc
// @Summary Register a new user
// @Description A route to register a new user
// @Tags User
// @Accept json
// @Param userRequest body views.UserRequest true "userRequest"
// @Router /register [post]
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

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "User registered successfully!")
}

// Login godoc
// @Summary Login for a user
// @Description A route for user login
// @Tags User
// @Accept json
// @Param userRequest body views.UserRequest true "userRequest"
// @Router /login [post]
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
		http.Error(w, "User not authorized", http.StatusUnauthorized)
		return
	}

	sessionToken := utils.GenerateToken(32)
	// csrfToken := utils.GenerateToken(32)

	user.SessionToken = sessionToken
	// user.CsrfToken = csrfToken

	err = database.Db.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "commitinder_session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "csrf_token",
	// 	Value:    csrfToken,
	// 	Expires:  time.Now().Add(24 * time.Hour),
	// 	HttpOnly: false,
	// })
	http.SetCookie(w, &http.Cookie{
		Name:     "user",
		Value:    user.Username,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	fmt.Fprintln(w, "Login successful!")
}

// Logout godoc
// @Summary Logout a user
// @Description A route to logout a user by cookie
// @Tags User
// @Router /logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {
	err := authorize(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var user models.User

	err = getUserByCookie(&user, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user.SessionToken = ""
	// user.CsrfToken = ""

	err = database.Db.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "commitinder_session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "csrf_token",
	// 	Value:    "",
	// 	Expires:  time.Now().Add(-time.Hour),
	// 	HttpOnly: true,
	// })
	fmt.Fprintln(w, "Logged out successfully!")
}
