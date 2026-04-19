package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the game!")
	
	// Game loop
	for {
		fmt.Print("Enter a letter guess (or 'quit' to exit): ")
		
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		
		input = strings.TrimSpace(input)
		
		// Check for quit command
		if input == "quit" {
			fmt.Println("Thanks for playing!")
			break
		}
		
		// Validate input - must be a single letter
		if len(input) != 1 || !isLetter(input[0]) {
			fmt.Println("Please enter a single letter.")
			continue
		}
		
		// Process the guess (placeholder logic)
		guess := strings.ToLower(input)
		fmt.Printf("You guessed: %s\n", guess)
		
		// Placeholder for game logic
		// In a real implementation, this would check against the word to guess
		// and handle win/lose conditions
		
		// For now, just continue the loop
	}
}

// isLetter checks if a character is a letter
func isLetter(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}