package models

type CommitMessage struct {
	Id      int `gorm:"primaryKey;autoIncrement:true"`
	Message string
	ModelId int `gorm:"index:commit_member"`
	DiffId  int `gorm:"index:commit_member"`

	Model Model `gorm:"foreignKey:ModelId;constraint:onDelete:CASCADE"`
	Diff  Diff  `gorm:"foreignKey:DiffId;constraint:onDelete:CASCADE"`
}
