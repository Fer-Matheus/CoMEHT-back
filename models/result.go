package models

type Result struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	UserId   int
	DuelId   int
	UserDuel UserDuel `gorm:"foreignKey:UserId,DuelId;references:UserId,DuelId;constraint:OnDelete:CASCADE"`

	Aspect       string
	ChoiceTime   float64
	ChosenOption string
}
