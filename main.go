package main

import (
	"fmt"
)

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