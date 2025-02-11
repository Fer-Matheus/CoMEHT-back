package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedDuels(userId int) {

	var user models.User
	_ = database.Db.GetUserById(userId, &user)

	var duels []models.Duel

	_ = database.Db.GetAllDuels(&duels)

	if len(duels) > 0 {
		for _, duel := range duels {
			duel.Users = append(duel.Users, user)
			database.Db.UpdateDuel(&duel)
		}
		return
	}

	users := []models.User{user}

	seeds := []models.Duel{
		{Users: users, CommitMessageAId: 1, CommitMessageBId: 2},
		{Users: users, CommitMessageAId: 7, CommitMessageBId: 5},
		{Users: users, CommitMessageAId: 11, CommitMessageBId: 10},
		{Users: users, CommitMessageAId: 13, CommitMessageBId: 16},
		{Users: users, CommitMessageAId: 18, CommitMessageBId: 20},
		{Users: users, CommitMessageAId: 24, CommitMessageBId: 23},
		{Users: users, CommitMessageAId: 27, CommitMessageBId: 28},
		{Users: users, CommitMessageAId: 30, CommitMessageBId: 32},
		{Users: users, CommitMessageAId: 35, CommitMessageBId: 33},
		{Users: users, CommitMessageAId: 39, CommitMessageBId: 40},
		{Users: users, CommitMessageAId: 44, CommitMessageBId: 42},
		{Users: users, CommitMessageAId: 47, CommitMessageBId: 45},
		{Users: users, CommitMessageAId: 51, CommitMessageBId: 49},
		{Users: users, CommitMessageAId: 53, CommitMessageBId: 55},
		{Users: users, CommitMessageAId: 57, CommitMessageBId: 58},
		{Users: users, CommitMessageAId: 62, CommitMessageBId: 63},
		{Users: users, CommitMessageAId: 66, CommitMessageBId: 68},
		{Users: users, CommitMessageAId: 69, CommitMessageBId: 71},
		{Users: users, CommitMessageAId: 73, CommitMessageBId: 74},
		{Users: users, CommitMessageAId: 78, CommitMessageBId: 80},
		{Users: users, CommitMessageAId: 82, CommitMessageBId: 84},
		{Users: users, CommitMessageAId: 86, CommitMessageBId: 85},
		{Users: users, CommitMessageAId: 90, CommitMessageBId: 92},
		{Users: users, CommitMessageAId: 93, CommitMessageBId: 94},
		{Users: users, CommitMessageAId: 97, CommitMessageBId: 98},
		{Users: users, CommitMessageAId: 103, CommitMessageBId: 104},
		{Users: users, CommitMessageAId: 105, CommitMessageBId: 107},
		{Users: users, CommitMessageAId: 112, CommitMessageBId: 110},
		{Users: users, CommitMessageAId: 115, CommitMessageBId: 113},
		{Users: users, CommitMessageAId: 117, CommitMessageBId: 118},
		{Users: users, CommitMessageAId: 121, CommitMessageBId: 123},
		{Users: users, CommitMessageAId: 127, CommitMessageBId: 128},
		{Users: users, CommitMessageAId: 130, CommitMessageBId: 129},
		{Users: users, CommitMessageAId: 135, CommitMessageBId: 136},
		{Users: users, CommitMessageAId: 140, CommitMessageBId: 137},
		{Users: users, CommitMessageAId: 141, CommitMessageBId: 142},
		{Users: users, CommitMessageAId: 148, CommitMessageBId: 145},
		{Users: users, CommitMessageAId: 150, CommitMessageBId: 149},
		{Users: users, CommitMessageAId: 154, CommitMessageBId: 155},
		{Users: users, CommitMessageAId: 157, CommitMessageBId: 160},
		{Users: users, CommitMessageAId: 163, CommitMessageBId: 161},
		{Users: users, CommitMessageAId: 168, CommitMessageBId: 166},
		{Users: users, CommitMessageAId: 171, CommitMessageBId: 172},
		{Users: users, CommitMessageAId: 174, CommitMessageBId: 176},
		{Users: users, CommitMessageAId: 179, CommitMessageBId: 178},
		{Users: users, CommitMessageAId: 182, CommitMessageBId: 181},
		{Users: users, CommitMessageAId: 187, CommitMessageBId: 185},
		{Users: users, CommitMessageAId: 190, CommitMessageBId: 189},
		{Users: users, CommitMessageAId: 196, CommitMessageBId: 193},
		{Users: users, CommitMessageAId: 200, CommitMessageBId: 198},
		{Users: users, CommitMessageAId: 204, CommitMessageBId: 201},
		{Users: users, CommitMessageAId: 205, CommitMessageBId: 208},
		{Users: users, CommitMessageAId: 210, CommitMessageBId: 211},
		{Users: users, CommitMessageAId: 213, CommitMessageBId: 215},
		{Users: users, CommitMessageAId: 217, CommitMessageBId: 219},
		{Users: users, CommitMessageAId: 224, CommitMessageBId: 222},
		{Users: users, CommitMessageAId: 227, CommitMessageBId: 226},
		{Users: users, CommitMessageAId: 230, CommitMessageBId: 232},
		{Users: users, CommitMessageAId: 234, CommitMessageBId: 236},
		{Users: users, CommitMessageAId: 237, CommitMessageBId: 240},
		{Users: users, CommitMessageAId: 242, CommitMessageBId: 241},
		{Users: users, CommitMessageAId: 248, CommitMessageBId: 245},
		{Users: users, CommitMessageAId: 252, CommitMessageBId: 251},
		{Users: users, CommitMessageAId: 256, CommitMessageBId: 255},
		{Users: users, CommitMessageAId: 260, CommitMessageBId: 258},
		{Users: users, CommitMessageAId: 264, CommitMessageBId: 262},
		{Users: users, CommitMessageAId: 265, CommitMessageBId: 267},
		{Users: users, CommitMessageAId: 269, CommitMessageBId: 272},
		{Users: users, CommitMessageAId: 273, CommitMessageBId: 274},
		{Users: users, CommitMessageAId: 280, CommitMessageBId: 278},
		{Users: users, CommitMessageAId: 283, CommitMessageBId: 281},
		{Users: users, CommitMessageAId: 288, CommitMessageBId: 286},
		{Users: users, CommitMessageAId: 291, CommitMessageBId: 292},
		{Users: users, CommitMessageAId: 294, CommitMessageBId: 293},
		{Users: users, CommitMessageAId: 298, CommitMessageBId: 297},
		{Users: users, CommitMessageAId: 302, CommitMessageBId: 301},
		{Users: users, CommitMessageAId: 308, CommitMessageBId: 306},
		{Users: users, CommitMessageAId: 312, CommitMessageBId: 311},
		{Users: users, CommitMessageAId: 316, CommitMessageBId: 314},
		{Users: users, CommitMessageAId: 320, CommitMessageBId: 317},
		{Users: users, CommitMessageAId: 324, CommitMessageBId: 322},
		{Users: users, CommitMessageAId: 327, CommitMessageBId: 326},
		{Users: users, CommitMessageAId: 332, CommitMessageBId: 330},
		{Users: users, CommitMessageAId: 335, CommitMessageBId: 333},
		{Users: users, CommitMessageAId: 337, CommitMessageBId: 339},
		{Users: users, CommitMessageAId: 343, CommitMessageBId: 342},
		{Users: users, CommitMessageAId: 346, CommitMessageBId: 348},
		{Users: users, CommitMessageAId: 349, CommitMessageBId: 351},
		{Users: users, CommitMessageAId: 356, CommitMessageBId: 355},
		{Users: users, CommitMessageAId: 359, CommitMessageBId: 360},
		{Users: users, CommitMessageAId: 362, CommitMessageBId: 363},
		{Users: users, CommitMessageAId: 366, CommitMessageBId: 368},
		{Users: users, CommitMessageAId: 371, CommitMessageBId: 370},
		{Users: users, CommitMessageAId: 376, CommitMessageBId: 373},
		{Users: users, CommitMessageAId: 379, CommitMessageBId: 378},
		{Users: users, CommitMessageAId: 383, CommitMessageBId: 381},
		{Users: users, CommitMessageAId: 388, CommitMessageBId: 386},
		{Users: users, CommitMessageAId: 391, CommitMessageBId: 390},
		{Users: users, CommitMessageAId: 395, CommitMessageBId: 394},
		{Users: users, CommitMessageAId: 398, CommitMessageBId: 399},
		{Users: users, CommitMessageAId: 401, CommitMessageBId: 402},
		{Users: users, CommitMessageAId: 408, CommitMessageBId: 405},
		{Users: users, CommitMessageAId: 411, CommitMessageBId: 412},
		{Users: users, CommitMessageAId: 413, CommitMessageBId: 414},
		{Users: users, CommitMessageAId: 417, CommitMessageBId: 418},
		{Users: users, CommitMessageAId: 424, CommitMessageBId: 421},
		{Users: users, CommitMessageAId: 427, CommitMessageBId: 425},
		{Users: users, CommitMessageAId: 430, CommitMessageBId: 429},
		{Users: users, CommitMessageAId: 436, CommitMessageBId: 435},
		{Users: users, CommitMessageAId: 440, CommitMessageBId: 437},
		{Users: users, CommitMessageAId: 444, CommitMessageBId: 443},
		{Users: users, CommitMessageAId: 447, CommitMessageBId: 445},
		{Users: users, CommitMessageAId: 452, CommitMessageBId: 451},
		{Users: users, CommitMessageAId: 456, CommitMessageBId: 455},
		{Users: users, CommitMessageAId: 459, CommitMessageBId: 460},
		{Users: users, CommitMessageAId: 461, CommitMessageBId: 463},
		{Users: users, CommitMessageAId: 467, CommitMessageBId: 468},
		{Users: users, CommitMessageAId: 471, CommitMessageBId: 469},
		{Users: users, CommitMessageAId: 473, CommitMessageBId: 475},
		{Users: users, CommitMessageAId: 477, CommitMessageBId: 479}}
	_ = database.Db.SaveDuels(&seeds)
}
