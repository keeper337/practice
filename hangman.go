package main

import (
	"fmt"
	"strings"
	"unicode"
)

// HangmanGame represents a hangman game instance
type HangmanGame struct {
	wordToGuess string
	guessedLetters map[rune]bool
	maxWrongGuesses int
	wrongGuesses int
}

// NewHangmanGame creates a new hangman game with the given word
func NewHangmanGame(word string) *HangmanGame {
	return &HangmanGame{
		wordToGuess: strings.ToLower(word),
		guessedLetters: make(map[rune]bool),
		maxWrongGuesses: 6,
		wrongGuesses: 0,
	}
}

// DisplayWord shows the current state of the word with guessed letters
func (h *HangmanGame) DisplayWord() string {
	var display strings.Builder
	for _, char := range h.wordToGuess {
		if h.guessedLetters[char] {
			display.WriteRune(char)
			display.WriteString(" ")
		} else {
			display.WriteString("_ ")
		}
	}
	return strings.TrimSpace(display.String())
}

// GuessLetter processes a letter guess
func (h *HangmanGame) GuessLetter(letter rune) bool {
	// Convert to lowercase for case-insensitive comparison
	letter = unicode.ToLower(letter)
	
	if h.guessedLetters[letter] {
		return true // Already guessed
	}
	
	h.guessedLetters[letter] = true
	
	if !strings.ContainsRune(h.wordToGuess, letter) {
		h.wrongGuesses++
		return false
	}
	
	return true
}

// IsGameOver checks if the game is over (win or lose)
func (h *HangmanGame) IsGameOver() bool {
	return h.IsWinner() || h.IsLoser()
}

// IsWinner checks if the player has won
func (h *HangmanGame) IsWinner() bool {
	for _, char := range h.wordToGuess {
		if !h.guessedLetters[char] {
			return false
		}
	}
	return true
}

// IsLoser checks if the player has lost
func (h *HangmanGame) IsLoser() bool {
	return h.wrongGuesses >= h.maxWrongGuesses
}

// GetWrongGuesses returns the number of wrong guesses
func (h *HangmanGame) GetWrongGuesses() int {
	return h.wrongGuesses
}

// GetMaxWrongGuesses returns the maximum allowed wrong guesses
func (h *HangmanGame) GetMaxWrongGuesses() int {
	return h.maxWrongGuesses
}

// SetWordToGuess sets a new word to guess
func (h *HangmanGame) SetWordToGuess(word string) {
	h.wordToGuess = strings.ToLower(word)
	h.guessedLetters = make(map[rune]bool)
	h.wrongGuesses = 0
}

// DisplayHangman shows a visual representation of the hangman progress
func (h *HangmanGame) DisplayHangman() {
	fmt.Println("  +---+")
	fmt.Println("  |   |")
	
	switch h.wrongGuesses {
	case 0:
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 1:
		fmt.Println("  O   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 2:
		fmt.Println("  O   |")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 3:
		fmt.Println("  O   |")
		fmt.Println(" /|   |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 4:
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 5:
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" /    |")
		fmt.Println("=========")
	case 6:
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" / \\  |")
		fmt.Println("=========")
	}
}

func main() {
	// Example usage
	game := NewHangmanGame("hangman")
	
	fmt.Println("Welcome to Hangman!")
	fmt.Printf("Word: %s\n", game.DisplayWord())
	game.DisplayHangman()
	
	// Simulate some guesses
	game.GuessLetter('h')
	game.GuessLetter('a')
	game.GuessLetter('n')
	game.GuessLetter('g')
	
	fmt.Printf("Word: %s\n", game.DisplayWord())
	game.DisplayHangman()
	fmt.Printf("Wrong guesses: %d/%d\n", game.GetWrongGuesses(), game.GetMaxWrongGuesses())
}