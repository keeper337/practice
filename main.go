package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Word lists for different difficulties
var (
	easyWords   = []string{"cat", "dog", "sun", "car", "book", "tree", "fish", "bird", "moon", "star"}
	mediumWords = []string{"computer", "elephant", "butterfly", "mountain", "ocean", "guitar", "library", "garden", "castle", "river"}
	hardWords   = []string{"phenomenon", "extravagant", "serendipity", "quintessential", "ephemeral", "luminous", "mellifluous", "ubiquitous", "resplendent", "effervescent"}
)

func main() {
	fmt.Println("Welcome to the Word Game!")
	fmt.Println("==========================")
	
	// Get difficulty selection
	difficulty := getDifficultySelection()
	
	// Load appropriate word list based on difficulty
	var wordList []string
	switch difficulty {
	case "easy":
		wordList = easyWords
	case "medium":
		wordList = mediumWords
	case "hard":
		wordList = hardWords
	default:
		fmt.Println("Invalid difficulty selected. Using easy words by default.")
		wordList = easyWords
	}
	
	fmt.Printf("You selected %s difficulty.\n", difficulty)
	fmt.Println("Your word list:")
	for i, word := range wordList {
		fmt.Printf("%d. %s\n", i+1, word)
	}
}

func getDifficultySelection() string {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Println("\nPlease select a difficulty level:")
		fmt.Println("1. Easy")
		fmt.Println("2. Medium")
		fmt.Println("3. Hard")
		fmt.Print("Enter your choice (1-3): ")
		
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		
		switch input {
		case "1", "easy", "Easy":
			return "easy"
		case "2", "medium", "Medium":
			return "medium"
		case "3", "hard", "Hard":
			return "hard"
		default:
			fmt.Println("Invalid input. Please enter 1, 2, or 3, or the word 'easy', 'medium', or 'hard'.")
		}
	}
}