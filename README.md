# Difficulty Validator

This project provides a Go implementation for validating and scoring difficulty levels based on multiple factors.

## Features

- Validate if a difficulty level is within acceptable bounds (0-10)
- Normalize difficulty levels to a [0, 1] range
- Calculate a normalized score based on difficulty, effort, and time

## Usage

```go
validator := &DifficultyValidator{}

// Validate difficulty level
isValid := validator.ValidateDifficulty(7) // true

// Normalize difficulty level
normalized := validator.NormalizeDifficulty(7) // 0.7

// Calculate score
score := validator.CalculateScore(8, 6, 5) // 0.64
```