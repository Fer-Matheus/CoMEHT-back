package database

import (
	"commitinder/models"
	"fmt"
)

func (d *Database) SaveModel(model *models.Model) error {
	err := d.db.Create(&model).Error
	if err != nil {
		return fmt.Errorf("error saving a new model: " + err.Error())
	}
	return nil
}
func (d *Database) GetModel(modelId int, model *models.Model) error {
	err := d.db.Preload("CommitMessage").First(&model, "id = ?", modelId).Error
	if err != nil {
		return fmt.Errorf("error getting a model by id: " + err.Error())
	}
	return nil
}
func (d *Database) GetAllModels(models *[]models.Model) error {
	err := d.db.Preload("CommitMessages").Find(&models).Error
	if err != nil {
		return fmt.Errorf("error getting all models: " + err.Error())
	}
	return nil
}
func (d *Database) UpdateModel(model *models.Model) error {
	err := d.db.Preload("CommitMessage").Updates(&model).Error
	if err != nil {
		return fmt.Errorf("error updating a model: " + err.Error())
	}
	return nil
}
func (d *Database) DeleteModel(modelId int) error {
	err := d.db.Delete(&models.Model{}, "id = ?", modelId).Error
	if err != nil {
		return fmt.Errorf("error deleting a model, by id" + err.Error())
	}
	return nil
}
