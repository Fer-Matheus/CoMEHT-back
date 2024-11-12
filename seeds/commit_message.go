package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedCommitMessages(){
	seeds := []models.CommitMessage{
      {
        Id: 1,
        Message:
          "Update CHANGELOG: Fix STORM-558 by changing 'swap!' to 'reset!' in supervisor",
        DiffId: 1,
        ModelId: 1,
      },
      {
        Id: 2,
        Message: "Fix STORM-558",
        DiffId: 1,
        ModelId: 2,
      },
      {
        Id: 3,
        Message:
          "Refactor environment handling and add client environment info\n          Renamed getCommandlineWhitelistedClientEnv() to getWhitelistedClientEnv().\n          Added ClientEnv info item to capture and display the current client environment.\n          Updated tests to cover environment variable freezing (test_env_freezing).",
        ModelId: 1,
        DiffId: 2,
      },
      {
        Id: 4,
        Message:
          "Add an info item to show the currently inherited client environment .",
        DiffId: 2,
        ModelId: 3,
      },
	}

	_ = database.Db.SaveCommitMessages(&seeds)
	
}