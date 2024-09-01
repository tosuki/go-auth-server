package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func checkVariables(variables []string) error {
	for _, variableName := range variables {
		variableValue := os.Getenv(variableName)

		if variableValue == "" {
			return fmt.Errorf("missing %s on .env file", variableName)
		}
	}

	return nil
}

func InitializeEnv(variables ...string) error {
	err := godotenv.Load("../../.env")

	if err != nil {
		return fmt.Errorf("failed too load env variables due to %s", err.Error())
	}

	missingVariable := checkVariables(variables)

	if missingVariable != nil {
		panic(missingVariable.Error())
	}

	return nil
}
