package database

import (
	"commitinder/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db = Init()

type Database struct {
	db *gorm.DB
}

func Init() Database {

	dsn := "host=postgres user=admin password=admin port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("error connecting with the database: " + err.Error())
	}
	db.AutoMigrate(&models.Diff{}, &models.Model{}, &models.CommitMessage{}, &models.Duel{}, &models.Result{})

	return Database{db: db}
}
