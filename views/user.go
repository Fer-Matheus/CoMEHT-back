package views

import (
	"commitinder/models"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserRequest) ToModel() (user models.User) {
	user.Username = u.Username
	user.HashedPassword = u.Password
	user.CurrentDuelId = 1
	return user
}
type LoginResponse struct {
	Token string `json:"token"`
}
