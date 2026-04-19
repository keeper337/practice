package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// WordList contains a list of words for the Hangman game
var WordList = []string{
	"hangman", "computer", "programming", "golang", "development",
	"algorithm", "function", "variable", "string", "integer",
}

// Game represents a Hangman game instance
type Game struct {
	word        string
	guessed     string
	wrongGuesses int
	maxWrong    int
}

// NewGame creates a new Hangman game with a random word
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	word := WordList[rand.Intn(len(WordList))]
	return &Game{
		word:        word,
		guessed:     strings.Repeat("_", len(word)),
		wrongGuesses: 0,
		maxWrong:    6,
	}
}

// Guess handles a letter guess
func (g *Game) Guess(letter string) bool {
	if strings.Contains(g.guessed, letter) || g.wrongGuesses >= g.maxWrong {
		return false
	}

	correct := false
	newGuessed := make([]byte, len(g.word))

	for i, char := range g.word {
		if string(char) == letter {
			newGuessed[i] = letter[0]
			correct = true
		} else {
			newGuessed[i] = g.guessed[i]
		}
	}

	g.guessed = string(newGuessed)

	if !correct {
		g.wrongGuesses++
	}

	return correct
}

// IsWon checks if the game is won
func (g *Game) IsWon() bool {
	return !strings.Contains(g.guessed, "_")
}

// IsLost checks if the game is lost
func (g *Game) IsLost() bool {
	return g.wrongGuesses >= g.maxWrong
}

// Display shows the current state of the game
func (g *Game) Display() {
	fmt.Println("Word:", g.guessed)
	fmt.Printf("Wrong guesses: %d/%d\n", g.wrongGuesses, g.maxWrong)
}

// HangmanDraw prints a simple hangman drawing based on wrong guesses
func (g *Game) HangmanDraw() {
	switch g.wrongGuesses {
	case 0:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 1:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 2:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 3:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 4:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 5:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" /    |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 6:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" / \\  |")
		fmt.Println("      |")
		fmt.Println("=========")
	}
}

func main() {
	game := NewGame()
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Try to guess the word by suggesting letters.")

	for !game.IsWon() && !game.IsLost() {
		game.Display()
		game.HangmanDraw()
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

	game.Display()
	game.HangmanDraw()
	if game.IsWon() {
		fmt.Println("Congratulations! You won!")
	} else {
		fmt.Println("Game over! You lost.")
		fmt.Printf("The word was: %s\n", game.word)
	}
}