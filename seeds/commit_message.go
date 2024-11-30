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
		{
			Message: "typo - sort of",
			DiffId:  11,
			ModelId: 1,
		},
		{
			Message: "fix typo in javadoc",
			DiffId:  11,
			ModelId: 2,
		},
		{
			Message: "Style: Corrected language in comment Changed the language of a comment in the StyleTest.java file from German to English for better understanding.",
			DiffId:  11,
			ModelId: 3,
		},
		{
			Message: "Make SQL task work under Oracle",
			DiffId:  12,
			ModelId: 1,
		},
		{
			Message: "remove unnecessary code",
			DiffId:  12,
			ModelId: 2,
		},
		{
			Message: "Refactor: Modify loop condition in execSQL method This commit modifies the condition in the do-while loop in the execSQL method of the SQLExec.java file. The change is a refactor aimed at improving the structure or readability of the code. The execSQL method is responsible for executing a SQL statement and handling any SQL warnings or errors. The modification might affect how the method handles the SQL execution.",
			DiffId:  12,
			ModelId: 3,
		},
		{
			Message: "Allow references and properties to be null without causing NPE in BSF",
			DiffId:  13,
			ModelId: 1,
		},
		{
			Message: "fix npe in ScriptRunner",
			DiffId:  13,
			ModelId: 2,
		},
		{
			Message: "Fix: Add null check before declaring a bean in ScriptRunner In the `executeScript(String execName)` method of `ScriptRunner.java`, a null check is added before declaring a bean. If the value is null, the bean is undeclared. This change is made to ensure that the script is executed using the beanshell scripting framework.",
			DiffId:  13,
			ModelId: 3,
		},
		{
			Message: "raise log level of targets to match that of task events",
			DiffId:  14,
			ModelId: 1,
		},
		{
			Message: "fix checkstyle error",
			DiffId:  14,
			ModelId: 2,
		},
		{
			Message: "Refactor: Change logging level in targetStarted method The logging level in the `targetStarted` method of the `CommonsLoggingListener` class has been changed from `Project.MSG_DEBUG` to `Project.MSG_VERBOSE`. This is a refactor commit, meaning the code has been changed to improve its internal structure without altering its external behavior. The functionality of the `targetStarted` method, which logs the start of a target during a build process, remains the same.",
			DiffId:  14,
			ModelId: 3,
		},
		{
			Message: "Remove Steve's complaint ;) after Stefan and Peter's recent Location changes",
			DiffId:  15,
			ModelId: 1,
		},
		{
			Message: "fix a typo in javadoc",
			DiffId:  15,
			ModelId: 2,
		},
		{
			Message: "Style: Refactor comments in EchoXML.java This commit refactors the comments in the EchoXML.java file for better readability and understanding. A line of code has also been removed. The functionality of the EchoXML class, which writes XML content to a specified file or log, remains unchanged.",
			DiffId:  15,
			ModelId: 3,
		},
		{
			Message: "Remove useless logging",
			DiffId:  16,
			ModelId: 1,
		},
		{
			Message: "remove debug logging",
			DiffId:  16,
			ModelId: 2,
		},
		{
			Message: "Refactor: Remove logging statement from processSnapshots method The logging statement \"Processing snapshots.\" has been removed from the processSnapshots method in the DefaultRepositoryBrowsing class. This change is part of a refactoring effort to improve the code quality.",
			DiffId:  16,
			ModelId: 3,
		},
		{
			Message: "add @SInCE for new method",
			DiffId:  17,
			ModelId: 1,
		},
		{
			Message: "add missing javadoc",
			DiffId:  17,
			ModelId: 2,
		},
		{
			Message: "Style: Add comment to `deleteArtifact` method in `RepositoriesService.java` A comment line has been added to the `deleteArtifact` method in the `RepositoriesService.java` file. The comment indicates that permissions are checked in the implementation of the method. This change does not affect the functionality of the code.",
			DiffId:  17,
			ModelId: 3,
		},
		{
			Message: "deal with theoretical case of version scheme changing again",
			DiffId:  18,
			ModelId: 1,
		},
		{
			Message: "catch Throwable instead of DeweyDecimal",
			DiffId:  18,
			ModelId: 2,
		},
		{
			Message: "Refactor: Improved parsing and storing of java version number The parsing and storing of the java version number in JavaEnvUtils.java has been refactored. The java version number is now parsed using the DeweyDecimal class and stored as a string. This change improves the accuracy and readability of the java version number.",
			DiffId:  18,
			ModelId: 3,
		},
		{
			Message: "Correct date formatting",
			DiffId:  19,
			ModelId: 1,
		},
		{
			Message: "fix the test",
			DiffId:  19,
			ModelId: 2,
		},
		{
			Message: "Fix: Corrected date format in PropertyFileTest The testDirect() method in the PropertyFileTest class had an incorrect date format in the assertEquals method. The format was changed from \"%d/%02d/%d\" to \"%d/%02d/%02d\" to correctly represent the date. This change ensures that the test accurately checks the functionality of the PropertyFile class.",
			DiffId:  19,
			ModelId: 3,
		},
		{
			Message: "change required after plexus update",
			DiffId:  20,
			ModelId: 1,
		},
		{
			Message: "renamed getTestConfiguration to getProxyConfiguration",
			DiffId:  20,
			ModelId: 2,
		},
		{
			Message: "Refactor: Rename method to better reflect its functionality In the files `DefaultProxyManagerTest.java` and `LegacyProxyManagerTest.java`, the method `getTestConfiguration()` has been renamed to `getProxyConfiguration()`. This change does not alter the functionality of the `setUp()` method where it is used, but provides a more accurate description of the method's purpose, which is to create and return a proxy configuration for testing purposes.",
			DiffId:  20,
			ModelId: 3,
		},
		{
			Message: "repository scanning test fix for idependance",
			DiffId:  21,
			ModelId: 1,
		},
		{
			Message: "add missing dirtiescontext",
			DiffId:  21,
			ModelId: 2,
		},
		{
			Message: "Feat: Add @DirtiesContext annotation to ArchivaRepositoryScanningTaskExecutorAbstractTest This commit adds the @DirtiesContext annotation to the ArchivaRepositoryScanningTaskExecutorAbstractTest class. This annotation is used to indicate that the Spring ApplicationContext should be dirtied and thus automatically closed after each test method. This is typically used when a test method dirties the context.",
			DiffId:  21,
			ModelId: 3,
		},
		{
			Message: "Adding config file removal before tests",
			DiffId:  22,
			ModelId: 1,
		},
		{
			Message: "add support for delete",
			DiffId:  22,
			ModelId: 2,
		},
		{
			Message: "Feat: Add initialize method in AbstractRepositoryAdminTest This commit adds a new method 'initialize' in the AbstractRepositoryAdminTest.java file. The method is used to set up the test environment before each test. It deletes the existing configuration file if it exists and reloads the configuration. This ensures that the test environment is properly set up for each test case.",
			DiffId:  22,
			ModelId: 3,
		},
		{
			Message: "Fixing repository group implementation",
			DiffId:  23,
			ModelId: 1,
		},
		{
			Message: "fix the build",
			DiffId:  23,
			ModelId: 2,
		},
		{
			Message: "Feat: Update 'createRepositoryGroup' method and add 'setRepositoryRegistry' method This commit updates the 'createRepositoryGroup' method in the 'MavenRepositoryProvider' class to create a new repository group with the given configuration. It also adds a new method 'setRepositoryRegistry' to set the repository registry. These changes enhance the functionality of the 'MavenRepositoryProvider' class.",
			DiffId:  23,
			ModelId: 3,
		},
		{
			Message: "Improve timeout exception when an osgi service can't be found",
			DiffId:  24,
			ModelId: 1,
		},
		{
			Message: "improve error message for ReferenceRecipe",
			DiffId:  24,
			ModelId: 2,
		},
		{
			Message: "Fix: Handle different types of service availability in `getService()` In the `getService()` method of the `ReferenceRecipe.java` file, conditions have been added to handle different types of service availability. If the service is mandatory and not available, an exception is thrown after waiting for a specified timeout period. If the service is optional and not available, an exception is thrown immediately. This fix ensures that the method behaves correctly for different types of service availability.",
			DiffId:  24,
			ModelId: 3,
		},
		{
			Message: "Fix an invalid path on Windows platforms Paths.get calls the platform's Filesystem.getPath, which validates that the path contains legal characters for file names. Since * is an illegal character for Windows, Paths.get can't be used for globs.",
			DiffId:  25,
			ModelId: 1,
		},
		{
			Message: "fix broken test",
			DiffId:  25,
			ModelId: 2,
		},
		{
			Message: "Fix: Update path creation in testMultipleFiles method In the testMultipleFiles method of the AvroSourceTest.java file, the path to the AvroSource has been updated. Previously, the path was created using Paths.get(), but this has been changed to use new File(). This change ensures that the AvroSource can correctly read from multiple Avro files.",
			DiffId:  25,
			ModelId: 3,
		},
		{
			Message: "StreamingWriteFn: check if table exists before creating In StreamingWriteFn, every worker tries to create the table without checking beforehand whether the table exists. This behavior can lead to temporarily violating the BigQuery API quota limits on table.insert:",
			DiffId:  26,
			ModelId: 1,
		},
		{
			Message: "fix bug in BigQueryIO",
			DiffId:  26,
			ModelId: 2,
		},
		{
			Message: "Feat: Modify getOrCreateTable method in BigQueryIO.java This commit modifies the `getOrCreateTable` method in the `BigQueryIO.java` file. The method now checks if a BigQuery table exists based on the provided table specification. If the table does not exist, it creates a new one. This method is thread-safe.",
			DiffId:  26,
			ModelId: 3,
		},
		{
			Message: "Correct URL for subnetwork",
			DiffId:  27,
			ModelId: 1,
		},
		{
			Message: "fix typo in javadoc",
			DiffId:  27,
			ModelId: 2,
		},
		{
			Message: "Refactor: Update GCE subnetwork description and test This commit updates the description of the GCE subnetwork for launching workers in `DataflowPipelineWorkerPoolOptions.java` and modifies the `testSubnetworkConfig()` method in `DataflowPipelineTranslatorTest.java` to test if the subnetwork configuration is correctly set in the pipeline options and properly reflected in the job environment. The changes are part of a refactor and do not alter the functionality of the code.",
			DiffId:  27,
			ModelId: 3,
		},
		{
			Message: "BigQuery: swap from asSingleton to asIterable for Cleanup asIterable can be simpler for runners to implement as it does not require semantically that the PCollection being viewed contains exactly one element.",
			DiffId:  28,
			ModelId: 1,
		},
		{
			Message: "fix quality flaw",
			DiffId:  28,
			ModelId: 2,
		},
		{
			Message: "Refactor: Change handling of cleanup signal in PassThroughThenCleanup In the `expand` method of the `PassThroughThenCleanup.java` file, the handling of the cleanup signal has been refactored. Previously, the cleanup signal was viewed as a singleton with a default value of null. Now, it is viewed as an iterable. This change does not alter the functionality of the method but changes its internal structure.",
			DiffId:  28,
			ModelId: 3,
		},
		{
			Message: "Ignore HBaseIOTest temporarily",
			DiffId:  29,
			ModelId: 1,
		},
		{
			Message: "ignore failing HBaseIOTest",
			DiffId:  29,
			ModelId: 2,
		},
		{
			Message: "Fix: Ignore HBaseIOTest and clean up imports In this commit, the HBaseIOTest class is annotated with @Ignore, indicating that the tests in this class are not to be executed. Additionally, some unused import statements are removed to clean up the code.",
			DiffId:  29,
			ModelId: 3,
		},
		{
			Message: "Fix Regex#FindAll javadoc",
			DiffId:  30,
			ModelId: 1,
		},
		{
			Message: "fix javadoc warnings",
			DiffId:  30,
			ModelId: 2,
		},
		{
			Message: "Feat: Update method documentation and example in Regex.java This commit updates the method documentation and example usage in the `Regex.java` file. The return type in the documentation has been corrected to `PCollection<List<String>>` and the example usage has been updated accordingly.",
			DiffId:  30,
			ModelId: 3,
		},
	}

	_ = database.Db.SaveCommitMessages(&seeds)

}
