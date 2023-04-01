package cmd

import (
	"github.com/AlecAivazis/survey/v2"
)

func GetStorage() (string, error) {
	storage := ""
	prompt := &survey.Select{
		Message: "Choose a data storage option:",
		Options: []string{"MySQL", "Postgres"},
	}

	err := survey.AskOne(prompt, &storage)
	if err != nil {
		return "", err
	}

	return storage, nil
}
