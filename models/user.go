package models

type User struct {
	Id             int    `gorm:"primaryKey"`
	Username       string `gorm:"unique"`
	HashedPassword string
	SessionToken   string
	CurrentDuelId  int
}
