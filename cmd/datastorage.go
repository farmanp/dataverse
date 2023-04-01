import {
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/AlecAivazis/survey/v2"
}

func getStorage() (string, error) {
	storage := ""
	prompt := &survey.Select{
		Message: "Select a storage type:",
		Options: []string{"S3", "GCS", "Azure Blob Storage"}
	}

	err := survey.AskOne(prompt, &storage)
	if err != nil {
		return "", err
	}

	return storage, nil
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Select storage: %s\n", storage)
}