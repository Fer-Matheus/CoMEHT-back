package models

type Diff struct {
	Id             int `gorm:"primaryKey;autoIncrement:true"`
	ContentPath    string
	CommitMessages []CommitMessage
}
