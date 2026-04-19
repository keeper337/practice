package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Word lists for different difficulty levels
var easyWords = []string{"cat", "dog", "sun", "car", "book"}
var mediumWords = []string{"hello", "world", "computer", "program", "golang"}
var hardWords = []string{"algorithm", "development", "programming", "implementation", "functionality"}

// Game state structure
type GameState struct {
	word        string
	blanks      []rune
	lives       int
	wordLength  int
}

// Select a random word from the specified difficulty list
func selectRandomWord(difficulty string) string {
	var words []string
	switch difficulty {
	case "easy":
		words = easyWords
	case "medium":
		words = mediumWords
	case "hard":
		words = hardWords
	default:
		words = easyWords // default to easy if invalid difficulty
	}

	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

// Initialize game state with blanks and lives
func initializeGame(difficulty string) GameState {
	word := selectRandomWord(difficulty)
	blanks := make([]rune, len(word))
	for i := range blanks {
		blanks[i] = '_'
	}

	return GameState{
		word:       word,
		blanks:     blanks,
		lives:      6, // default lives
		wordLength: len(word),
	}
}

// Display the current state of the game
func displayGame(gameState GameState) {
	fmt.Printf("Word: %s\n", string(gameState.blanks))
	fmt.Printf("Lives: %d\n", gameState.lives)
	fmt.Printf("Word length: %d\n", gameState.wordLength)
}

func main() {
	// Initialize game with easy difficulty
	game := initializeGame("easy")
	
	// Display initial game state
	displayGame(game)
}