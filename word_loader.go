package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LoadWords loads word list based on difficulty level
// This function is already correctly implemented
func LoadWords(difficulty string) ([]string, error) {
	var filename string
	
	switch strings.ToLower(difficulty) {
	case "easy":
		filename = "easy_words.txt"
	case "medium":
		filename = "medium_words.txt"
	case "hard":
		filename = "hard_words.txt"
	default:
		return nil, fmt.Errorf("unsupported difficulty level: %s", difficulty)
	}
	
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", filename, err)
	}
	defer file.Close()
	
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}
	
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}
	
	return words, nil
}

func main() {
	// Example usage
	fmt.Println("Loading easy words...")
	easyWords, err := LoadWords("easy")
	if err != nil {
		fmt.Printf("Error loading easy words: %v\n", err)
		return
	}
	fmt.Printf("Loaded %d easy words\n", len(easyWords))
	
	fmt.Println("Loading medium words...")
	mediumWords, err := LoadWords("medium")
	if err != nil {
		fmt.Printf("Error loading medium words: %v\n", err)
		return
	}
	fmt.Printf("Loaded %d medium words\n", len(mediumWords))
	
	fmt.Println("Loading hard words...")
	hardWords, err := LoadWords("hard")
	if err != nil {
		fmt.Printf("Error loading hard words: %v\n", err)
		return
	}
	fmt.Printf("Loaded %d hard words\n", len(hardWords))
}