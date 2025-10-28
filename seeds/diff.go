package seeds

import (
	"comeht/database"
	"comeht/models"
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
		{
			ContentPath: "31.txt",
		},
		{

			ContentPath: "32.txt",
		},
		{

			ContentPath: "33.txt",
		},
		{

			ContentPath: "34.txt",
		},
		{

			ContentPath: "35.txt",
		},
		{

			ContentPath: "36.txt",
		},
		{

			ContentPath: "37.txt",
		},
		{

			ContentPath: "38.txt",
		},
		{

			ContentPath: "39.txt",
		},
		{

			ContentPath: "40.txt",
		},
		{
			ContentPath: "41.txt",
		},
		{

			ContentPath: "42.txt",
		},
		{

			ContentPath: "43.txt",
		},
		{

			ContentPath: "44.txt",
		},
		{

			ContentPath: "45.txt",
		},
		{

			ContentPath: "46.txt",
		},
		{

			ContentPath: "47.txt",
		},
		{

			ContentPath: "48.txt",
		},
		{

			ContentPath: "49.txt",
		},
		{

			ContentPath: "50.txt",
		},
		{
			ContentPath: "51.txt",
		},
		{

			ContentPath: "52.txt",
		},
		{

			ContentPath: "53.txt",
		},
		{

			ContentPath: "54.txt",
		},
		{

			ContentPath: "55.txt",
		},
		{

			ContentPath: "56.txt",
		},
		{

			ContentPath: "57.txt",
		},
		{

			ContentPath: "58.txt",
		},
		{

			ContentPath: "59.txt",
		},
		{

			ContentPath: "60.txt",
		},
	}
	_ = database.Db.SaveDiffs(&diffs)
}
