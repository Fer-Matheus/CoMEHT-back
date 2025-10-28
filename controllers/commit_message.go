package controllers

import (
	"comeht/database"
	"comeht/models"
	"comeht/views"
	"encoding/json"
	"net/http"
	"strconv"
)

// SaveCommitMessage godoc
//	@Summary		Save a new commit message
//	@Description	A route to save a new commit message
//	@Tags			Commit Message
//	@Accept			json
//	@Produce		json
//	@Param			commitMessageRequest	body		views.CommitMessageRequest	true	"commitMessageRequest"
//	@Success		201						{object}	views.CommitMessageResponse
//	@Router			/commit_messages [post]
func SaveCommitMessage(w http.ResponseWriter, r *http.Request) {

	var commitMessageRequest views.CommitMessageRequest
	var commitMessageResponse views.CommitMessageResponse

	err := json.NewDecoder(r.Body).Decode(&commitMessageRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var model models.Model
	err = database.Db.GetModel(commitMessageRequest.ModelId, &model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var diff models.Diff
	err = database.Db.GetDiff(commitMessageRequest.DiffId, &diff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	commitMessage := commitMessageRequest.ToModel()
	err = database.Db.SaveCommitMessage(&commitMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	model.CommitMessages = append(model.CommitMessages, commitMessage)
	diff.CommitMessages = append(diff.CommitMessages, commitMessage)

	err = database.Db.UpdateModel(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	err = database.Db.UpdateDiff(&diff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	commitMessageResponse.FromModelToView(&commitMessage)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commitMessageResponse)
}

// GetCommitMessage godoc
//	@Summary		Get a commit message
//	@Description	A route to get a commit message by id
//	@Tags			Commit Message
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		201	{object}	views.CommitMessageResponse
//	@Router			/commit_messages/{id} [get]
func GetCommitMessage(w http.ResponseWriter, r *http.Request) {
	var commitMessage models.CommitMessage
	var commitMessageResponse views.CommitMessageResponse

	id := r.PathValue("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.Db.GetCommitMessage(intId, &commitMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	commitMessageResponse.FromModelToView(&commitMessage)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commitMessageResponse)
}

// GetAllCommitMessages godoc
//	@Summary		Get all commit messages
//	@Description	A route to get all commit messages
//	@Tags			Commit Message
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	views.CommitMessagesResponse
//	@Router			/commit_messages [get]
func GetAllCommitMessages(w http.ResponseWriter, r *http.Request) {
	var modelsCommitMessages []models.CommitMessage
	var commitMessageResponse views.CommitMessageResponse
	var commitMessages views.CommitMessagesResponse

	err := database.Db.GetAllCommitMessages(&modelsCommitMessages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	for _, modelDiff := range modelsCommitMessages {
		commitMessageResponse.FromModelToView(&modelDiff)
		commitMessages.CommitMessages = append(commitMessages.CommitMessages, commitMessageResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(commitMessages)
}
