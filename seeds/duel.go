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
		{CommitMessageAId: 2, CommitMessageBId: 1, Users: users},
		{CommitMessageAId: 5, CommitMessageBId: 6, Users: users},
		{CommitMessageAId: 7, CommitMessageBId: 9, Users: users},
		{CommitMessageAId: 12, CommitMessageBId: 11, Users: users},
		{CommitMessageAId: 13, CommitMessageBId: 15, Users: users},
		{CommitMessageAId: 17, CommitMessageBId: 18, Users: users},
		{CommitMessageAId: 19, CommitMessageBId: 21, Users: users},
		{CommitMessageAId: 24, CommitMessageBId: 22, Users: users},
		{CommitMessageAId: 27, CommitMessageBId: 25, Users: users},
		{CommitMessageAId: 30, CommitMessageBId: 28, Users: users},
		{CommitMessageAId: 33, CommitMessageBId: 32, Users: users},
		{CommitMessageAId: 34, CommitMessageBId: 36, Users: users},
		{CommitMessageAId: 37, CommitMessageBId: 38, Users: users},
		{CommitMessageAId: 41, CommitMessageBId: 42, Users: users},
		{CommitMessageAId: 43, CommitMessageBId: 45, Users: users},
		{CommitMessageAId: 47, CommitMessageBId: 48, Users: users},
		{CommitMessageAId: 51, CommitMessageBId: 50, Users: users},
		{CommitMessageAId: 53, CommitMessageBId: 52, Users: users},
		{CommitMessageAId: 56, CommitMessageBId: 55, Users: users},
		{CommitMessageAId: 59, CommitMessageBId: 58, Users: users},
		{CommitMessageAId: 61, CommitMessageBId: 63, Users: users},
		{CommitMessageAId: 66, CommitMessageBId: 64, Users: users},
		{CommitMessageAId: 67, CommitMessageBId: 69, Users: users},
		{CommitMessageAId: 71, CommitMessageBId: 72, Users: users},
		{CommitMessageAId: 75, CommitMessageBId: 73, Users: users},
		{CommitMessageAId: 76, CommitMessageBId: 77, Users: users},
		{CommitMessageAId: 80, CommitMessageBId: 79, Users: users},
		{CommitMessageAId: 83, CommitMessageBId: 82, Users: users},
		{CommitMessageAId: 86, CommitMessageBId: 85, Users: users},
		{CommitMessageAId: 88, CommitMessageBId: 90, Users: users},
	}
	_ = database.Db.SaveDuels(&seeds)
}
