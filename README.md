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
- `start_new_round(secret_word, max_wrong_guesses=None)` starts a new round after win/loss and can optionally reconfigure attempts.
- `reset_round(secret_word, max_wrong_guesses=None)` remains available as a compatibility alias.
- Starting/resetting a round is intentionally scoped to completed rounds; attempting it while a round is still in progress raises `RuntimeError`.

### Acceptance criteria mapped to tests

- Completed-round-only reset scope:
  `test_start_new_round_rejected_while_round_in_progress`,
  `test_start_new_round_allowed_after_win_and_resets_state`,
  `test_start_new_round_allowed_after_loss_and_resets_state`,
  `test_reset_round_alias_uses_same_completed_round_scope`.
- No partial mutation on invalid new-round secret:
  `test_start_new_round_validation_is_atomic_on_invalid_secret`,
  `test_reset_round_alias_validation_is_atomic_on_invalid_secret`.
