package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedModels() {
	models := []models.Model{
		{
			Name: "ChatGPT",
		},
		{
			Name: "NNGen",
		},
		{
			Name: "Human",
		},
	}
	_ = database.Db.SaveModels(&models)
}
