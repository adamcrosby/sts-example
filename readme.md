# STS Example #
This go code shows how to take an optional ARN and ExternalID contraint and assume a role via STS.

```golang
creds = stscreds.NewCredentials(sess, arn, func(p *stscreds.AssumeRoleProvider) {
  p.ExternalID = &externalID
})
```

The key here is to use the `stscreds` provider for aws.Config, rather than work with the `sts` package and service directly.
