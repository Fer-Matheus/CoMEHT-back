package models

type CommitMessage struct {
	Id      int `gorm:"primaryKey;autoIncrement:true"`
	Message string
	Model   Model
	Diff    Diff
}
