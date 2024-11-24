package models

type Duel struct {
	Id               int `gorm:"primaryKey;autoIncrement:true"`
	CommitMessageAId int `gorm:"uniqueIndex:user_duel"`
	CommitMessageBId int `gorm:"uniqueIndex:user_duel"`

	Users []User `gorm:"many2many:user_duels"`

	CommitMessageA CommitMessage `gorm:"foreignKey:CommitMessageAId;constrait:OnDelete:CASCADE"`
	CommitMessageB CommitMessage `gorm:"foreignKey:CommitMessageBId;constrait:OnDelete:CASCADE"`
}
