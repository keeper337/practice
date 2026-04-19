# Hangman Game Rules

## Objective
Guess the hidden word by suggesting letters one at a time. The player wins by correctly guessing all letters in the word before making too many incorrect guesses.

## How to Play
1. A random word from the word list is selected
2. The player is shown the word with underscores representing unguessed letters
3. The player suggests a letter (one at a time)
4. If the letter is in the word, it is revealed in its correct positions
5. If the letter is not in the word, a part of the hangman drawing is added
6. The game continues until the word is guessed or the hangman is complete

## Game Elements
- **Word List**: Contains programming-related terms
- **Wrong Guesses**: Maximum of 6 wrong guesses allowed before losing
- **Visual Hangman**: A simple ASCII hangman drawing that updates with each wrong guess
- **Win/Loss Detection**: The game automatically detects when the player wins or loses

## Controls
- Enter a single letter to make a guess
- Press Enter after typing your letter