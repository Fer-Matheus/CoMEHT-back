package seeds

import (
	"comeht/database"
	"comeht/models"
)

func SeedModels() {
	models := []models.Model{
		{
			Name: "Human",
		},
		{
			Name: "FIRA",
		},
		{
			Name: "OMG",
		},
	}
	_ = database.Db.SaveModels(&models)
}
