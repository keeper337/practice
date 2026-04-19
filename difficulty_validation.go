package main

import (
	"fmt"
	"math"
)

// DifficultyValidator provides methods for validating and scoring difficulty levels.
type DifficultyValidator struct{}

// ValidateDifficulty checks if the provided difficulty level is within acceptable bounds.
func (dv *DifficultyValidator) ValidateDifficulty(level float64) bool {
	return level >= 0 && level <= 10
}

// NormalizeDifficulty scales the difficulty level to a [0, 1] range.
func (dv *DifficultyValidator) NormalizeDifficulty(level float64) float64 {
	if level < 0 {
		return 0
	}
	if level > 10 {
		return 1
	}
	return level / 10
}

// CalculateScore computes a normalized score based on difficulty and other factors.
func (dv *DifficultyValidator) CalculateScore(difficulty, effort, time float64) float64 {
	normalizedDifficulty := dv.NormalizeDifficulty(difficulty)
	normalizedEffort := math.Min(effort/10, 1.0)
	normalizedTime := math.Min(time/10, 1.0)

	// Weighted average: difficulty (50%), effort (30%), time (20%)
	score := (normalizedDifficulty*0.5 + normalizedEffort*0.3 + normalizedTime*0.2)
	return score
}

func main() {
	validator := &DifficultyValidator{}

	// Example usage
	fmt.Println("Validating difficulty level 7:", validator.ValidateDifficulty(7))
	fmt.Println("Normalizing difficulty level 7:", validator.NormalizeDifficulty(7))
	fmt.Println("Calculating score for difficulty=8, effort=6, time=5:", validator.CalculateScore(8, 6, 5))
}