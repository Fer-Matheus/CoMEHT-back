package seeds

import (
	"comeht/database"
	"comeht/models"
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
		{Users: users, CommitMessageAId: 1, CommitMessageBId: 3},
		{Users: users, CommitMessageAId: 6, CommitMessageBId: 4},
		{Users: users, CommitMessageAId: 7, CommitMessageBId: 8},
		{Users: users, CommitMessageAId: 10, CommitMessageBId: 11},
		{Users: users, CommitMessageAId: 15, CommitMessageBId: 13},
		{Users: users, CommitMessageAId: 16, CommitMessageBId: 17},
		{Users: users, CommitMessageAId: 19, CommitMessageBId: 20},
		{Users: users, CommitMessageAId: 22, CommitMessageBId: 23},
		{Users: users, CommitMessageAId: 26, CommitMessageBId: 27},
		{Users: users, CommitMessageAId: 30, CommitMessageBId: 28},
		{Users: users, CommitMessageAId: 32, CommitMessageBId: 31},
		{Users: users, CommitMessageAId: 35, CommitMessageBId: 36},
		{Users: users, CommitMessageAId: 39, CommitMessageBId: 38},
		{Users: users, CommitMessageAId: 41, CommitMessageBId: 42},
		{Users: users, CommitMessageAId: 44, CommitMessageBId: 45},
		{Users: users, CommitMessageAId: 47, CommitMessageBId: 48},
		{Users: users, CommitMessageAId: 50, CommitMessageBId: 51},
		{Users: users, CommitMessageAId: 52, CommitMessageBId: 53},
		{Users: users, CommitMessageAId: 55, CommitMessageBId: 57},
		{Users: users, CommitMessageAId: 59, CommitMessageBId: 58},
		{Users: users, CommitMessageAId: 61, CommitMessageBId: 63},
		{Users: users, CommitMessageAId: 66, CommitMessageBId: 65},
		{Users: users, CommitMessageAId: 69, CommitMessageBId: 68},
		{Users: users, CommitMessageAId: 72, CommitMessageBId: 70},
		{Users: users, CommitMessageAId: 74, CommitMessageBId: 73},
		{Users: users, CommitMessageAId: 77, CommitMessageBId: 78},
		{Users: users, CommitMessageAId: 79, CommitMessageBId: 81},
		{Users: users, CommitMessageAId: 83, CommitMessageBId: 82},
		{Users: users, CommitMessageAId: 87, CommitMessageBId: 86},
		{Users: users, CommitMessageAId: 90, CommitMessageBId: 89},
		{Users: users, CommitMessageAId: 91, CommitMessageBId: 93},
		{Users: users, CommitMessageAId: 95, CommitMessageBId: 94},
		{Users: users, CommitMessageAId: 99, CommitMessageBId: 97},
		{Users: users, CommitMessageAId: 102, CommitMessageBId: 100},
		{Users: users, CommitMessageAId: 104, CommitMessageBId: 105},
		{Users: users, CommitMessageAId: 108, CommitMessageBId: 107},
		{Users: users, CommitMessageAId: 110, CommitMessageBId: 111},
		{Users: users, CommitMessageAId: 113, CommitMessageBId: 112},
		{Users: users, CommitMessageAId: 115, CommitMessageBId: 116},
		{Users: users, CommitMessageAId: 120, CommitMessageBId: 118},
		{Users: users, CommitMessageAId: 121, CommitMessageBId: 123},
		{Users: users, CommitMessageAId: 124, CommitMessageBId: 126},
		{Users: users, CommitMessageAId: 129, CommitMessageBId: 127},
		{Users: users, CommitMessageAId: 132, CommitMessageBId: 130},
		{Users: users, CommitMessageAId: 135, CommitMessageBId: 134},
		{Users: users, CommitMessageAId: 138, CommitMessageBId: 136},
		{Users: users, CommitMessageAId: 141, CommitMessageBId: 140},
		{Users: users, CommitMessageAId: 142, CommitMessageBId: 144},
		{Users: users, CommitMessageAId: 145, CommitMessageBId: 147},
		{Users: users, CommitMessageAId: 148, CommitMessageBId: 150},
		{Users: users, CommitMessageAId: 153, CommitMessageBId: 151},
		{Users: users, CommitMessageAId: 155, CommitMessageBId: 154},
		{Users: users, CommitMessageAId: 159, CommitMessageBId: 158},
		{Users: users, CommitMessageAId: 162, CommitMessageBId: 160},
		{Users: users, CommitMessageAId: 163, CommitMessageBId: 164},
		{Users: users, CommitMessageAId: 166, CommitMessageBId: 167},
		{Users: users, CommitMessageAId: 169, CommitMessageBId: 170},
		{Users: users, CommitMessageAId: 173, CommitMessageBId: 172},
		{Users: users, CommitMessageAId: 175, CommitMessageBId: 177},
		{Users: users, CommitMessageAId: 179, CommitMessageBId: 180},
	}
	_ = database.Db.SaveDuels(&seeds)
}
