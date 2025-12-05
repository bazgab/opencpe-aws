package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// Flag declaration
var flagConfig string
var flagPolicy string
var flagRegion string

var rootCmd = &cobra.Command{
	Use:     "opencpe-aws",
	Short:   "All-in-one Cloud-Native Policy Engine with sensible defaults",
	Version: "0.0.1",
	Long: `OpenCPE is a tool for managing cloud resources with heavily opinionated defaults.
	
A comprehensive policy reference and usage instructions can be found at https://github.com/bazgab/opencpe-aws

Note: All Global Flags are required.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// In case there are no flags being passed, return the help page
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			return
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(notifyCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringVar(&flagConfig, "config", "", "Configuration file")
	rootCmd.PersistentFlags().StringVar(&flagPolicy, "policy", "", "Policy to be executed")
	rootCmd.PersistentFlags().StringVar(&flagRegion, "region", "", "Region to be queried")

}
