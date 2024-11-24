package controllers

import (
	"bytes"
	"commitinder/database"
	"commitinder/models"
	"encoding/csv"
	"net/http"
	"strconv"
)

// GetAllResults godoc
// @Summary Get all results
// @Description Route to get all results
// @Tags Results
// @Security Bearer
// @Router /results [get]
func GetAllResults(w http.ResponseWriter, r *http.Request) {

	var results []models.Result

	_, err := authorize(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
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
		"user_id",
		"duel_id",
		"aspect",
		"choice_time",
		"chosen_option",
	})

	for _, result := range results {
		write.Write([]string{
			strconv.Itoa(result.UserId),
			strconv.Itoa(result.DuelId),
			result.Aspect,
			strconv.FormatFloat(result.ChoiceTime, 'f', 4, 64),
			result.ChosenOption,
		})
	}

	write.Flush()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Length", strconv.Itoa(buffer.Len()))
	w.Write(buffer.Bytes())
}
