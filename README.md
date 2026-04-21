# practice

Personal practice repository.

## Hangman game state model

This repository now includes a minimal game-state implementation in `hangman.py`.

### Behaviors covered

- Game initializes with a hidden word and max wrong guesses.
- Player can submit letter guesses one at a time via `guess_letter`.
- Correct guesses reveal all matching positions through `masked_word`.
- Incorrect guesses increment `wrong_guesses` and reduce `remaining_attempts`.
- Game ends with `GameStatus.WON` when all letters are revealed.
- Game ends with `GameStatus.LOST` when attempts reach zero.
- `start_new_round(secret_word)` starts a new round after win/loss.
- `reset_round(secret_word)` remains available as a compatibility alias.
