package models

type Result struct {
	Id           int    `gorm:"primaryKey;autoIncrement;uniqueIndex:aspect_duel"`
	Aspect       string 
	ChosenOption string
	ChoiseTime   string

	DuelId int  `gorm:"uniqueIndex:aspect_duel"`
	Duel   Duel `gorm:"foreignKey:DuelId;constrait:OnDelete:CASCADE"`
}

type Duel struct {
	Id               int  `gorm:"primaryKey;autoIncrement:true"`
	CommitMessageAId int  `gorm:"uniqueIndex:user_duel"`
	CommitMessageBId int  `gorm:"uniqueIndex:user_duel"`
	UserId           int  `gorm:"uniqueIndex:user_duel"`
	User             User `gorm:"foreignKey:UserId;constrait:OnDelete:CASCADE"`
	Results          []Result

	CommitMessageA CommitMessage `gorm:"foreignKey:CommitMessageAId;constrait:OnDelete:CASCADE"`
	CommitMessageB CommitMessage `gorm:"foreignKey:CommitMessageBId;constrait:OnDelete:CASCADE"`
}
