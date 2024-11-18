package models

type User struct {
	Id             int    `gorm:"primaryKey"`
	Username       string `gorm:"unique"`
	HashedPassword string
	CurrentDuelId  int
}
