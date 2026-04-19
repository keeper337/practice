# Hangman Game Play Guide

## Overview
This guide provides step-by-step instructions for playing the Hangman game. The goal is to guess a hidden word by suggesting letters within a limited number of attempts.

## How to Play
1. A random word is selected and displayed as a series of blanks (underscores) representing each letter.
2. Suggest a letter you believe might be in the word.
3. If the letter is correct, it will be revealed in its correct position(s).
4. If the letter is incorrect, a part of the hangman drawing will be added.
5. Continue guessing letters until:
   - You correctly guess the entire word (you win), or
   - The hangman is fully drawn (you lose).

## Step-by-Step Gameplay Example

### Starting the Game
- A hidden word is selected (e.g., "hangman").
- Display shows: `_ _ _ _ _ _ _`
- Number of incorrect guesses allowed: 6

### Making a Guess
1. Player suggests a letter, e.g., 'a'
2. If 'a' is in the word:
   - Display updates to: `_ a _ _ _ _ _`
3. If 'a' is not in the word:
   - A part of the hangman drawing is added
   - Incorrect guesses counter increases

### Continuing Play
- Continue guessing letters until either:
  - All letters are correctly guessed (win)
  - Hangman drawing is complete (lose)

## Tips for Success
1. Start with common vowels (a, e, i, o, u) as they appear frequently.
2. Look for repeated letters and their positions.
3. Keep track of previously guessed letters to avoid repeats.
4. Consider the word length when making guesses.

## Game End Conditions
- **Win**: All letters in the hidden word are correctly guessed.
- **Lose**: The hangman drawing is completed with too many incorrect guesses.

## Example Round
1. Hidden word: "computer"
2. Initial display: `_ _ _ _ _ _ _ _`
3. Player guesses 'e' → `_ e _ _ _ _ _ _`
4. Player guesses 'o' → `_ e o _ _ _ _ _`
5. Continue until win or loss.

## Conclusion
Hangman is a classic word-guessing game that helps improve vocabulary and spelling skills. Enjoy the challenge!