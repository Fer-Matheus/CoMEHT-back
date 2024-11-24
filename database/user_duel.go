package database

import (
	"commitinder/models"
	"fmt"
)

func (d *Database) SaveUserDuel(userDuel *models.UserDuel) error {
	err := d.db.Create(&userDuel).Error
	if err != nil {
		return fmt.Errorf("error saving the user duel: %s", err)
	}
	return nil
}

func (d *Database) SaveResults(userDuel *models.UserDuel) error {

	var count int64
	_ = d.db.Find(&models.Result{}, "user_id = ? AND duel_id = ?", userDuel.UserId, userDuel.DuelId).Count(&count)

	if count > 0 {
		return nil
	}

	err := d.db.Select("Results").Updates(&userDuel).Error
	if err != nil {
		return fmt.Errorf("error updating the user duel: %s", err)
	}
	return nil
}

func (d *Database) GetAllResults(results *[]models.Result) error {
	
	err := d.db.Find(&results).Error
	if err != nil {
		return fmt.Errorf("error getting all results: %s", err)
	}

	return nil
}