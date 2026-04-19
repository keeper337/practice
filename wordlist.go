package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// WordList represents a collection of words organized by difficulty level
type WordList struct {
	Easy   []string
	Medium []string
	Hard   []string
}

// NewWordList creates a new WordList with sample data
func NewWordList() *WordList {
	return &WordList{
		Easy: []string{
			"cat", "dog", "sun", "car", "book",
			"tree", "bird", "fish", "cake", "moon",
		},
		Medium: []string{
			"computer", "elephant", "beautiful", "wonderful", "delicious",
			"fantastic", "amazing", "incredible", "extraordinary", "magnificent",
		},
		Hard: []string{
			"phenomenon", "ubiquitous", "serendipity", "quintessential", "ephemeral",
			"mellifluous", "luminous", "effervescent", "resplendent", "ethereal",
		},
	}
}

// GetWordsByDifficulty returns words based on the specified difficulty level
func (wl *WordList) GetWordsByDifficulty(difficulty string) []string {
	switch strings.ToLower(difficulty) {
	case "easy":
		return wl.Easy
	case "medium":
		return wl.Medium
	case "hard":
		return wl.Hard
	default:
		return []string{}
	}
}

// GetRandomWordByDifficulty returns a random word from the specified difficulty level
func (wl *WordList) GetRandomWordByDifficulty(difficulty string) string {
	words := wl.GetWordsByDifficulty(difficulty)
	if len(words) == 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

// LoadWordList loads word lists based on difficulty levels
func LoadWordList() *WordList {
	return NewWordList()
}

func main() {
	wordList := LoadWordList()

	fmt.Println("Easy words:", wordList.GetWordsByDifficulty("easy"))
	fmt.Println("Medium words:", wordList.GetWordsByDifficulty("medium"))
	fmt.Println("Hard words:", wordList.GetWordsByDifficulty("hard"))

	fmt.Println("Random easy word:", wordList.GetRandomWordByDifficulty("easy"))
	fmt.Println("Random medium word:", wordList.GetRandomWordByDifficulty("medium"))
	fmt.Println("Random hard word:", wordList.GetRandomWordByDifficulty("hard"))
}