package main

import (
	"testing"
)

func TestHangmanGame_IsWinner(t *testing.T) {
	game := NewHangmanGame()
	// Guess all letters correctly
	for _, letter := range game.Word {
		game.GuessLetter(letter)
	}
	
	if !game.IsWinner() {
		t.Errorf("Expected game to be a winner, but it wasn't")
	}
}

func TestHangmanGame_IsLoser(t *testing.T) {
	game := NewHangmanGame()
	// Make incorrect guesses until losing
	for i := 0; i < game.MaxIncorrectGuesses; i++ {
		game.GuessLetter('z') // Guess a letter not in the word
	}
	
	if !game.IsLoser() {
		t.Errorf("Expected game to be a loser, but it wasn't")
	}
}

func TestHangmanGame_IsGameOver(t *testing.T) {
	// Test win condition
	game1 := NewHangmanGame()
	for _, letter := range game1.Word {
		game1.GuessLetter(letter)
	}
	
	if !game1.IsGameOver() {
		t.Errorf("Expected game to be over (win), but it wasn't")
	}

	// Test lose condition
	game2 := NewHangmanGame()
	for i := 0; i < game2.MaxIncorrectGuesses; i++ {
		game2.GuessLetter('z')
	}
	
	if !game2.IsGameOver() {
		t.Errorf("Expected game to be over (lose), but it wasn't")
	}
}