package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Word lists for different difficulty levels
var (
	easyWords   = []string{"cat", "dog", "sun", "car", "book"}
	mediumWords = []string{"hello", "world", "computer", "program", "golang"}
	hardWords   = []string{"algorithm", "development", "programming", "implementation", "abstraction"}
)

// GameState represents the current state of the game
type GameState struct {
	word        string
	guessedWord string
	lives       int
	letters     map[rune]bool
}

// NewGameState initializes a new game with a random word from the specified difficulty
func NewGameState(difficulty string) *GameState {
	rand.Seed(time.Now().UnixNano())
	
	var wordList []string
	switch difficulty {
	case "easy":
		wordList = easyWords
	case "medium":
		wordList = mediumWords
	case "hard":
		wordList = hardWords
	default:
		wordList = easyWords // default to easy if invalid difficulty
	}
	
	selectedWord := wordList[rand.Intn(len(wordList))]
	
	// Initialize guessed word with blanks
	guessedWord := strings.Repeat("_", len(selectedWord))
	
	return &GameState{
		word:        selectedWord,
		guessedWord: guessedWord,
		lives:       6, // typical number of lives
		letters:     make(map[rune]bool),
	}
}

// GuessLetter processes a letter guess
func (gs *GameState) GuessLetter(letter rune) bool {
	if _, exists := gs.letters[letter]; exists {
		return false // already guessed
	}
	
	gs.letters[letter] = true
	
	// Check if the letter is in the word
	found := strings.Contains(gs.word, string(letter))
	
	if !found {
		gs.lives--
		return false
	}
	
	// Update guessed word with the correct letter
	newGuessedWord := ""
	for _, char := range gs.word {
		if _, guessed := gs.letters[char]; guessed {
			newGuessedWord += string(char)
		} else {
			newGuessedWord += "_"
		}
	}
	gs.guessedWord = newGuessedWord
	
	return true
}

// IsGameOver checks if the game is over (either won or lost)
func (gs *GameState) IsGameOver() bool {
	return gs.lives <= 0 || !strings.Contains(gs.guessedWord, "_")
}

// IsGameWon checks if the player has won
func (gs *GameState) IsGameWon() bool {
	return !strings.Contains(gs.guessedWord, "_")
}

func main() {
	fmt.Println("Welcome to Word Blanks Game!")
	
	// Initialize game with medium difficulty
	game := NewGameState("medium")
	
	fmt.Printf("Word length: %d\n", len(game.word))
	fmt.Printf("Initial word: %s\n", game.guessedWord)
	fmt.Printf("Lives: %d\n", game.lives)
	
	// Example guesses
	exampleGuesses := []rune{'e', 'o', 'l', 'h', 'a', 'r', 't'}
	for _, guess := range exampleGuesses {
		fmt.Printf("\nGuessing letter '%c': ", guess)
		if game.GuessLetter(guess) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
		fmt.Printf("Current word: %s\n", game.guessedWord)
		fmt.Printf("Lives remaining: %d\n", game.lives)
		
		if game.IsGameOver() {
			if game.IsGameWon() {
				fmt.Println("Congratulations! You won!")
			} else {
				fmt.Println("Game over! You lost.")
			}
			break
		}
	}
}