package cmd

import (
	"fmt"
	"github.com/bazgab/opencpe/config"
	"github.com/spf13/cobra"
	"os"
)

// Flag declaration
var flagPolicy string
var flagRegion string

var rootCmd = &cobra.Command{
	Use:   "opencpe",
	Short: "All-in-one Cloud-Native Policy Engine with sensible defaults",
	Long: `OpenCPE is a tool for managing cloud resources with heavily opinionated defaults.

Note: All Global flags are required
	
A comprehensive policy reference and usage instructions can be found at https://github.com/bazgab/opencpe
`,
	Run: func(cmd *cobra.Command, args []string) {
		// In case there are no flags being passed, return the help page
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			return
		}

		fmt.Println("Initializing: ")
		fmt.Printf("Policy: %s\n", flagPolicy)
		fmt.Printf("Region: %s\n", flagRegion)

		config.LoadConfig()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVar(&flagRegion, "region", "", "Region to be queried (default is $OPENCPE_DEFAULT_REGION)")
	rootCmd.PersistentFlags().StringVar(&flagPolicy, "policy", "", "Policy to be executed")

	err := rootCmd.MarkPersistentFlagRequired("region")
	if err != nil {
		panic(err)
	}

	err = rootCmd.MarkPersistentFlagRequired("policy")
	if err != nil {
		panic(err)
	}

}
