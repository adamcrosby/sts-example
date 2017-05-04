package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	var arn string
	var externalID string

	const (
		defaultARN   = ""
		arnUsage     = "The ARN of the role you need to assume"
		defaultExtID = ""
		extIDUsage   = "The ExternalID constraint, if applicable for the role you need to assume"
		region       = "us-east-1"
	)
	flag.StringVar(&arn, "arn", defaultARN, arnUsage)
	flag.StringVar(&externalID, "extid", defaultExtID, extIDUsage)
	flag.Parse()

	sess := session.Must(session.NewSession())
	conf := createConfig(arn, externalID, region, sess)

	fmt.Println("This should print the S3 buckets available in your account.  If you passed in an ARN, it will print the S3 buckets in the Assumed Role account.")
	s3Svc := s3.New(sess, &conf)
	var input *s3.ListBucketsInput
	resp, _ := s3Svc.ListBuckets(input)
	fmt.Println(resp)
}

func createConfig(arn string, externalID string, region string, sess *session.Session) aws.Config {

	conf := aws.Config{Region: aws.String(region)}
	if arn != "" {
		// if ARN flag is passed in, we need to be able ot assume role here
		var creds *credentials.Credentials
		if externalID != "" {
			// If externalID flag is passed, we need to include it in credentials struct
			creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {
				p.ExternalID = &externalID
			})
		} else {
			creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {})
		}
		conf.Credentials = creds
	}
	return conf
}
