package views

import "comeht/models"

type ModelRequest struct {
	Name string `json:"name"`
}

func (m *ModelRequest) ToModel() (model models.Model) {
	model.Name = m.Name
	return model
}

type ModelResponse struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	CommitMessages []int  `json:"commit_messages"`
}

func (m *ModelResponse) FromModelToView(model *models.Model) {
	m.Id = model.Id
	m.Name = model.Name
	for _, commitMessage := range model.CommitMessages {
		m.CommitMessages = append(m.CommitMessages, commitMessage.Id)
	}
}

type ModelsResponse struct {
	Models []ModelResponse `json:"models"`
}
