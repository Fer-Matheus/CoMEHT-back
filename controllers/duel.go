package controllers

import (
	"commitinder/database"
	"commitinder/models"
	"commitinder/views"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// GetDuel godoc
//
//	@Summary		Get a new duel
//	@Description	A route to get a duel for a user
//	@Tags			Duel
//	@Produce		json
//	@Security Bearer
//	@Success		200				{object}	views.DuelResponse
//	@Router			/duels [get]
func GetDuel(w http.ResponseWriter, r *http.Request) {
	userId, err := authorize(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var user models.User
	err = database.Db.GetUserById(userId, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var duel models.Duel
	err = database.Db.GetDuel(user.CurrentDuelId, &duel)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var commitMessageA models.CommitMessage

	err = database.Db.GetCommitMessage(duel.CommitMessageAId, &commitMessageA)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var commitMessageB models.CommitMessage
	err = database.Db.GetCommitMessage(duel.CommitMessageBId, &commitMessageB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var diff models.Diff

	err = database.Db.GetDiff(commitMessageA.DiffId, &diff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	content, _ := os.ReadFile("diffs/" + diff.ContentPath)

	duelResponse := &views.DuelResponse{
		DuelId:         duel.Id,
		DiffContent:    string(content),
		CommitMessageA: commitMessageA.Message,
		CommitMessageB: commitMessageB.Message,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content_type", "application/json")
	json.NewEncoder(w).Encode(duelResponse)
}

// SaveResults godoc
//
//	@Summary		Save the duel results
//	@Description	A route to save a duel results
//	@Tags			Duel
//	@Accept			json
//	@Produce		json
//	@Security Bearer
//	@Param			results	body	views.ResultsRequest	true	"results"
//	@Router			/results [post]
func SaveResults(w http.ResponseWriter, r *http.Request) {
	var results views.ResultsRequest
	var duel models.Duel

	userId, err := authorize(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.Db.GetDuel(results.DuelId, &duel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var user models.User

	err = database.Db.GetUserById(userId, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if user.CurrentDuelId != duel.Id {
		http.Error(w, "Results already collected", http.StatusConflict)
		return
	}

	var userDuel models.UserDuel

	userDuel.User = user
	userDuel.UserId = user.Id
	userDuel.Duel = duel
	userDuel.DuelId = duel.Id

	// err = database.Db.SaveUserDuel(&userDuel)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusConflict)
	// 	return
	// }

	for _, result := range results.Options{

		var modelResult models.Result

		modelResult.Aspect = result.Aspect
		modelResult.ChoiceTime = result.ChoiceTime
		modelResult.ChosenOption = result.ChosenOption

		modelResult.DuelId = userDuel.DuelId
		modelResult.UserId = userDuel.UserId
		modelResult.UserDuel = userDuel

		userDuel.Results = append(userDuel.Results, modelResult)
	}

	err = database.Db.SaveResults(&userDuel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	user.CurrentDuelId += 1

	err = database.Db.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content_type", "application/json")
	fmt.Fprintln(w, "Results save successfully")
}
