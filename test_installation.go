package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if INSTALLATION.md exists
	if _, err := os.Stat("INSTALLATION.md"); os.IsNotExist(err) {
		fmt.Println("ERROR: INSTALLATION.md file is missing")
		os.Exit(1)
	}

	// Read the content of INSTALLATION.md
	content, err := os.ReadFile("INSTALLATION.md")
	if err != nil {
		fmt.Printf("ERROR: Could not read INSTALLATION.md: %v\n", err)
		os.Exit(1)
	}

	installationContent := string(content)

	// Validate that the file contains expected sections
	expectedSections := []string{
		"Prerequisites",
		"Installation Steps",
		"Cross-Platform Compatibility",
		"Troubleshooting",
	}

	for _, section := range expectedSections {
		if !strings.Contains(installationContent, section) {
			fmt.Printf("ERROR: INSTALLATION.md is missing expected section: %s\n", section)
			os.Exit(1)
		}
	}

	fmt.Println("SUCCESS: INSTALLATION.md file exists and contains all required sections")
}