package models

type CommitMessage struct {
	Id      int `gorm:"primaryKey;autoIncrement:true"`
	Message string
	ModelId int
	DiffId  int
}
