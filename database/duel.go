package database

import (
	"commitinder/models"
	"fmt"
)

func (d *Database) SaveDuel(duel *models.Duel) error {
	err := d.db.Create(&duel).Error
	if err != nil {
		return fmt.Errorf("error saving the duel: " + err.Error())
	}
	return nil
}
func (d *Database) SaveDuels(duels *[]models.Duel) error {
	err := d.db.Create(&duels).Error
	if err != nil {
		return fmt.Errorf("error saving new duels: " + err.Error())
	}
	return nil
}
func (d *Database) GetDuel(duelId int, duel *models.Duel) error {
	err := d.db.Preload("Users").First(&duel, "id = ?", duelId).Error
	if err != nil {
		return fmt.Errorf("error getting a duel, by id: " + err.Error())
	}
	return nil
}
func (d *Database) GetAllDuels(duels *[]models.Duel) error {
	err := d.db.Preload("Users").Find(&duels).Error
	if err != nil {
		return fmt.Errorf("error getting all duels: " + err.Error())
	}
	return nil
}
func (d *Database) UpdateDuel(duel *models.Duel) error {
	err := d.db.Preload("Users").Updates(&duel).Error
	if err != nil {
		return fmt.Errorf("error updating a duel: " + err.Error())
	}
	return nil
}
func (d *Database) UpdateDuels(duels *[]models.Duel) error {
	err := d.db.Preload("Users").Updates(&duels).Error
	if err != nil {
		return fmt.Errorf("error updating a duel: " + err.Error())
	}
	return nil
}
func (d *Database) DeleteDuel(duelId int) error {
	err := d.db.Delete(&models.Duel{}, "id = ?", duelId).Error
	if err != nil {
		return fmt.Errorf("error deleting the duel: " + err.Error())
	}
	return nil
}