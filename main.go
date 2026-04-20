package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	words := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew"}
	word := words[rand.Intn(len(words))]
	guessed := make(map[rune]bool)
	incorrectGuesses := 0
	maxIncorrectGuesses := 6

	fmt.Println("Welcome to Hangman!")
	for incorrectGuesses < maxIncorrectGuesses {
		fmt.Print("\nWord: ")
		for _, char := range word {
			if guessed[char] {
				fmt.Printf("%c ", char)
			} else {
				fmt.Print("_ ")
			}
		}

		// Display hangman ASCII art based on incorrect guesses
		displayHangman(incorrectGuesses)

		var input string
		fmt.Print("\nEnter a letter: ")
		fmt.Scanln(&input)

		if len(input) != 1 || !strings.IsLetter(input[0]) {
			fmt.Println("Please enter a single letter.")
			continue
		}

		letter := rune(strings.ToLower(input)[0])
		if guessed[letter] {
			fmt.Println("You already guessed that letter.")
			continue
		}

		guessed[letter] = true

		if strings.ContainsRune(word, letter) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
			incorrectGuesses++
		}

		if isWordGuessed(word, guessed) {
			fmt.Println("\nCongratulations! You've won!")
			return
		}
	}

	fmt.Println("\nGame over! You've lost.")
	fmt.Printf("The word was: %s\n", word)
	displayHangman(incorrectGuesses)
}

func isWordGuessed(word string, guessed map[rune]bool) bool {
	for _, char := range word {
		if !guessed[char] {
			return false
		}
	}
	return true
}

func displayHangman(incorrectGuesses int) {
	hangmanParts := []string{
		"  +---+",
		"  |   |",
		"      |",
		"      |",
		"      |",
		"      |",
		"=========",
	}

	switch incorrectGuesses {
	case 1:
		hangmanParts[2] = "  O   |"
	case 2:
		hangmanParts[2] = "  O   |"
		hangmanParts[3] = "  |   |"
	case 3:
		hangmanParts[2] = "  O   |"
		hangmanParts[3] = " /|   |"
	case 4:
		hangmanParts[2] = "  O   |"
		hangmanParts[3] = " /|\\  |"
	case 5:
		hangmanParts[2] = "  O   |"
		hangmanParts[3] = " /|\\  |"
		hangmanParts[4] = " /    |"
	case 6:
		hangmanParts[2] = "  O   |"
		hangmanParts[3] = " /|\\  |"
		hangmanParts[4] = " / \\  |"
	}

	for _, part := range hangmanParts {
		fmt.Println(part)
	}
	fmt.Println()
}