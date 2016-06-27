package main

import (
	"fmt"
	"os"
)

func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {

	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}
func containsString(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

