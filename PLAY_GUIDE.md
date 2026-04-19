# How to Play Hangman

## Running the Game
To play the Hangman game, simply run:
```bash
go run main.go
```

## Game Interface
The game will display:
1. The current state of the word with guessed letters revealed
2. A visual hangman drawing that updates with each wrong guess
3. Prompts for letter input

## Example Gameplay
```
Try to guess the word by suggesting letters.

Enter a letter: a
Correct guess!

Enter a letter: b
Incorrect guess!

Enter a letter: c
Correct guess!
```

## Game End Conditions
- **Win**: All letters in the word are guessed correctly
- **Loss**: The hangman drawing is complete (6 wrong guesses)