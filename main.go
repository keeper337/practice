package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Game represents the state of the game
type Game struct {
	word        string
	guessed     []bool
	lives       int
	maxLives    int
	wordLength  int
}

// WordList represents different difficulty levels
type WordList struct {
	Easy   []string
	Medium []string
	Hard   []string
}

// InitializeGame initializes a new game with a random word from the specified difficulty
func InitializeGame(difficulty string) *Game {
	// Define word lists for different difficulties
	wordLists := WordList{
		Easy:   []string{"cat", "dog", "sun", "car", "book"},
		Medium: []string{"hello", "world", "computer", "program", "language"},
		Hard:   []string{"algorithm", "development", "programming", "application", "framework"},
	}

	// Select a random word based on difficulty
	var selectedWord string
	switch difficulty {
	case "easy":
		selectedWord = getRandomWord(wordLists.Easy)
	case "medium":
		selectedWord = getRandomWord(wordLists.Medium)
	case "hard":
		selectedWord = getRandomWord(wordLists.Hard)
	default:
		selectedWord = getRandomWord(wordLists.Easy) // Default to easy if invalid difficulty
	}

	// Initialize game state
	game := &Game{
		word:       strings.ToLower(selectedWord),
		guessed:    make([]bool, len(selectedWord)),
		lives:      6,
		maxLives:   6,
		wordLength: len(selectedWord),
	}

	return game
}

// getRandomWord returns a random word from the given list
func getRandomWord(wordList []string) string {
	rand.Seed(time.Now().UnixNano())
	return wordList[rand.Intn(len(wordList))]
}

// DisplayGameState shows the current state of the game
func (g *Game) DisplayGameState() {
	fmt.Print("Word: ")
	for i, char := range g.word {
		if g.guessed[i] {
			fmt.Printf("%c ", char)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Printf("\nLives left: %d\n", g.lives)
}

// MakeGuess processes a guess
func (g *Game) MakeGuess(guess string) bool {
	if len(guess) != 1 {
		fmt.Println("Please enter only one letter.")
		return false
	}

	guess = strings.ToLower(guess)
	guessedIndex := -1

	// Check if the guessed letter is in the word
	for i, char := range g.word {
		if string(char) == guess && !g.guessed[i] {
			g.guessed[i] = true
			guessedIndex = i
		}
	}

	// If the letter wasn't found, lose a life
	if guessedIndex == -1 {
		g.lives--
		fmt.Printf("Letter '%s' is not in the word.\n", guess)
		return false
	}

	fmt.Printf("Good guess! Letter '%s' is in the word.\n", guess)
	return true
}

// IsGameOver checks if the game is over (win or lose)
func (g *Game) IsGameOver() bool {
	// Check for win condition (all letters guessed)
	for _, guessed := range g.guessed {
		if !guessed {
			return false
		}
	}

	// If all letters are guessed, player wins
	if strings.Join(g.getGuessedWord(), "") == g.word {
		fmt.Println("Congratulations! You've won!")
		return true
	}

	// Check for lose condition (no lives left)
	if g.lives <= 0 {
		fmt.Println("Game over! You've run out of lives.")
		fmt.Printf("The word was: %s\n", g.word)
		return true
	}

	return false
}

// getGuessedWord returns the current state of the guessed word as a string
func (g *Game) getGuessedWord() []string {
	result := make([]string, len(g.word))
	for i, guessed := range g.guessed {
		if guessed {
			result[i] = string(g.word[i])
		} else {
			result[i] = "_"
		}
	}
	return result
}

func main() {
	fmt.Println("Welcome to the Word Guessing Game!")
	
	// Initialize game with a random word from easy difficulty
	game := InitializeGame("easy")
	
	// Display initial game state
	game.DisplayGameState()
	
	// Simple example of making guesses
	game.MakeGuess("c")
	game.DisplayGameState()
	
	game.MakeGuess("a")
	game.DisplayGameState()
	
	game.MakeGuess("t")
	game.DisplayGameState()
	
	// Check if game is over
	if game.IsGameOver() {
		fmt.Println("Game ended.")
	} else {
		fmt.Println("Keep playing!")
	}
}