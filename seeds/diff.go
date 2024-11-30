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
		{

			ContentPath: "9.txt",
		},
		{

			ContentPath: "10.txt",
		},
		{
			ContentPath: "11.txt",
		},
		{

			ContentPath: "12.txt",
		},
		{

			ContentPath: "13.txt",
		},
		{

			ContentPath: "14.txt",
		},
		{

			ContentPath: "15.txt",
		},
		{

			ContentPath: "16.txt",
		},
		{

			ContentPath: "17.txt",
		},
		{

			ContentPath: "18.txt",
		},
		{

			ContentPath: "19.txt",
		},
		{

			ContentPath: "20.txt",
		},
		{
			ContentPath: "21.txt",
		},
		{

			ContentPath: "22.txt",
		},
		{

			ContentPath: "23.txt",
		},
		{

			ContentPath: "24.txt",
		},
		{

			ContentPath: "25.txt",
		},
		{

			ContentPath: "26.txt",
		},
		{

			ContentPath: "27.txt",
		},
		{

			ContentPath: "28.txt",
		},
		{

			ContentPath: "29.txt",
		},
		{

			ContentPath: "30.txt",
		},
	}
	_ = database.Db.SaveDiffs(&diffs)
}
