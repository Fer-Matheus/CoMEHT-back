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
		for _, duel := range duels{
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
			Users: users,
		},
		{
			CommitMessageAId: 3,
			CommitMessageBId: 4,
			Users: users,
		},
	}
	_ = database.Db.SaveDuels(&seeds)
}
