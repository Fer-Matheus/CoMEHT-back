package models

type Result struct {
	Option       string
	TimeToChoice string
	DuelId       int
}

type Duel struct {
	Id             int `gorm:"primaryKey;autoIncrement:true"`
	CommitMessageA CommitMessage
	CommitMessageB CommitMessage
	Results        []Result
}
