package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Game represents the hangman game state
type Game struct {
	word        string
	guessed     []bool
	lives       int
	maxLives    int
	guessedLetters []rune
}

// NewGame creates a new hangman game
func NewGame(word string) *Game {
	return &Game{
		word:        strings.ToUpper(word),
		guessed:     make([]bool, len(word)),
		lives:       6,
		maxLives:    6,
		guessedLetters: []rune{},
	}
}

// Guess processes a letter guess
func (g *Game) Guess(letter rune) bool {
	// Convert letter to uppercase for comparison
	upperLetter := strings.ToUpper(string(letter))
	if len(upperLetter) == 0 {
		return false // Invalid input
	}
	letter = rune(upperLetter[0])
	
	// Check if letter was already guessed
	for _, l := range g.guessedLetters {
		if l == letter {
			return false // Already guessed
		}
	}
	
	g.guessedLetters = append(g.guessedLetters, letter)
	
	// Check if letter is in the word
	found := false
	for i, char := range g.word {
		if char == letter {
			g.guessed[i] = true
			found = true
		}
	}
	
	// Decrement lives if letter not found
	if !found {
		g.lives--
	}
	
	return found


// GetLives returns the current number of lives
func (g *Game) GetLives() int {
	return g.lives
}

// GetGuessed returns the current guessed state
func (g *Game) GetGuessed() []bool {
	return g.guessed
}

// IsWon checks if the game is won
func (g *Game) IsWon() bool {
	for _, guessed := range g.guessed {
		if !guessed {
			return false
		}
	}
	return true
}

// IsLost checks if the game is lost
func (g *Game) IsLost() bool {
	return g.lives <= 0
}

// DisplayWord returns the word with guessed letters revealed
func (g *Game) DisplayWord() string {
	var result strings.Builder
	for i, char := range g.word {
		if g.guessed[i] {
			result.WriteRune(char)
		} else {
			result.WriteRune('_')
		}
		result.WriteRune(' ')
	}
	return result.String()
}

// DisplayHangman returns a visual representation of hangman
func (g *Game) DisplayHangman() string {
	stages := []string{
		`  +---+
  |   |
      |
      |
      |
      |
		`  +---+
  |   |
  O   |
      |
      |
      |
		`  +---+
  |   |
  O   |
  |   |
      |
      |
		`  +---+
  |   |
  O   |
 /|   |
      |
      |
		`  +---+
  |   |
  O   |
 /|\  |
      |
      |
		`  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
		`  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
	}
	
	stageIndex := g.maxLives - g.lives
	if stageIndex < 0 {
		stageIndex = 0
	}
	if stageIndex >= len(stages) {
		stageIndex = len(stages) - 1
	}
	
	return stages[stageIndex]
}

// DisplayGuessedLetters returns the list of guessed letters
func (g *Game) DisplayGuessedLetters() string {
	return fmt.Sprintf("Guessed letters: %s", string(g.guessedLetters))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	words := []string{"GOLANG", "COMPUTER", "PROGRAM", "DEVELOPER", "FUNCTION", "VARIABLE"}
	word := words[rand.Intn(len(words))]
	
	game := NewGame(word)
	
	fmt.Println("Welcome to Hangman!")
	
	for !game.IsWon() && !game.IsLost() {
		fmt.Println("\n" + game.DisplayHangman())
		fmt.Println("Word: " + game.DisplayWord())
		fmt.Println(game.DisplayGuessedLetters())
		fmt.Printf("Lives remaining: %d\n", game.GetLives())
		
		var input string
		fmt.Print("Guess a letter: ")
		fmt.Scanln(&input)
		
		if len(input) > 0 {
			letter := rune(strings.ToUpper(input)[0])
			if game.Guess(letter) {
				fmt.Println("Good guess!")
			} else {
				fmt.Println("Wrong guess!")
			}
		}
	}
	
	fmt.Println("\n" + game.DisplayHangman())
	fmt.Println("Word: " + game.DisplayWord())
	
	if game.IsWon() {
		fmt.Println("Congratulations! You won!")
	} else {
		fmt.Println("Game over! You lost!")
		fmt.Printf("The word was: %s\n", game.word)
	}
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Word lists for different difficulties
var (
	easyWords   = []string{"cat", "dog", "sun", "car", "book", "tree", "fish", "bird", "moon", "star"}
	mediumWords = []string{"computer", "elephant", "butterfly", "mountain", "ocean", "guitar", "library", "garden", "castle", "river"}
	hardWords   = []string{"phenomenon", "extravagant", "serendipity", "quintessential", "ephemeral", "luminous", "mellifluous", "ubiquitous", "resplendent", "effervescent"}
)

func main() {
	fmt.Println("Welcome to the Word Game!")
	fmt.Println("==========================")
	
	// Get difficulty selection
	difficulty := getDifficultySelection()
	
	// Load appropriate word list based on difficulty
	var wordList []string
	switch difficulty {
	case "easy":
		wordList = easyWords
	case "medium":
		wordList = mediumWords
	case "hard":
		wordList = hardWords
	default:
		fmt.Println("Invalid difficulty selected. Using easy words by default.")
		wordList = easyWords
	}
	
	fmt.Printf("You selected %s difficulty.\n", difficulty)
	fmt.Println("Your word list:")
	for i, word := range wordList {
		fmt.Printf("%d. %s\n", i+1, word)
	}
}

func getDifficultySelection() string {
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Println("\nPlease select a difficulty level:")
		fmt.Println("1. Easy")
		fmt.Println("2. Medium")
		fmt.Println("3. Hard")
		fmt.Print("Enter your choice (1-3): ")
		
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		
		switch input {
		case "1", "easy", "Easy":
			return "easy"
		case "2", "medium", "Medium":
			return "medium"
		case "3", "hard", "Hard":
			return "hard"
		default:
			fmt.Println("Invalid input. Please enter 1, 2, or 3, or the word 'easy', 'medium', or 'hard'.")
		}
	}
}