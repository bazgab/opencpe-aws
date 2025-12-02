package policies

import (
	"fmt"
	"github.com/bazgab/opencpe/config"
)

func InstanceAge2Days(profile string, region string) {
	config.LoadConfig(profile, region)
	fmt.Println("This is from the policies package")
	fmt.Printf("Profile: %s\n", profile)
	fmt.Printf("Region: %s\n", region)
}
