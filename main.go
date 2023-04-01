package main

import (
	"fmt"
	"os"

	"github.com/fpirzada/dataverse/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dataverse",
	Short: "Dataverse CLI is a command line interface for generating data engineering platforms",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Welcome to Dataverse CLI")
	},
}

func Execute() {
	storage, err := cmd.GetStorage()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
