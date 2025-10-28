package database

import (
	"comeht/models"
	"fmt"
)

func (d *Database) SaveUser(user *models.User) error {
	err := d.db.Create(&user).Error
	if err != nil {
		return fmt.Errorf("error saving a new user: " + err.Error())
	}
	return nil
}
func (d *Database) SaveUsers(users *[]models.User) error {
	err := d.db.Create(&users).Error
	if err != nil {
		return fmt.Errorf("error saving new users: " + err.Error())
	}
	return nil
}
func (d *Database) GetUser(username string, user *models.User) error {
	err := d.db.First(&user, "username = ?", username).Error
	if err != nil {
		return fmt.Errorf("error getting a user by username: " + err.Error())
	}
	return nil
}
func (d *Database) GetUserById(userId int, user *models.User) error {
	err := d.db.First(&user, "id = ?", userId).Error
	if err != nil {
		return fmt.Errorf("error getting a user by username: " + err.Error())
	}
	return nil
}
func (d *Database) GetAllUser(users *[]models.User) error {
	err := d.db.Find(&users).Error
	if err != nil {
		return fmt.Errorf("error getting all users: " + err.Error())
	}
	return nil
}
func (d *Database) UpdateUser(user *models.User) error {
	err := d.db.Where("id = ?", user.Id).Updates(&user).Error
	if err != nil {
		return fmt.Errorf("error updating a user: " + err.Error())
	}
	return nil
}
func (d *Database) DeleteUser(username string) error {
	err := d.db.Delete(&models.Diff{}, "username = ?", username).Error
	if err != nil {
		return fmt.Errorf("error deleting a user, by username" + err.Error())
	}
	return nil
}
