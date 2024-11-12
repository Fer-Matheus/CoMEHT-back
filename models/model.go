package models

type Model struct {
	Id             int    `gorm:"primaryKey;autoIncrement:true"`
	Name           string `gorm:"unique;index"`
	CommitMessages []CommitMessage
}
