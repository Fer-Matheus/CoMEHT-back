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
		{
			ContentPath: "61.txt",
		},
		{
			ContentPath: "62.txt",
		},
		{
			ContentPath: "63.txt",
		},
		{
			ContentPath: "64.txt",
		},
		{
			ContentPath: "65.txt",
		},
		{
			ContentPath: "66.txt",
		},
		{
			ContentPath: "67.txt",
		},
		{
			ContentPath: "68.txt",
		},
		{
			ContentPath: "69.txt",
		},
		{
			ContentPath: "70.txt",
		},
		{
			ContentPath: "71.txt",
		},
		{
			ContentPath: "72.txt",
		},
		{
			ContentPath: "73.txt",
		},
		{
			ContentPath: "74.txt",
		},
		{
			ContentPath: "75.txt",
		},
		{
			ContentPath: "76.txt",
		},
		{
			ContentPath: "77.txt",
		},
		{
			ContentPath: "78.txt",
		},
		{
			ContentPath: "79.txt",
		},
		{
			ContentPath: "80.txt",
		},
		{
			ContentPath: "81.txt",
		},
		{
			ContentPath: "82.txt",
		},
		{
			ContentPath: "83.txt",
		},
		{
			ContentPath: "84.txt",
		},
		{
			ContentPath: "85.txt",
		},
		{
			ContentPath: "86.txt",
		},
		{
			ContentPath: "87.txt",
		},
		{
			ContentPath: "88.txt",
		},
		{
			ContentPath: "89.txt",
		},
		{
			ContentPath: "90.txt",
		},
		{
			ContentPath: "91.txt",
		},
		{
			ContentPath: "92.txt",
		},
		{
			ContentPath: "93.txt",
		},
		{
			ContentPath: "94.txt",
		},
		{
			ContentPath: "95.txt",
		},
		{
			ContentPath: "96.txt",
		},
		{
			ContentPath: "97.txt",
		},
		{
			ContentPath: "98.txt",
		},
		{
			ContentPath: "99.txt",
		},
		{
			ContentPath: "100.txt",
		},
		{
			ContentPath: "101.txt",
		},
		{
			ContentPath: "102.txt",
		},
		{
			ContentPath: "103.txt",
		},
		{
			ContentPath: "104.txt",
		},
		{
			ContentPath: "105.txt",
		},
		{
			ContentPath: "106.txt",
		},
		{
			ContentPath: "107.txt",
		},
		{
			ContentPath: "108.txt",
		},
		{
			ContentPath: "109.txt",
		},
		{
			ContentPath: "110.txt",
		},
		{
			ContentPath: "111.txt",
		},
		{
			ContentPath: "112.txt",
		},
		{
			ContentPath: "113.txt",
		},
		{
			ContentPath: "114.txt",
		},
		{
			ContentPath: "115.txt",
		},
		{
			ContentPath: "116.txt",
		},
		{
			ContentPath: "117.txt",
		},
		{
			ContentPath: "118.txt",
		},
		{
			ContentPath: "119.txt",
		},
		{
			ContentPath: "120.txt",
		},
	}
	_ = database.Db.SaveDiffs(&diffs)
}
