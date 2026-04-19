package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// WordList contains possible words for the hangman game
var WordList = []string{
	"computer", "programming", "algorithm", "function", "variable",
	"string", "integer", "boolean", "array", "slice", "map", "struct",
	"interface", "channel", "goroutine", "mutex", "panic", "recover",
	"defer", "method", "pointer", "package", "import", "export", "public",
	"private", "protected", "const", "var", "type", "return", "break",
	"continue", "for", "if", "else", "switch", "case", "default",
}

// HangmanGame represents a hangman game instance
type HangmanGame struct {
	word          string
	guessedWord   string
	guessedLetters map[rune]bool
	incorrectGuesses int
	maxIncorrectGuesses int
}

// NewHangmanGame creates a new hangman game with a random word
func NewHangmanGame() *HangmanGame {
	rand.Seed(time.Now().UnixNano())
	word := WordList[rand.Intn(len(WordList))]
	
	guessedWord := ""
	for _, char := range word {
		if char == ' ' {
			guessedWord += " "
		} else {
			guessedWord += "_"
		}
	}
	
	return &HangmanGame{
		word:                word,
		guessedWord:         guessedWord,
		guessedLetters:      make(map[rune]bool),
		incorrectGuesses:    0,
		maxIncorrectGuesses: 6,
	}
}

// GuessLetter processes a letter guess
func (h *HangmanGame) GuessLetter(letter rune) bool {
	letter = rune(strings.ToLower(string(letter))[0])
	
	// Check if letter was already guessed
	if h.guessedLetters[letter] {
		return false // Already guessed
	}
	
	h.guessedLetters[letter] = true
	
	// Check if the letter is in the word
	wordLower := strings.ToLower(h.word)
	found := strings.Contains(wordLower, string(letter))
	
	if !found {
		h.incorrectGuesses++
		return false
	}
	
	// Update guessed word with correct letters
	newGuessedWord := ""
	for i, char := range h.word {
		if char == ' ' {
			newGuessedWord += " "
		} else if strings.ToLower(string(char)) == string(letter) {
			newGuessedWord += string(char)
		} else {
			newGuessedWord += string(h.guessedWord[i])
		}
	}
	h.guessedWord = newGuessedWord
	
	return true
}

// IsGameOver checks if the game is over (win or lose)
func (h *HangmanGame) IsGameOver() bool {
	return h.IsWin() || h.IsLose()
}

// IsWin checks if the player has won
func (h *HangmanGame) IsWin() bool {
	return !strings.Contains(h.guessedWord, "_")
}

// IsLose checks if the player has lost
func (h *HangmanGame) IsLose() bool {
	return h.incorrectGuesses >= h.maxIncorrectGuesses
}

// GetGuessedWord returns the current state of the guessed word
func (h *HangmanGame) GetGuessedWord() string {
	return h.guessedWord
}

// GetIncorrectGuesses returns the number of incorrect guesses
func (h *HangmanGame) GetIncorrectGuesses() int {
	return h.incorrectGuesses
}

// GetMaxIncorrectGuesses returns the maximum allowed incorrect guesses
func (h *HangmanGame) GetMaxIncorrectGuesses() int {
	return h.maxIncorrectGuesses
}

// GetGuessedLetters returns all letters that have been guessed
func (h *HangmanGame) GetGuessedLetters() []rune {
	letters := make([]rune, 0, len(h.guessedLetters))
	for letter := range h.guessedLetters {
		letters = append(letters, letter)
	}
	return letters
}

// DisplayGame displays the current game state
func (h *HangmanGame) DisplayGame() {
	fmt.Printf("Word: %s\n", h.guessedWord)
	fmt.Printf("Incorrect guesses: %d/%d\n", h.incorrectGuesses, h.maxIncorrectGuesses)
	fmt.Printf("Guessed letters: %v\n", h.GetGuessedLetters())
}

func main() {
	game := NewHangmanGame()
	
	fmt.Println("Welcome to Hangman!")
	
	for !game.IsGameOver() {
		game.DisplayGame()
		
		var input string
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&input)
		
		if len(input) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}
		
		letter := rune(input[0])
		if !game.GuessLetter(letter) {
			fmt.Printf("Letter '%c' was already guessed or is incorrect.\n", letter)
		} else {
			fmt.Printf("Good guess! Letter '%c' is in the word.\n", letter)
		}
		
		fmt.Println()
	}
	
	game.DisplayGame()
	
	if game.IsWin() {
		fmt.Println("Congratulations! You won!")
	} else {
		fmt.Println("Game over! You lost.")
		fmt.Printf("The word was: %s\n", game.word)
	}
}