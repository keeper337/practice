package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Game represents the state of our game
type Game struct {
	PlayerName string
	Score      int
	Lives      int
	Level      int
}

// NewGame creates a new game instance
func NewGame() *Game {
	return &Game{
		PlayerName: "Player",
		Score:      0,
		Lives:      3,
		Level:      1,
	}
}

// DisplayGameState shows the current state of the game
func (g *Game) DisplayGameState() {
	fmt.Println("=====================================")
	fmt.Println("              GAME STATE")
	fmt.Println("=====================================")
	fmt.Printf("Player: %s\n", g.PlayerName)
	fmt.Printf("Score: %d\n", g.Score)
	fmt.Printf("Lives: %d\n", g.Lives)
	fmt.Printf("Level: %d\n", g.Level)
	fmt.Println("=====================================")
}

// GetInput prompts the user for input with clear instructions
func GetInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimSpace(input)
}

// ValidateAndProcessInput processes user input with error handling
func (g *Game) ValidateAndProcessInput(input string) bool {
	switch strings.ToLower(input) {
	case "up", "u":
		g.Score += 10
		fmt.Println("Moved up! +10 points")
		return true
	case "down", "d":
		g.Score += 5
		fmt.Println("Moved down! +5 points")
		return true
	case "left", "l":
		g.Score += 7
		fmt.Println("Moved left! +7 points")
		return true
	case "right", "r":
		g.Score += 8
		fmt.Println("Moved right! +8 points")
		return true
	case "quit", "q":
		fmt.Println("Thanks for playing!")
		return false
	case "help", "h":
		g.ShowHelp()
		return true
	default:
		fmt.Println("Invalid input! Please try again.")
		fmt.Println("Valid commands: up/u, down/d, left/l, right/r, help/h, quit/q")
		return true
	}
}

// ShowHelp displays available commands
func (g *Game) ShowHelp() {
	fmt.Println("\n=== HELP ===")
	fmt.Println("Commands:")
	fmt.Println("  up/u     - Move up")
	fmt.Println("  down/d   - Move down")
	fmt.Println("  left/l   - Move left")
	fmt.Println("  right/r  - Move right")
	fmt.Println("  help/h   - Show this help")
	fmt.Println("  quit/q   - Quit the game")
	fmt.Println("============\n")
}

// UpdateLevel checks if player should advance to next level
func (g *Game) UpdateLevel() {
	if g.Score >= g.Level*100 {
		g.Level++
		fmt.Printf("Congratulations! You've reached Level %d!\n", g.Level)
	}
}

// PlayGame runs the main game loop
func (g *Game) PlayGame() {
	fmt.Println("=====================================")
	fmt.Println("           WELCOME TO THE GAME")
	fmt.Println("=====================================")
	
	// Get player name
	name := GetInput("Enter your name: ")
	if name != "" {
		g.PlayerName = name
	}
	
	fmt.Println("Type 'help' for available commands.")
	fmt.Println()
	
	for {
		g.DisplayGameState()
		
		input := GetInput("Enter command (up/down/left/right/quit/help): ")
		
		if !g.ValidateAndProcessInput(input) {
			break
		}
		
		g.UpdateLevel()
		
		// Add a small delay to make it easier to read
		fmt.Println()
	}
}

func main() {
	game := NewGame()
	game.PlayGame()
}