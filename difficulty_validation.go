package main

import (
	"fmt"
	"math"
)

// DifficultyValidator validates difficulty levels for tasks
type DifficultyValidator struct{}

// ValidateDifficulty checks if the difficulty level is within acceptable bounds
func (dv *DifficultyValidator) ValidateDifficulty(difficulty float64) (bool, error) {
	if difficulty < 0 {
		return false, fmt.Errorf("difficulty cannot be negative")
	}
	
	if difficulty > 100 {
		return false, fmt.Errorf("difficulty cannot exceed 100")
	}
	
	return true, nil
}

// NormalizeDifficulty ensures difficulty is within [0, 100] range
func (dv *DifficultyValidator) NormalizeDifficulty(difficulty float64) float64 {
	if difficulty < 0 {
		return 0
	}
	if difficulty > 100 {
		return 100
	}
	return difficulty
}

// CalculateDifficultyScore computes a normalized difficulty score
func (dv *DifficultyValidator) CalculateDifficultyScore(tasks []int) (float64, error) {
	if len(tasks) == 0 {
		return 0, fmt.Errorf("no tasks provided")
	}
	
	var sum int
	for _, task := range tasks {
		if task < 0 {
			return 0, fmt.Errorf("task difficulty cannot be negative")
		}
		sum += task
	}
	
	average := float64(sum) / float64(len(tasks))
	return math.Min(average, 100), nil
}

func main() {
	validator := &DifficultyValidator{}
	
	// Test cases
	testCases := []float64{50, -10, 150, 75.5}
	
	for _, difficulty := range testCases {
		valid, err := validator.ValidateDifficulty(difficulty)
		if err != nil {
			fmt.Printf("Validation failed for %v: %v\n", difficulty, err)
		} else {
			fmt.Printf("Validation result for %v: %v\n", difficulty, valid)
		}
	}
	
	// Test normalization
	fmt.Printf("Normalized 120: %v\n", validator.NormalizeDifficulty(120))
	fmt.Printf("Normalized -10: %v\n", validator.NormalizeDifficulty(-10))
	
	// Test score calculation
	tasks := []int{30, 40, 50, 60}
	score, err := validator.CalculateDifficultyScore(tasks)
	if err != nil {
		fmt.Printf("Error calculating score: %v\n", err)
	} else {
		fmt.Printf("Average difficulty score: %v\n", score)
	}
}