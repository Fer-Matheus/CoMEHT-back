package seeds

import (
	"commitinder/database"
	"commitinder/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	password, _ := bcrypt.GenerateFromPassword([]byte("uga"), 10)
	user := models.User{
		Username:       "Matheus",
		HashedPassword: string(password),
		CurrentDuelId:  1,
	}
	_ = database.Db.SaveUser(&user)
	SeedDuels(user.Id)
}
