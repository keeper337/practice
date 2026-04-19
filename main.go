package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// WordList contains a list of words for the Hangman game
var WordList = []string{
	"apple", "banana", "cherry", "date", "elderberry",
	"fig", "grape", "honeydew", "kiwi", "lemon",
	"mango", "nectarine", "orange", "papaya", "quince",
	"raspberry", "strawberry", "tangerine", "watermelon", "blueberry",
}

// Game represents a Hangman game instance
type Game struct {
	word      string
	guessed   []bool
	attempts  int
	maxAttempts int
}

// NewGame creates a new Hangman game with a random word
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	word := WordList[rand.Intn(len(WordList))]
	return &Game{
		word:      word,
		guessed:   make([]bool, len(word)),
		attempts:  0,
		maxAttempts: 6,
	}
}

// Guess checks if a letter is in the word and updates the game state
func (g *Game) Guess(letter string) bool {
	if len(letter) != 1 || !strings.IsLower(letter) {
		return false
	}

	letter = strings.ToLower(letter)
	found := false

	for i, char := range g.word {
		if string(char) == letter {
			g.guessed[i] = true
			found = true
		}
	}

	if !found {
		g.attempts++
	}
	return found
}

// IsWon checks if the player has guessed all letters
func (g *Game) IsWon() bool {
	for _, guessed := range g.guessed {
		if !guessed {
			return false
		}
	}
	return true
}

// IsLost checks if the player has exceeded maximum attempts
func (g *Game) IsLost() bool {
	return g.attempts >= g.maxAttempts
}

// DisplayWord shows the current state of the word with guessed letters
func (g *Game) DisplayWord() string {
	var result []string
	for i, char := range g.word {
		if g.guessed[i] {
			result = append(result, string(char))
		} else {
			result = append(result, "_")
		}
	}
	return strings.Join(result, " ")
}

// DisplayHangman shows the hangman drawing based on attempts
func (g *Game) DisplayHangman() {
	stages := []string{
		`  +---+
  |   |
      |
      |
      |
      |
=========`,
		`  +---+
  |   |
  O   |
      |
      |
      |
=========`,
		`  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
		`  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
		`  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
		`  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
		`  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
	}

	if g.attempts < len(stages) {
		fmt.Println(stages[g.attempts])
	}
}

// DisplayAttempts shows the number of remaining attempts
func (g *Game) DisplayAttempts() {
	fmt.Printf("Attempts left: %d/%d\n", g.maxAttempts-g.attempts, g.maxAttempts)
}

func main() {
	game := NewGame()
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Try to guess the word by suggesting letters.")

	for !game.IsWon() && !game.IsLost() {
		fmt.Println("\nWord:", game.DisplayWord())
		game.DisplayAttempts()
		game.DisplayHangman()
		fmt.Print("Enter a letter: ")
		var input string
		fmt.Scanln(&input)
		if len(input) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}
		letter := strings.ToLower(input)
		if !game.Guess(letter) {
			fmt.Println("Incorrect guess!")
		} else {
			fmt.Println("Correct guess!")
		}
	}

	fmt.Println("\nWord:", game.DisplayWord())
	game.DisplayHangman()
	if game.IsWon() {
		fmt.Println("Congratulations! You won!")
	} else {
		fmt.Println("Game over! You lost.")
		fmt.Printf("The word was: %s\n", game.word)
	}
}