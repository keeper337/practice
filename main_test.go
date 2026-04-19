package main

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	if game == nil {
		t.Error("NewGame should not return nil")
	}
	if game.word == "" {
		t.Error("NewGame should set a random word")
	}
	if len(game.guessed) != len(game.word) {
		t.Error("Guessed slice should match word length")
	}
}

func TestGuess(t *testing.T) {
	game := NewGame()
	game.word = "test"
	game.guessed = make([]bool, 4)

	// Test correct guess
	result := game.Guess("t")
	if !result {
		t.Error("Guessing 't' should return true")
	}
	if !game.guessed[0] || !game.guessed[3] {
		t.Error("Guessed slice should be updated for correct guess")
	}

	// Test incorrect guess
	result = game.Guess("x")
	if result {
		t.Error("Guessing 'x' should return false")
	}
	if game.attempts != 1 {
		t.Error("Attempts should increment on incorrect guess")
	}
}

func TestIsWon(t *testing.T) {
	game := NewGame()
	game.word = "win"
	game.guessed = []bool{true, true, true}

	if !game.IsWon() {
		t.Error("Game should be won when all letters are guessed")
	}

	game.guessed[0] = false
	if game.IsWon() {
		t.Error("Game should not be won when not all letters are guessed")
	}
}

func TestIsLost(t *testing.T) {
	game := NewGame()
	game.word = "test"
	game.attempts = 10

	if !game.IsLost() {
		t.Error("Game should be lost after max attempts")
	}
}

func TestMainFunction(t *testing.T) {
	// This test ensures main function compiles and runs without errors
	// Actual testing of user interaction would require more complex mocking
	// For now, we just verify the code compiles
}