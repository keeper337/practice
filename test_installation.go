package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if INSTALLATION.md exists
	if _, err := os.Stat("INSTALLATION.md"); os.IsNotExist(err) {
		fmt.Println("FAIL: INSTALLATION.md file does not exist")
		os.Exit(1)
	}

	// Read the content of INSTALLATION.md
	content, err := os.ReadFile("INSTALLATION.md")
	if err != nil {
		fmt.Printf("FAIL: Error reading INSTALLATION.md: %v\n", err)
		os.Exit(1)
	}

	// Convert to string for easier checking
	text := string(content)

	// Check for key sections
	requiredSections := []string{
		"Installation Instructions",
		"Prerequisites",
		"Installation Steps",
		"Cross-Platform Compatibility",
		"Troubleshooting",
	}

	for _, section := range requiredSections {
		if !strings.Contains(text, section) {
			fmt.Printf("FAIL: Required section '%s' not found in INSTALLATION.md\n", section)
			os.Exit(1)
		}
	}

	// Check for specific content
	requiredContent := []string{
		"Go (version 1.19 or higher)",
		"git clone",
		"go build -o hangman main.go",
		"./hangman",
		"go version",
		"go vet",
	}

	for _, content := range requiredContent {
		if !strings.Contains(text, content) {
			fmt.Printf("FAIL: Required content '%s' not found in INSTALLATION.md\n", content)
			os.Exit(1)
		}
	}

	fmt.Println("PASS: INSTALLATION.md file exists and contains all required sections and content")
}