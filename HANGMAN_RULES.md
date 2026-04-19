# Hangman Game Rules

## Objective
The objective of Hangman is to guess a hidden word by suggesting letters within a certain number of guesses.

## How to Play
1. A secret word is chosen and displayed as a series of underscores ( _ ) representing each letter.
2. Players guess one letter at a time.
3. If the guessed letter is in the word, it is revealed in its correct position(s).
4. If the guessed letter is not in the word, a part of the hangman figure is drawn.
5. The game continues until:
   - The player guesses all letters in the word (they win), or
   - The hangman figure is completed (they lose).

## Example Gameplay
Let's say the secret word is "hangman":

1. Initial display: `_ _ _ _ _ _ _`
2. Player guesses 'a':
   - Word becomes: `_ a _ _ a _ _`
3. Player guesses 'n':
   - Word becomes: `_ a n _ a n _`
4. Player guesses 'g':
   - Word becomes: `g a n _ a n _`
5. Player guesses 'h':
   - Word becomes: `h a n _ a n _`
6. Player guesses 'm':
   - Word becomes: `h a n m a n _`
7. Player guesses 'x':
   - A part of the hangman is drawn (incorrect guess)
8. Player guesses 'l':
   - Word becomes: `h a n m a n l`
9. Player wins by guessing all letters.

## Tips
- Start with common vowels (a, e, i, o, u) and consonants.
- Keep track of guessed letters to avoid repeating mistakes.
- Try to identify patterns in the word based on revealed letters.

## Conclusion
Hangman is a classic word-guessing game that helps improve vocabulary and spelling skills. Enjoy playing!