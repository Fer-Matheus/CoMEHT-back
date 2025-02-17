package controllers

import (
	"bytes"
	"commitinder/database"
	"commitinder/models"
	"commitinder/utils"
	"commitinder/views"
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func downloadHandler(filePath string, w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(filePath)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filePath)
	os.Remove(filePath)
}

// GetAllResults godoc
// @Summary Get all results
// @Description Route to get all results
// @Tags Results
// @Security Bearer
// @Param login body views.UserRequest true "login"
// @Router /results [put]
func GetAllResults(w http.ResponseWriter, r *http.Request) {

	var results []models.Result

	var userRequest views.UserRequest
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.Db.GetUser("admin", &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !utils.CheckHashPassword(userRequest.Password, user.HashedPassword) {
		http.Error(w, "password doesn't match", http.StatusUnauthorized)
		return
	}

	err = database.Db.GetAllResults(&results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var buffer bytes.Buffer

	write := csv.NewWriter(&buffer)

	write.Write([]string{
		"user_name",
		"duel_id",
		"aspect",
		"model_name_A",
		"model_name_B",
		"model_name_winner",
		"time",
	})

	timeSum := 0

	for _, result := range results {

		var user models.User

		_ = database.Db.GetUserById(result.UserId, &user)

		var duel models.Duel

		_ = database.Db.GetDuel(result.DuelId, &duel)

		var commitMessageA models.CommitMessage
		var commitMessageB models.CommitMessage

		_ = database.Db.GetCommitMessage(duel.CommitMessageAId, &commitMessageA)
		_ = database.Db.GetCommitMessage(duel.CommitMessageBId, &commitMessageB)

		winner := ""

		if result.ChosenOption == "A" {
			winner = commitMessageA.Model.Name
		} else {
			winner = commitMessageB.Model.Name
		}

		timeSum += int(result.ChoiceTime)

		write.Write([]string{
			user.Username,
			strconv.Itoa(result.DuelId),
			result.Aspect,
			commitMessageA.Model.Name,
			commitMessageB.Model.Name,
			winner,
			strconv.Itoa(int(result.ChoiceTime) / 60) +	"m and " + strconv.Itoa(int(result.ChoiceTime) % 60) +"s",	
		})
	}

	days := timeSum / 86400
	timeSum = timeSum % 86400

	hours := timeSum / 3600
	timeSum = timeSum % 3600

	minutes := timeSum / 60
	seconds := timeSum % 60

	write.Write([]string{
		"",
		"",
		"",
		"",
		"",
		"",
		strconv.Itoa(days) + "d, " + strconv.Itoa(hours) + "h, " + strconv.Itoa(minutes) + "m, " + strconv.Itoa(seconds) +"s",	
	})

	write.Flush()

	timestamp := time.Now().UTC().String()

	os.WriteFile(timestamp, buffer.Bytes(), 0600)

	downloadHandler(timestamp, w, r)
}