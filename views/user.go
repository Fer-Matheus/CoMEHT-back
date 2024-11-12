package views

import (
	"commitinder/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRegistration struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserRegistration) ToModel() (user models.User) {
	user.Username = u.Username
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	user.HashedPassword = string(hashedPassword)
	user.CurrentDuelId = 1
	return user
}
