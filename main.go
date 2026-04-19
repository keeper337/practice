package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Welcome to the Game!")
	fmt.Println("====================")
	
	// Display game state
	displayGameState()
	
	// Prompt for input with clear instructions
	fmt.Print("\nEnter your move (1-9) or 'quit' to exit: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	
	// Validate input
	if input == "quit" {
		fmt.Println("Thanks for playing!")
		return
	}
	
	move, err := strconv.Atoi(input)
	if err != nil || move < 1 || move > 9 {
		fmt.Println("Error: Invalid input. Please enter a number between 1 and 9.")
		return
	}
	
	// Process the move
	fmt.Printf("You entered move: %d\n", move)
	fmt.Println("Processing your move...")
	
	// Display updated game state
	displayGameState()
}

func displayGameState() {
	fmt.Println("\nCurrent Game State:")
	fmt.Println(" 1 | 2 | 3 ")
	fmt.Println("-----------")
	fmt.Println(" 4 | 5 | 6 ")
	fmt.Println("-----------")
	fmt.Println(" 7 | 8 | 9 ")
	fmt.Println("\nEach number represents a position on the board.")
}