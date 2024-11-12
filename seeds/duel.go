package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedDuels(userId int) {
	seeds := []models.Duel{
		{
			CommitMessageAId: 1,
			CommitMessageBId: 2,
			UserId:           userId,
		},
		{
			CommitMessageAId: 3,
			CommitMessageBId: 4,
			UserId:           userId,
		},
	}
	_ = database.Db.SaveDuels(&seeds)
}
