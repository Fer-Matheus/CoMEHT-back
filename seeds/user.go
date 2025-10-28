package seeds

import (
	"comeht/database"
	"comeht/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), 10)
	user := models.User{
		Username:       "admin",
		HashedPassword: string(password),
		CurrentDuelId:  1,
	}
	_ = database.Db.SaveUser(&user)
	SeedDuels(user.Id)
}
