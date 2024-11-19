package views

import "commitinder/models"

type DuelResponse struct {
	DuelId         int    `json:"duel_id"`
	DiffContent    string `json:"diff_content"`
	CommitMessageA string `json:"commit_message_a"`
	CommitMessageB string `json:"commit_message_b"`
}

func (d *DuelResponse) FromModelToView(duel *models.Duel) {
	d.DuelId = duel.Id
	d.DiffContent = duel.CommitMessageA.Diff.Content
	d.CommitMessageA = duel.CommitMessageA.Message
	d.CommitMessageB = duel.CommitMessageB.Message
}

type Options struct {
	Aspect       string  `gorm:"aspect"`
	ChoiseTime   float64 `json:"choise_time"`
	ChosenOption string  `json:"chosen_option"`
}

type ResultsRequest struct {
	DuelId  int       `json:"duel_id"`
	Options []Options `json:"options"`
}
