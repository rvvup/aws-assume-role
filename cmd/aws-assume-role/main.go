package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"log"
)

const (
	defaultRoleSessionName = "aws-assume-role"
)

var (
	roleARNArg, externalIDArg, roleSessionNameArg string
	verboseArg, versionArg                        bool
	version                                       string
)

func init() {
	flag.StringVar(&roleARNArg, "role-arn", "", "The ARN of the role to assume")
	flag.StringVar(&externalIDArg, "external-id", "", "An external ID to use when assuming the role")
	flag.StringVar(&roleSessionNameArg, "role-session-name", defaultRoleSessionName, "An identifier for the assumed role session (A default is used if not supplied)")

	flag.BoolVar(&verboseArg, "verbose", false, "Enable verbose output")
	flag.BoolVar(&versionArg, "version", false, "Print version")

	flag.Parse()
}

func assumeRole() *aws.Credentials {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client := sts.NewFromConfig(cfg)
	id, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal(err)
	}

	if verboseArg {
		log.Printf("initial auth arn: %s\n", aws.ToString(id.Arn))
	}

	stsClient := sts.NewFromConfig(cfg)
	provider := stscreds.NewAssumeRoleProvider(stsClient, roleARNArg, func(o *stscreds.AssumeRoleOptions) {
		o.RoleARN = roleARNArg
		o.RoleSessionName = roleSessionNameArg
		if externalIDArg != "" {
			o.ExternalID = &externalIDArg
		}
	})

	cfg.Credentials = aws.NewCredentialsCache(provider)
	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &creds
}

func printVersion() {
	fmt.Println(version)
}

func printEnvCredentials(credentials aws.Credentials) {
	fmt.Printf("AWS_ACCESS_KEY_ID=%s AWS_SECRET_ACCESS_KEY=%s AWS_SESSION_TOKEN=%s", credentials.AccessKeyID, credentials.SecretAccessKey, credentials.SessionToken)
}

func main() {
	if versionArg {
		printVersion()
		return
	}
	creds := assumeRole()
	printEnvCredentials(*creds)
}
