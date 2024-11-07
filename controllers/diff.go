package controllers

import (
	"commitinder/database"
	"commitinder/models"
	"commitinder/views"
	"encoding/json"
	"net/http"
	"strconv"
)

// SaveDiff godoc
// @Summary Save a new diff
// @Description A route to save a new diff
// @Tags Diff
// @Accept json
// @Produce json
// @Param diffRequest body views.DiffRequest true "diffRequest"
// @Success 201 {object} views.DiffResponse
// @Router /diffs [post]
func SaveDiff(w http.ResponseWriter, r *http.Request) {

	var diffRequest views.DiffRequest
	var diffResponse views.DiffResponse

	err := json.NewDecoder(r.Body).Decode(&diffRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	diff := diffRequest.ToModel()
	err = database.Db.SaveDiff(&diff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	diffResponse.FormModelToView(&diff)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diffResponse)
}

// GetDiff godoc
// @Summary Get a diff
// @Description A route to get a diff by id
// @Tags Diff
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 201 {object} views.DiffResponse
// @Router /diffs [get]
func GetDiff(w http.ResponseWriter, r *http.Request) {
	var diff models.Diff
	var diffResponse views.DiffResponse

	id := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.Db.GetDiff(intId, &diff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	diffResponse.FormModelToView(&diff)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diffResponse)
}


