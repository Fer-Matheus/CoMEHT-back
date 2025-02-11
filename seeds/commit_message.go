package seeds

import (
	"commitinder/database"
	"commitinder/models"
)

func SeedCommitMessages() {
	seeds := []models.CommitMessage{
		{
			Message: "changed gogrid to not use ssh ctor",
			DiffId:  1,
			ModelId: 1,
		},
		{
			Message: "fix the test",
			DiffId:  1,
			ModelId: 2,
		},
		{
			Message: "Refactor: Update SSH connection in testShellAccess method In the testShellAccess() method of the GoGridLiveTestDisabled.java file, the way an SSH connection is established has been refactored. Instead of creating a new SshjSshClient instance directly, the method now retrieves an instance of SshClient.Factory from the context's injector and uses it to create the SshClient. This change improves the structure of the code without altering its functionality.",
			DiffId:  1,
			ModelId: 3,
		},
		{
			Message: "Added intro image .",
			DiffId:  1,
			ModelId: 4,
		},
		{
			Message: "Added test to check that the key is decoded as expected",
			DiffId:  2,
			ModelId: 1,
		},
		{
			Message: "add unit test for Multimap of",
			DiffId:  2,
			ModelId: 2,
		},
		{
			Message: "Feat: Add test for parsing query with keys requiring decoding This commit adds a new test method `testParseQueryWithKeysThatRequireDecoding` to the `ModifyRequestTest` class. The test checks the parsing of a query with keys that require decoding.",
			DiffId:  2,
			ModelId: 3,
		},
		{
			Message: "Added intro image .",
			DiffId:  2,
			ModelId: 4,
		},
		{
			Message: "Allow lcoation '--mem--' on the command line (TDB, in-memory). Same effect as --memTDB",
			DiffId:  3,
			ModelId: 1,
		},
		{
			Message: "improve exception message",
			DiffId:  3,
			ModelId: 2,
		},
		{
			Message: "Feat: Add in-memory TDB dataset handling in FusekiCmd This commit adds the ability to handle in-memory TDB datasets in the FusekiCmd class. Specifically, the processModulesAndArgs() method has been modified to check if the directory is equal to Names.memName. If so, it logs that the TDB dataset is in-memory. If not, it checks if the directory exists and logs the directory of the TDB dataset. This feature enhances the flexibility of handling TDB datasets.",
			DiffId:  3,
			ModelId: 3,
		},
		{
			Message: "Add TestsuiteSanityTestCase to check testsuite environment .",
			DiffId:  3,
			ModelId: 4,
		},
		{
			Message: "added default list op",
			DiffId:  4,
			ModelId: 1,
		},
		{
			Message: "add list to ELB",
			DiffId:  4,
			ModelId: 2,
		},
		{
			Message: "Feat: Add list method to ELB class This commit introduces a new feature to the ELB class in the labs/elb module. A new method `list` has been added, which takes a `LoadBalancerClient` as an argument and returns an iterable of `LoadBalancer` objects. This method lists all load balancers using the provided `loadBalancerClient` and default options.",
			DiffId:  4,
			ModelId: 3,
		},
		{
			Message: "Added intro image .",
			DiffId:  4,
			ModelId: 4,
		},
		{
			Message: "Remove Nova instance metadata limitRackspace cloud limits instance metadata to 5 key-value pairs, but upstream Nova only sets the limit at 128 by default.  This patch removes the limit entirely; the official python clients don't check it and the server is responsible for enforcing it anyway.",
			DiffId:  5,
			ModelId: 1,
		},
		{
			Message: "remove unused method",
			DiffId:  5,
			ModelId: 2,
		},
		{
			Message: "Fix: Remove limit on metadata key-value pairs in CreateServerOptions The previous implementation of the metadata method in the CreateServerOptions class had a limit on the number of key-value pairs that could be supplied per server. This commit removes that limit, in line with the findings of issue JCLOUDS-507 that Nova instance metadata is not limited to 5 pairs.",
			DiffId:  5,
			ModelId: 3,
		},
		{
			Message: "Added intro image .",
			DiffId:  5,
			ModelId: 4,
		},
		{
			Message: "zone -> region in Cinder AvailabilityZoneApiLiveTest",
			DiffId:  6,
			ModelId: 1,
		},
		{
			Message: "fix unit test",
			DiffId:  6,
			ModelId: 2,
		},
		{
			Message: "Refactor: Update setupContext method in AvailabilityZoneApiLiveTest In the `setupContext()` method of the `AvailabilityZoneApiLiveTest.java` file, the testing context setup has been updated. Previously, the availability zone API was initialized for the first configured zone. Now, the availability zone API is initialized for a specified region. This change improves the structure of the code without altering its behavior.",
			DiffId:  6,
			ModelId: 3,
		},
		{
			Message: "Remove unused import",
			DiffId:  6,
			ModelId: 4,
		},
		{
			Message: "Re-adding Reflection2.constructor Removed in 671749d but used downstream in jclouds-labs",
			DiffId:  7,
			ModelId: 1,
		},
		{
			Message: "add Reflection2 constructor",
			DiffId:  7,
			ModelId: 2,
		},
		{
			Message: "Feat: Add constructor method to Reflection2 class This commit adds a new method `constructor` to the `Reflection2` class. This method returns an `Invokable` object that reflects a constructor present in the `TypeToken` type.",
			DiffId:  7,
			ModelId: 3,
		},
		{
			Message: "Added intro image .",
			DiffId:  7,
			ModelId: 4,
		},
		{
			Message: "Fix - was using base URI for request URI.",
			DiffId:  8,
			ModelId: 1,
		},
		{
			Message: "fix typo in HttpOp",
			DiffId:  8,
			ModelId: 2,
		},
		{
			Message: "Refactor: Update URI determination in execHttpPost method In the execHttpPost method, the way the request URI is determined has been updated. The method determineBaseIRI(url) has been replaced with determineRequestURI(url). This change does not alter the functionality of the method but improves its structure.",
			DiffId:  8,
			ModelId: 3,
		},
		{
			Message: "Move BikePark * to correct location .",
			DiffId:  8,
			ModelId: 4,
		},		
	}
	_ = database.Db.SaveCommitMessages(&seeds)
}
