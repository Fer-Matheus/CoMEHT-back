package views

import "comeht/models"

type CommitMessageRequest struct {
	Message string `json:"message"`
	ModelId int    `json:"model_id"`
	DiffId  int    `json:"diff_id"`
}

func (c *CommitMessageRequest) ToModel() (commitMessage models.CommitMessage) {
	commitMessage.Message = c.Message
	commitMessage.DiffId = c.DiffId
	commitMessage.ModelId = c.ModelId
	return commitMessage
}

type CommitMessageResponse struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	ModelId int    `json:"model_id"`
	DiffId  int    `json:"diff_id"`
}

func (c *CommitMessageResponse) FromModelToView(commitMessage *models.CommitMessage) {
	c.Id = commitMessage.Id
	c.Message = commitMessage.Message
	c.ModelId = commitMessage.ModelId
	c.DiffId = commitMessage.DiffId
}

type CommitMessagesResponse struct {
	CommitMessages []CommitMessageResponse `json:"commit_messages"`
}
