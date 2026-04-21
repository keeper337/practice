import pytest

from hangman import GameStatus, HangmanGame


def test_correct_guess_reveals_letters_without_consuming_attempt() -> None:
    game = HangmanGame("letter")

    result = game.guess_letter("e")

    assert result is True
    assert game.masked_word == "_e__e_"
    assert game.wrong_guesses == 0
    assert game.remaining_attempts == 6
    assert game.status is GameStatus.IN_PROGRESS


def test_incorrect_guess_consumes_attempt_and_tracks_letter() -> None:
    game = HangmanGame("letter")

    result = game.guess_letter("x")

    assert result is False
    assert game.masked_word == "______"
    assert game.wrong_guesses == 1
    assert game.remaining_attempts == 5
    assert game.guessed_letters == {"x"}
    assert game.status is GameStatus.IN_PROGRESS


def test_duplicate_guess_is_idempotent_for_correct_letter() -> None:
    game = HangmanGame("banana")
    first = game.guess_letter("a")
    before = (set(game.guessed_letters), game.wrong_guesses, game.status, game.masked_word)

    second = game.guess_letter("a")

    assert first is True
    assert second is True
    assert (set(game.guessed_letters), game.wrong_guesses, game.status, game.masked_word) == before


def test_duplicate_guess_is_idempotent_for_incorrect_letter() -> None:
    game = HangmanGame("banana")
    first = game.guess_letter("x")
    before = (set(game.guessed_letters), game.wrong_guesses, game.status, game.masked_word)

    second = game.guess_letter("x")

    assert first is False
    assert second is False
    assert (set(game.guessed_letters), game.wrong_guesses, game.status, game.masked_word) == before


def test_winning_guess_sets_won_state_and_blocks_further_guesses() -> None:
    game = HangmanGame("go")
    game.guess_letter("g")
    assert game.status is GameStatus.IN_PROGRESS

    result = game.guess_letter("o")

    assert result is True
    assert game.masked_word == "go"
    assert game.status is GameStatus.WON
    with pytest.raises(RuntimeError, match="round is complete"):
        game.guess_letter("x")


def test_losing_guess_sets_lost_state_and_blocks_further_guesses() -> None:
    game = HangmanGame("a", max_wrong_guesses=2)
    game.guess_letter("x")
    assert game.status is GameStatus.IN_PROGRESS

    result = game.guess_letter("y")

    assert result is False
    assert game.wrong_guesses == 2
    assert game.status is GameStatus.LOST
    with pytest.raises(RuntimeError, match="round is complete"):
        game.guess_letter("z")


def _win_round(game: HangmanGame) -> None:
    for letter in sorted(set(game.secret_word)):
        game.guess_letter(letter)


def _lose_round(game: HangmanGame) -> None:
    for letter in ["x", "y", "z", "q", "v", "w", "k", "j"]:
        if game.status is GameStatus.LOST:
            break
        if letter not in game.secret_word:
            game.guess_letter(letter)


def test_start_new_round_rejected_while_round_in_progress() -> None:
    game = HangmanGame("apple")
    game.guess_letter("a")

    with pytest.raises(RuntimeError, match="in progress"):
        game.start_new_round("berry", max_wrong_guesses=4)

    assert game.secret_word == "apple"
    assert game.masked_word == "a____"
    assert game.max_wrong_guesses == 6
    assert game.status is GameStatus.IN_PROGRESS


def test_start_new_round_allowed_after_win_and_resets_state() -> None:
    game = HangmanGame("pear")
    _win_round(game)
    assert game.status is GameStatus.WON

    game.start_new_round("melon", max_wrong_guesses=5)

    assert game.secret_word == "melon"
    assert game.masked_word == "_____"
    assert game.guessed_letters == set()
    assert game.wrong_guesses == 0
    assert game.max_wrong_guesses == 5
    assert game.status is GameStatus.IN_PROGRESS


def test_start_new_round_allowed_after_loss_and_resets_state() -> None:
    game = HangmanGame("abc", max_wrong_guesses=2)
    _lose_round(game)
    assert game.status is GameStatus.LOST

    game.start_new_round("kite")

    assert game.secret_word == "kite"
    assert game.masked_word == "____"
    assert game.guessed_letters == set()
    assert game.wrong_guesses == 0
    assert game.max_wrong_guesses == 2
    assert game.status is GameStatus.IN_PROGRESS


def test_reset_round_alias_uses_same_completed_round_scope() -> None:
    game = HangmanGame("cloud")

    with pytest.raises(RuntimeError, match="in progress"):
        game.reset_round("storm")

    _win_round(game)
    game.reset_round("storm")

    assert game.secret_word == "storm"
    assert game.status is GameStatus.IN_PROGRESS


def test_start_new_round_validation_is_atomic_on_invalid_secret() -> None:
    game = HangmanGame("pear")
    _win_round(game)

    with pytest.raises(ValueError, match="secret_word"):
        game.start_new_round("123", max_wrong_guesses=2)

    assert game.secret_word == "pear"
    assert game.max_wrong_guesses == 6
    assert game.status is GameStatus.WON


def test_reset_round_alias_validation_is_atomic_on_invalid_secret() -> None:
    game = HangmanGame("abc", max_wrong_guesses=2)
    _lose_round(game)

    with pytest.raises(ValueError, match="secret_word"):
        game.reset_round("!!!", max_wrong_guesses=4)

    assert game.secret_word == "abc"
    assert game.max_wrong_guesses == 2
    assert game.status is GameStatus.LOST
