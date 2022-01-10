func setCors() {
	//    go run s3_put_bucket_cors.go -b BUCKET_NAME get put

	//// snippet-start:[s3.go.set_cors.vars]
	//bucketPtr := flag.String("b", "mahdicpp", "Bucket to set CORS on, (required)")
	//
	//flag.Parse()
	//
	//if *bucketPtr == "" {
	//	exitErrorf("-b <bucket> Bucket name required")
	//}
	var bucket string = "mahdicpp"
	methods := filterMethods(flag.Args())
	// snippet-end:[s3.go.set_cors.vars]

	// Initialize a session
	// snippet-start:[s3.go.set_cors.session]
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("02b5f145-6b3e-468c-94d1-3e02bb834f0d", "630a75917ce5d93e5f7cb82b6b3e7d8d3e0d64645d49728ae46f308ec36d90b2", ""),
	})
	svc := s3.New(sess, &aws.Config{
		Region:   aws.String("default"),
		Endpoint: aws.String("https://s3.ir-thr-at1.arvanstorage.com"),
	})
	// snippet-end:[s3.go.set_cors.session]

	// Create a CORS rule for the bucket
	// snippet-start:[s3.go.set_cors.rule]
	rule := s3.CORSRule{
		AllowedHeaders: aws.StringSlice([]string{"Authorization"}),
		AllowedOrigins: aws.StringSlice([]string{"*"}),
		MaxAgeSeconds:  aws.Int64(3000),

		// Add HTTP methods CORS request that were specified in the CLI.
		AllowedMethods: aws.StringSlice(methods),
	}
	// snippet-end:[s3.go.set_cors.rule]

	// Create the parameters for the PutBucketCors API call, add add
	// the rule created to it.
	// snippet-start:[s3.go.set_cors.put]
	params := s3.PutBucketCorsInput{
		Bucket: aws.String(bucket),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: []*s3.CORSRule{&rule},
		},
	}

	_, err = svc.PutBucketCors(&params)
	if err != nil {
		// Print the error message
		exitErrorf("Unable to set Bucket %q's CORS, %v", bucket, err)
	}

	// Print the updated CORS config for the bucket
	fmt.Printf("Updated bucket %q CORS for %v\n", bucket, methods)
	// snippet-end:[s3.go.set_cors.put]
}
