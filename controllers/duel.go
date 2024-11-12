package controllers

import (
	"commitinder/database"
	"commitinder/models"
	"commitinder/views"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetDuel godoc
// @Summary Get a new duel
// @Description A route to get a duel for a user
// @Tags Duel
// @Produce json
// @Success 200 {object} views.DuelResponse
// @Router /duels [get]
func GetDuel(w http.ResponseWriter, r *http.Request) {
	err := authorize(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var user models.User

	err = getUserByCookie(&user, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var duel models.Duel

	err = database.Db.GetDuel(user.CurrentDuelId, &duel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

	duelResponse := &views.DuelResponse{
		DuelId: duel.Id,
		DiffContent: diff.Content,
		CommitMessageA: commitMessageA.Message,
		CommitMessageB: commitMessageB.Message,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content_type", "application/json")
	json.NewEncoder(w).Encode(duelResponse)
}

// SaveResults godoc
// @Summary Save the duel results
// @Description A route to save a duel results
// @Tags Duel
// @Accept json
// @Produce json
// @Param results body views.ResultsRequest true "results"
// @Router /results [post]
func SaveResults(w http.ResponseWriter, r *http.Request) {
	var results views.ResultsRequest
	var duel models.Duel
	var modelResult models.Result

	err := authorize(r)
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

	for _, result := range results.Options {
		modelResult.ChosenOption = result.ChosenOption
		modelResult.ChoiseTime = result.ChoiseTime
		modelResult.Duel = duel
		modelResult.DuelId = duel.Id

		err = database.Db.SaveResult(&modelResult)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		duel.Results = append(duel.Results, modelResult)
	}

	err = database.Db.UpdateDuel(&duel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	duel.User.CurrentDuelId = +1

	err = database.Db.UpdateUser(&duel.User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content_type", "application/json")
	fmt.Fprintln(w, "Results save successfully")
}
