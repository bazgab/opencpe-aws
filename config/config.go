package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"log"
)

func LoadConfig(profile string, region string) {
	// Expected JSON value from configuration file
	_, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}
}
