package cmd

import (
	"fmt"
	"github.com/bazgab/opencpe/utils/logging"
	"github.com/spf13/cobra"
)

type Config struct {
	Authentication Authentication `json:"authentication"`
	Notification   Notification   `json:"notification"`
	IgnoredTags    IgnoredTags    `json:"ignored_tags"`
}

type Authentication struct {
	AwsProfile string `json:"aws_profile"`
}

type Notification struct {
	SmtpHost      string `json:"smtp_host"`
	SmtpPort      int    `json:"smtp_port"`
	EmailFrom     string `json:"email_from"`
	EmailPassword string `json:"email_password"`
}

type IgnoredTags struct {
	Owner   string `json:"owner"`
	Project string `json:"project"`
}

var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "only notify resource owners",
	Long:  `This command only notifies resource owners of policy infringement, as opposed to notify-and-delete.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Check if the global flag "policy" was actually set by the user
		if !cmd.Flags().Changed("policy") {
			return fmt.Errorf("required flag 'policy' not set for the 'notify' command")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Checking if global flags are working
		fmt.Println("Query Request Output:")
		logging.TextRequestOutputLogger("Printing values from global flags", flagConfig, flagPolicy, flagRegion)

	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)

}
