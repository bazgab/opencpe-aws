package errors

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/bazgab/opencpe-aws/utils/logging"
	"log"
	"os"
	"strconv"
)

func IdentityCheck(profile string, region string, account int) {
	targetProfile := profile
	targetRegion := region
	expectedAccount := account

	fmt.Println("Identity Check:")
	fmt.Println("----- ENVIRONMENT DIAGNOSTICS -----")
	if os.Getenv("AWS_PROFILE") == "" {
		fmt.Println("AWS_PROFILE env var not set.")
	} else {
		fmt.Printf("AWS_PROFILE env var: '%s'\n", os.Getenv("AWS_PROFILE"))
	}
	if os.Getenv("AWS_REGION") == "" {
		fmt.Println("AWS_REGION env var not set.")
	} else {
		fmt.Printf("AWS_REGION env var: '%s'\n", os.Getenv("AWS_REGION"))
	}
	fmt.Println("-----------------------------------")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(targetProfile),
		config.WithRegion(targetRegion),
	)
	if err != nil {
		log.Fatalf("Config Load Error: %v", err)
	}

	client := sts.NewFromConfig(cfg)

	identity, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalf("Authentication Failed: %v\n(Check if you ran 'aws sso login --profile desired-profile')", err)
	}

	currentAccount, err := strconv.Atoi(*identity.Account)
	if err != nil {
		log.Fatalf("Failed to convert account ID to int: %v", err)
	}

	fmt.Println("------ AUTHENTICATION RESULT ------")
	fmt.Printf("Connected Account ID: %d\n", currentAccount)
	fmt.Printf("Expected Account ID:  %d\n", expectedAccount)
	fmt.Printf("Connected ARN:        %s\n", *identity.Arn)
	fmt.Println("-----------------------------------")

	if currentAccount != expectedAccount {
		fmt.Println("❌ MISMATCH DETECTED!")
		fmt.Println("You are NOT connected to the account in your config file.")
		fmt.Println("Your terminal is likely pointing to a different set of credentials.")
	} else {
		fmt.Println("✅ Account ID matches.")
		fmt.Println("If this matches, but you see 0 instances, verify the INSTANCE ID in the console.")
	}
	logging.BreakerLine()
}
