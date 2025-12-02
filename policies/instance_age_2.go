package policies

import (
	"fmt"
	"github.com/bazgab/opencpe/config"
)

func InstanceAge2Days(profile string, region string) {
	config.LoadConfig(profile, region)
	fmt.Println("Policy: Instance-age-2-Days")
	fmt.Println("-- Description: Checks for EC2 instances that are over 2 days old")
	fmt.Printf("Profile: %s\n", profile)
	fmt.Printf("Region: %s\n", region)
}
