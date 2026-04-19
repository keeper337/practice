package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Word lists for different difficulty levels
var (
	easyWords   = []string{"cat", "dog", "sun", "car", "book", "tree", "fish", "bird", "moon", "star"}
	mediumWords = []string{"computer", "elephant", "mountain", "butterfly", "ocean", "library", "guitar", "chocolate", "adventure", "rainbow"}
	hardWords   = []string{"phenomenon", "extravagant", "serendipity", "quintessential", "mellifluous", "ephemeral", "ubiquitous", "luminous", "resilient", "voracious"}
)

func main() {
	fmt.Println("Welcome to the Word Game!")
	fmt.Println("==========================")
	fmt.Println("Please select a difficulty level:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	fmt.Print("\nEnter your choice (1-3): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	choice := strings.TrimSpace(input)
	
	switch choice {
	case "1":
		fmt.Println("\nYou selected Easy difficulty.")
		loadWordList(easyWords)
	case "2":
		fmt.Println("\nYou selected Medium difficulty.")
		loadWordList(mediumWords)
	case "3":
		fmt.Println("\nYou selected Hard difficulty.")
		loadWordList(hardWords)
	default:
		fmt.Println("\nInvalid choice. Please run the program again and select 1, 2, or 3.")
		return
	}
}

func loadWordList(words []string) {
	fmt.Println("Loading word list...")
	fmt.Println("Words for this difficulty level:")
	for i, word := range words {
		fmt.Printf("%d. %s\n", i+1, word)
	}
	fmt.Println("\nGame would start now with these words.")
}