package models

type UserDuel struct {
	UserId int  `gorm:"primaryKey"`
	DuelId int  `gorm:"primaryKey"`
	User   User `gorm:"foreignKey:UserId;constrait:OnDelete:CASCADE"`
	Duel   Duel `gorm:"foreignKey:DuelId;constrait:OnDelete:CASCADE"`

	Results []Result `gorm:"foreignKey:UserId,DuelId;references:UserId,DuelId"`
}
func (UserDuel) TableName() string {
    return "user_duels"
}