# practice

Personal practice repository.

## Hangman Game

This is a simple implementation of the Hangman game in Go. The game allows players to guess letters in a hidden word, with a limited number of incorrect guesses before losing.

### Features

- Guess letters in a word
- Track correct and incorrect guesses
- Display current state of the word with guessed letters revealed and underscores for unguessed letters
- Win or lose conditions based on guesses
- Visual representation of hangman progress

### How to Run

To run the game, you need to have Go installed on your system. Then execute:

```bash
go run hangman.go
```

### Game Logic

The game starts with a predefined word (in this case "hangman"). Players can guess letters one at a time. The game will display the current state of the word with guessed letters revealed and underscores for unguessed letters.

If a player guesses incorrectly, their wrong guess count increases. If they exceed the maximum number of allowed wrong guesses (default is 6), they lose the game.

## Improvements

- Added support for custom words
- Improved error handling
- Enhanced user interface with better feedback messages
- Fixed bug in letter guessing logic
- Implemented visual representation of hangman progress