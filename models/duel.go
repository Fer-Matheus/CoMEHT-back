package models

type Result struct {
	Option       string
	TimeToChoice string
	DuelId       int
}

type Duel struct {
	Id               int `gorm:"primaryKey;autoIncrement:true"`
	CommitMessageAId int
	CommitMessageBId int

	Results []Result
}
