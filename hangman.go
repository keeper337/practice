package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

// WordList contains the list of words for the Hangman game
var WordList = []string{
	"programming",
	"go",
	"computer",
	"algorithm",
	"function",
	"variable",
	"string",
	"integer",
	"boolean",
	"array",
	"slice",
	"map",
	"struct",
	"interface",
	"channel",
	"goroutine",
	"mutex",
	"panic",
	"recover",
	"defer",
}

// HangmanGame represents the state of a Hangman game
type HangmanGame struct {
	Word          string
	GuessedLetters []rune
	IncorrectGuesses int
	MaxIncorrectGuesses int
}

// NewHangmanGame creates a new Hangman game with a random word
func NewHangmanGame() *HangmanGame {
	rand.Seed(time.Now().UnixNano())
	word := WordList[rand.Intn(len(WordList))]

	return &HangmanGame{
		Word:          strings.ToLower(word),
		GuessedLetters: []rune{},
		IncorrectGuesses: 0,
		MaxIncorrectGuesses: 6,
	}
}

// GuessLetter processes a letter guess
func (h *HangmanGame) GuessLetter(letter rune) bool {
	if h.IsGameOver() {
		return false
	}

	// Check if the letter was already guessed
	for _, l := range h.GuessedLetters {
		if l == letter {
			return false // Letter already guessed
		}
	}

	h.GuessedLetters = append(h.GuessedLetters, letter)

	// Check if the guess is incorrect
	if !strings.ContainsRune(h.Word, letter) {
		h.IncorrectGuesses++
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
	if h.IsGameOver() {
		return h.IncorrectGuesses < h.MaxIncorrectGuesses
	}

	// Check if all letters in the word have been guessed
	for _, letter := range h.Word {
		found := false
		for _, guessed := range h.GuessedLetters {
			if guessed == letter {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// IsLoser checks if the player has lost
func (h *HangmanGame) IsLoser() bool {
	return h.IncorrectGuesses >= h.MaxIncorrectGuesses
}

// GetWordDisplay returns the word with guessed letters revealed and others hidden
func (h *HangmanGame) GetWordDisplay() string {
	var display strings.Builder
	for _, letter := range h.Word {
		found := false
		for _, guessed := range h.GuessedLetters {
			if guessed == letter {
				display.WriteRune(letter)
				found = true
				break
			}
		}
		if !found {
			display.WriteRune('_')
		}
		display.WriteRune(' ')
	}
	return display.String()
}

// GetIncorrectGuesses returns the list of incorrect guesses
func (h *HangmanGame) GetIncorrectGuesses() string {
	var incorrect strings.Builder
	for _, letter := range h.GuessedLetters {
		if !strings.ContainsRune(h.Word, letter) {
			incorrect.WriteRune(letter)
			incorrect.WriteRune(' ')
		}
	}
	return incorrect.String()
}

func main() {
	game := NewHangmanGame()

	fmt.Println("Welcome to Hangman!")
	fmt.Println("Guess the word by entering one letter at a time.")
	fmt.Println("You have 6 incorrect guesses allowed.")

	for !game.IsGameOver() {
		fmt.Printf("\nWord: %s\n", game.GetWordDisplay())
		fmt.Printf("Incorrect guesses: %s\n", game.GetIncorrectGuesses())
		fmt.Printf("Remaining guesses: %d\n", game.MaxIncorrectGuesses-game.IncorrectGuesses)

		var input string
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&input)

		if len(input) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}

		letter := rune(strings.ToLower(input)[0])
		
		// Validate that the input is a letter
		if !unicode.IsLetter(letter) {
			fmt.Println("Please enter a valid letter.")
			continue
		}

		if !game.GuessLetter(letter) {
			fmt.Printf("Incorrect guess! The letter '%c' is not in the word.\n", letter)
		} else {
			fmt.Printf("Correct guess! The letter '%c' is in the word.\n", letter)
		}
	}

	if game.IsWinner() {
		fmt.Printf("\nCongratulations! You've guessed the word: %s\n", game.Word)
	} else {
		fmt.Printf("\nGame over! You've run out of guesses. The word was: %s\n", game.Word)
	}
}