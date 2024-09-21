package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for true {
		// TODO: implement Ui with tview (makes this an TLI i guess)
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What would you like to do?")
		fmt.Println("1: Search, 0: Exit")
		fmt.Printf("> ")
		input, err := CleanInput(reader)
		if input == "0" {
			os.Exit(0)
		}

		fmt.Println("Enter filenpath:")
		fmt.Printf("> ")
		path, err := CleanInput(reader) //path := "C:/Temp/Arbeitszeugnis.txt"
		Check(err)
		absPath := PreparePath(path)

		fmt.Println("Enter searchString:")
		fmt.Printf("> ")
		searchString, err := CleanInput(reader)
		Check(err)

		foundInFile, err := SearchInFile(absPath, searchString)
		Check(err)

		if foundInFile {
			fmt.Println("Its there!")
		} else {
			fmt.Println("Not there chief")
		}

		fmt.Println()
	}

}

// CleanInput reads command line input and trims the deliminator at the end
func CleanInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	cleanedInput := strings.TrimSuffix(input, "\n")
	return cleanedInput, err
}

// PreparePath cleans a path and makes it absolute.
func PreparePath(path string) string {
	cleanedPath := filepath.Clean(path)
	absPath, _ := filepath.Abs(cleanedPath)
	return absPath
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
