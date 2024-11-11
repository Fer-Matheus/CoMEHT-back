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
	for _, model := range models {
		err := database.Db.SaveModel(&model)
		if err != nil {
			continue
		}
	}
}
