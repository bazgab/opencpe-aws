package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadConfig() {
	// Expected JSON value from configuration file
	_, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(""),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}
}
