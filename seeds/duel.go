package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedDuels(userId int) {

	var user models.User
	_ = database.Db.GetUserById(userId, &user)

	var duels []models.Duel

	_ = database.Db.GetAllDuels(&duels)

	if len(duels) > 0 {
		for _, duel := range duels {
			duel.Users = append(duel.Users, user)
			database.Db.UpdateDuel(&duel)
		}
		return
	}

	users := []models.User{user}

	seeds := []models.Duel{
		{
			CommitMessageAId: 1,
			CommitMessageBId: 2,
			Users:            users,
		},
		{
			CommitMessageAId: 4,
			CommitMessageBId: 5,
			Users:            users,
		},
		{
			CommitMessageAId: 7,
			CommitMessageBId: 8,
			Users:            users,
		},
		{
			CommitMessageAId: 10,
			CommitMessageBId: 11,
			Users:            users,
		},
		{
			CommitMessageAId: 13,
			CommitMessageBId: 14,
			Users:            users,
		},
		{
			CommitMessageAId: 16,
			CommitMessageBId: 17,
			Users:            users,
		},
		{
			CommitMessageAId: 19,
			CommitMessageBId: 20,
			Users:            users,
		},
		{
			CommitMessageAId: 22,
			CommitMessageBId: 23,
			Users:            users,
		},
		{
			CommitMessageAId: 25,
			CommitMessageBId: 26,
			Users:            users,
		},
		{
			CommitMessageAId: 28,
			CommitMessageBId: 29,
			Users:            users,
		},
		{
			CommitMessageAId: 1,
			CommitMessageBId: 3,
			Users:            users,
		},
		{
			CommitMessageAId: 4,
			CommitMessageBId: 6,
			Users:            users,
		},
		{
			CommitMessageAId: 7,
			CommitMessageBId: 9,
			Users:            users,
		},
		{
			CommitMessageAId: 10,
			CommitMessageBId: 12,
			Users:            users,
		},
		{
			CommitMessageAId: 13,
			CommitMessageBId: 15,
			Users:            users,
		},
		{
			CommitMessageAId: 16,
			CommitMessageBId: 18,
			Users:            users,
		},
		{
			CommitMessageAId: 19,
			CommitMessageBId: 21,
			Users:            users,
		},
		{
			CommitMessageAId: 22,
			CommitMessageBId: 24,
			Users:            users,
		},
		{
			CommitMessageAId: 25,
			CommitMessageBId: 27,
			Users:            users,
		},
		{
			CommitMessageAId: 28,
			CommitMessageBId: 30,
			Users:            users,
		},
		{
			CommitMessageAId: 2,
			CommitMessageBId: 3,
			Users:            users,
		},

		{
			CommitMessageAId: 5,
			CommitMessageBId: 6,
			Users:            users,
		},

		{
			CommitMessageAId: 8,
			CommitMessageBId: 9,
			Users:            users,
		},

		{
			CommitMessageAId: 11,
			CommitMessageBId: 12,
			Users:            users,
		},

		{
			CommitMessageAId: 14,
			CommitMessageBId: 15,
			Users:            users,
		},

		{
			CommitMessageAId: 17,
			CommitMessageBId: 18,
			Users:            users,
		},

		{
			CommitMessageAId: 20,
			CommitMessageBId: 21,
			Users:            users,
		},

		{
			CommitMessageAId: 23,
			CommitMessageBId: 24,
			Users:            users,
		},

		{
			CommitMessageAId: 26,
			CommitMessageBId: 27,
			Users:            users,
		},

		{
			CommitMessageAId: 29,
			CommitMessageBId: 30,
			Users:            users,
		},
	}
	_ = database.Db.SaveDuels(&seeds)
}
