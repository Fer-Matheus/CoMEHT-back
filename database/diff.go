package database

import (
	"commitinder/models"
	"fmt"
)

func (d *Database) SaveDiff(diff *models.Diff) error {
	err := d.db.Create(&diff).Error
	if err != nil {
		return fmt.Errorf("error saving a new diff: " + err.Error())
	}
	return nil
}
func (d *Database) GetDiff(diffId int, diff *models.Diff) error {
	err := d.db.Preload("CommitMessage").First(&diff, "id = ?", diffId).Error
	if err != nil {
		return fmt.Errorf("error getting a diff by id: " + err.Error())
	}
	return nil
}
func (d *Database) GetAllDiffs(diffs *[]models.Diff) error {
	err := d.db.Preload("CommitMessages").Find(&diffs).Error
	if err != nil {
		return fmt.Errorf("error getting all diffs: " + err.Error())
	}
	return nil
}
func (d *Database) UpdateDiff(diff *models.Diff) error {
	err := d.db.Preload("CommitMessage").Updates(&diff).Error
	if err != nil {
		return fmt.Errorf("error updating a diff: " + err.Error())
	}
	return nil
}
func (d *Database) DeleteDiff(diffId int) error {
	err := d.db.Delete(&models.Diff{}, "id = ?", diffId).Error
	if err != nil {
		return fmt.Errorf("error deleting a diff, by id" + err.Error())
	}
	return nil
}