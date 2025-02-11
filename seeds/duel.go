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
		{Users: users, CommitMessageAId: 3, CommitMessageBId: 1},
		{Users: users, CommitMessageAId: 5, CommitMessageBId: 7},
		{Users: users, CommitMessageAId: 11, CommitMessageBId: 10},
		{Users: users, CommitMessageAId: 16, CommitMessageBId: 13},
		{Users: users, CommitMessageAId: 19, CommitMessageBId: 18},
		{Users: users, CommitMessageAId: 24, CommitMessageBId: 22},
		{Users: users, CommitMessageAId: 26, CommitMessageBId: 28},
		{Users: users, CommitMessageAId: 32, CommitMessageBId: 29}}
	_ = database.Db.SaveDuels(&seeds)
}
