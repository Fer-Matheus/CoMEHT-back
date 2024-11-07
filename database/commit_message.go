package database

import (
	"commitinder/models"
	"fmt"
)

func (d *Database) SaveCommitMessage(commitMessage *models.CommitMessage) error {
	err := d.db.Preload("Models").Preload("Diffs").Create(&commitMessage).Error
	if err != nil {
		return fmt.Errorf("error saving a new commit message: " + err.Error())
	}
	return nil
}
func (d *Database) GetCommitMessage(commitMessageId int, commitMessage *models.CommitMessage) error {
	err := d.db.Preload("Diffs").Preload("Models").First(&commitMessage, "id = ?", commitMessageId).Error
	if err != nil {
		return fmt.Errorf("error getting a commit message by id: " + err.Error())
	}
	return nil
}
func (d *Database) GetAllCommitMessages(commitMessages *[]models.CommitMessage) error {
	err := d.db.Preload("Diffs").Preload("Models").Find(&commitMessages).Error
	if err != nil {
		return fmt.Errorf("error getting all commit messages: " + err.Error())
	}
	return nil
}
func (d *Database) UpdateCommitMessage(commitMessage *models.CommitMessage) error {
	err := d.db.Preload("Diffs").Preload("Models").Updates(&commitMessage).Error
	if err != nil {
		return fmt.Errorf("error updating a commit message: " + err.Error())
	}
	return nil
}
func (d *Database) DeleteCommitMessage(commitMessageId int) error {
	err := d.db.Delete(&models.CommitMessage{}, "id = ?", commitMessageId).Error
	if err != nil {
		return fmt.Errorf("error deleting a commit message: " + err.Error())
	}
	return nil
}
