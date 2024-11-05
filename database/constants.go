package database

import (
	"commitinder/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db = Init()

type Database struct {
	db *gorm.DB
}

func Init() (db Database) {

	dsn := "host=postgres user=admin password=admin port=5432"
	db.db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.db.AutoMigrate(&models.Diff{}, &models.Model{}, &models.CommitMessage{}, &models.Duel{}, &models.Result{})

	return db
}