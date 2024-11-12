package controllers

import (
	"commitinder/database"
	"commitinder/models"
	"commitinder/views"
	"encoding/json"
	"net/http"
	"strconv"
)

// SaveModel godoc
// @Summary Save a new model
// @Description A route to save a new model
// @Tags Model
// @Accept json
// @Produce json
// @Param modelRequest body views.ModelRequest true "modelRequest"
// @Success 201 {object} views.ModelResponse
// @Router /models [post]
func SaveModel(w http.ResponseWriter, r *http.Request) {

	var modelRequest views.ModelRequest
	var modelResponse views.ModelResponse

	err := json.NewDecoder(r.Body).Decode(&modelRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := modelRequest.ToModel()
	err = database.Db.SaveModel(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	modelResponse.FromModelToView(&model)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modelResponse)
}

// GetModel godoc
// @Summary Get a model
// @Description A route to get a model by id
// @Tags Model
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} views.ModelResponse
// @Router /models/{id} [get]
func GetModel(w http.ResponseWriter, r *http.Request) {
	var model models.Model
	var modelResponse views.ModelResponse

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.Db.GetModel(intId, &model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	modelResponse.FromModelToView(&model)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modelResponse)
}

// GetAllModels godoc
// @Summary Get all models
// @Description A route to get all models
// @Tags Model
// @Accept json
// @Produce json
// @Success 200 {object} views.ModelsResponse
// @Router /models [get]
func GetAllModels(w http.ResponseWriter, r *http.Request) {
	var modelModels []models.Model
	var modelReponse views.ModelResponse
	var models views.ModelsResponse

	err := database.Db.GetAllModels(&modelModels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	for _, modelDiff := range modelModels {
		modelReponse.FromModelToView(&modelDiff)
		models.Models = append(models.Models, modelReponse)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models)
}
