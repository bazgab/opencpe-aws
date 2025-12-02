package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/bazgab/opencpe/policies"
	"github.com/bazgab/opencpe/utils/logging"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

type Config struct {
	Authentication struct {
		AwsProfile string `json:"aws_profile"`
	} `json:"authentication"`

	Notification struct {
		SmtpHost      string `json:"smtp_host"`
		SmtpPort      int    `json:"smtp_port"`
		EmailFrom     string `json:"email_from"`
		EmailPassword string `json:"email_password"`
	} `json:"notification"`

	IgnoredTags struct {
		Owner   []string `json:"owner"`
		Project []string `json:"project"`
	} `json:"ignored_tags"`
}

var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "notify resource owners of policy infringement",
	Long:  `This command only notifies resource owners of policy infringement, as opposed to notify-and-delete.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Check if the global flag "policy" was actually set by the user
		RequiredFlags := []string{"config", "policy", "region"}
		for _, flag := range RequiredFlags {
			if !cmd.Flags().Changed(flag) {
				return fmt.Errorf("required flag %s not set for the 'notify' command", flag)
			}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Checking if global flags are working
		fmt.Println("OpenCPE - Notify")
		logging.TextRequestOutputLogger("Query Request Output", flagConfig, flagPolicy, flagRegion)

		cfgFile, err := os.ReadFile(flagConfig)
		if err != nil {
			slog.Error("Error when opening file: ", err)
		}

		var cfg Config

		err = json.Unmarshal(cfgFile, &cfg)
		if err != nil {
			slog.Error("Error during Unmarshal(): ", err)
		}
		logging.BreakerLine()
		fmt.Println("")
		fmt.Println("Loaded Configuration:")
		fmt.Printf("-- Authentication.Profile: %s\n", cfg.Authentication.AwsProfile)
		fmt.Printf("-- Notification.SMTP_Host: %s\n", cfg.Notification.SmtpHost)
		fmt.Printf("-- Notification.SMTP_Port: %d\n", cfg.Notification.SmtpPort)
		fmt.Printf("-- Notification.From_Email: %s\n", cfg.Notification.EmailFrom)
		for _, owner := range cfg.IgnoredTags.Owner {
			fmt.Printf("-- IgnoredTags.Owner: %s\n", owner)
		}
		for _, project := range cfg.IgnoredTags.Project {
			fmt.Printf("-- IgnoredTags.Project: %s\n", project)
		}
		logging.BreakerLine()
		fmt.Println()

		//Check for policy
		if flagPolicy == "instance-age-2-days" {
			policies.InstanceAge2Days(cfg.Authentication.AwsProfile, flagRegion)
		}

	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)

}
