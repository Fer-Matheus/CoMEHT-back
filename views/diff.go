package views

import "comeht/models"

type DiffRequest struct {
	Content string `json:"content"`
}

func (d *DiffRequest) ToModel() (diff models.Diff) {
	diff.ContentPath = d.Content
	return diff
}

type DiffResponse struct {
	Id             int    `json:"id"`
	Content        string `json:"content"`
	CommitMessages []int  `json:"commit_messages"`
}

func (d *DiffResponse) FromModelToView(diff *models.Diff) {
	d.Id = diff.Id
	d.Content = diff.ContentPath

	for _, commitMessage := range diff.CommitMessages {
		d.CommitMessages = append(d.CommitMessages, commitMessage.Id)
	}
}

type DiffsResponse struct {
	Diffs []DiffResponse `json:"diffs"`
}
