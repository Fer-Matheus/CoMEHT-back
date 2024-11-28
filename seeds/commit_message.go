package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedCommitMessages() {
	seeds := []models.CommitMessage{
		{
			Message: "Revert \"Fix unnecessary type casting in Test\" This reverts commit f3ab77c.",
			DiffId:  1,
			ModelId: 1,
		},
		{
			Message: "remove unused parameter",
			DiffId:  1,
			ModelId: 2,
		},
		{
			Message: "Refactor: Specify type of HashMap in `testTimelineMetricCachePrecisionUpdates()`In the `testTimelineMetricCachePrecisionUpdates()` method in the `TimelineMetricCacheTest.java` file, the type of the HashMap `valueMap` was specified in its initialization. This change improves the code's internal structure without altering its external behavior.",
			DiffId:  1,
			ModelId: 3,
		},
		{
			Message: "Added some javadocs for OS condition.",
			DiffId:  2,
			ModelId: 1,
		},
		{
			Message: "add javadoc for Os setFamily",
			DiffId:  2,
			ModelId: 2,
		},
		{
			Message: "Feat: Improve documentation in Os.java Added detailed comments to the setFamily and eval methods in the Os.java file. The comments provide information about the functionality of the methods and the expected input for the setFamily method.",
			DiffId:  2,
			ModelId: 3,
		},
		{
			Message: "Last weeks changes took away helpful information about why exec failed to run; this puts it back in.",
			DiffId:  3,
			ModelId: 1,
		},
		{
			Message: "improve error message",
			DiffId:  3,
			ModelId: 2,
		},
		{
			Message: "Fix: Provide more detailed exception message in `runExec` method In the `runExec` method of `ExecTask.java`, the exception message when the command line process execution fails has been updated to include more information about the exception. This change improves the error handling by providing more context about the exception, which can be helpful for debugging.",
			DiffId:  3,
			ModelId: 3,
		},
		{
			Message: "Made complex loop counter a simple constant Removing a Audit violation",
			DiffId:  4,
			ModelId: 1,
		},
		{
			Message: "fix a bug in XMLValidateTask",
			DiffId:  4,
			ModelId: 2,
		},
		{
			Message: "Refactor: Optimize loop in XMLValidateTask This commit introduces a minor optimization in the XMLValidateTask class. A new variable 'size' is introduced to hold the size of 'm_dtdLocations', which is then used in the for loop condition. This change improves the structure of the code without altering its functionality.",
			DiffId:  4,
			ModelId: 3,
		},
		{
			Message: "testcase wouldn't compile - slipping fingers?",
			DiffId:  5,
			ModelId: 1,
		},
		{
			Message: "fix broken test",
			DiffId:  5,
			ModelId: 2,
		},
		{
			Message: "Fix: Correct object used to set date in testDateISO() method In the 'testDateISO()' method of the 'DateUtilsTest.java' file, the object used to set the calendar date has been corrected from 'emailtaskcal' to 'cal'. This change ensures that the method tests the conversion of a date to an ISO 8601 string correctly.",
			DiffId:  5,
			ModelId: 3,
		},
		{
			Message: "Delete some unused variables",
			DiffId:  6,
			ModelId: 1,
		},
		{
			Message: "remove unused field",
			DiffId:  6,
			ModelId: 2,
		},
		{
			Message: "Refactor: Remove unused variables in Translate.java In both the 'proposal/myrmidon/src/main/org/apache/tools/ant/taskdefs/optional/i18n/Translate.java' and 'proposal/myrmidon/src/todo/org/apache/tools/ant/taskdefs/optional/i18n/Translate.java' files, the private Locale variable 'locale' and the StringBuffer 'newline' in the 'translate' method have been removed. These variables were not used, and their removal simplifies the code without affecting its functionality.",
			DiffId:  6,
			ModelId: 3,
		},
		{
			Message: "quick fix for Ant.java use of property in Ant#reinit",
			DiffId:  7,
			ModelId: 1,
		},
		{
			Message: "fix a bug in ComponentHelper ",
			DiffId:  7,
			ModelId: 2,
		},
		{
			Message: "Fix: Handle null task in createTask method In the createTask method of the ComponentHelper class, added a condition to handle a case where the task is null and the taskType equals \"property\". In this case, a new task definition is added and a new task is created. This is a fix for a situation where the \"property\" task is used before the project is initialized.",
			DiffId:  7,
			ModelId: 3,
		},
		{
			Message: "cancel previous change",
			DiffId:  8,
			ModelId: 1,
		},
		{
			Message: "fix a bug in the Translate",
			DiffId:  8,
			ModelId: 2,
		},
		{
			Message: "Fix: Update resource map loading in Translate.java This commit updates the method loadResourceMaps() in Translate.java to load resource maps for a specific locale and the default locale, processing each bundle file in a specific order. The change enhances the specificity of the resource map loading process.",
			DiffId:  8,
			ModelId: 3,
		},
		{
			Message: "Fix incorrect finalize override",
			DiffId:  9,
			ModelId: 1,
		},
		{
			Message: "make finalize public",
			DiffId:  9,
			ModelId: 2,
		},
		{
			Message: "Refactor: Change visibility of finalize method in CBZip2OutputStream The visibility of the `finalize` method in `CBZip2OutputStream.java` has been changed from `public` to `protected`, and a call to `super.finalize()` has been added. This ensures that the object is properly closed when it's garbage collected.",
			DiffId:  9,
			ModelId: 3,
		},
		{
			Message: "mistake found ... oops ... Corrected",
			DiffId:  10,
			ModelId: 1,
		},
		{
			Message: "fix the build",
			DiffId:  10,
			ModelId: 2,
		},
		{
			Message: "Fix: Correct classification of symbolic links in FTP.java In the method 'scandir' of the class 'FTP.java', symbolic links were previously classified as directories and added to 'dirsExcluded'. This commit corrects this misclassification by adding symbolic links to 'filesExcluded' instead. This change ensures that symbolic links are correctly identified and handled.",
			DiffId:  10,
			ModelId: 3,
		},
	}

	_ = database.Db.SaveCommitMessages(&seeds)

}
