import pytest

from hangman import GameStatus, HangmanGame


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
