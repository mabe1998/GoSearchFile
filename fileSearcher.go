package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// ReadFile reads text based files (.txt, .json, etc...) and returns its content.
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return "", err
	}
	return string(content), nil
}

func SearchInFile(path string, searchString string) (bool, error) {
	content, err := ReadFile(path)
	if err != nil {
		log.Default()
		return false, err
	}
	return strings.Contains(content, searchString), nil

}
