package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedDiffs() {

	diffs := []models.Diff{
		{
			ContentPath: "1.txt",
		},
		{

			ContentPath: "2.txt",
		},
		{

			ContentPath: "3.txt",
		},
		{

			ContentPath: "4.txt",
		},
		{

			ContentPath: "5.txt",
		},
		{

			ContentPath: "6.txt",
		},
		{

			ContentPath: "7.txt",
		},
		{

			ContentPath: "8.txt",
		},
	}
	_ = database.Db.SaveDiffs(&diffs)
}
