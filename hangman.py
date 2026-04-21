from dataclasses import dataclass, field
from enum import Enum
from typing import Set


class GameStatus(str, Enum):
    IN_PROGRESS = "in_progress"
    WON = "won"
    LOST = "lost"


@dataclass
class HangmanGame:
    secret_word: str
    max_wrong_guesses: int = 6
    guessed_letters: Set[str] = field(default_factory=set)
    wrong_guesses: int = 0
    status: GameStatus = GameStatus.IN_PROGRESS

    def __post_init__(self) -> None:
        if self.max_wrong_guesses <= 0:
            raise ValueError("max_wrong_guesses must be greater than zero")
        self._initialize_round(self.secret_word)

    @property
    def remaining_attempts(self) -> int:
        return self.max_wrong_guesses - self.wrong_guesses

    @property
    def masked_word(self) -> str:
        return "".join(
            char if char in self.guessed_letters else "_"
            for char in self.secret_word
        )

    def guess_letter(self, letter: str) -> bool:
        if self.status is not GameStatus.IN_PROGRESS:
            raise RuntimeError("round is complete, start a new round")
        if not isinstance(letter, str):
            raise ValueError("guess must be a single alphabetic letter")
        if not letter or len(letter) != 1 or not letter.isalpha():
            raise ValueError("guess must be a single alphabetic letter")

        normalized = letter.lower()
        is_duplicate_guess = normalized in self.guessed_letters
        if is_duplicate_guess:
            # Duplicate guesses are idempotent: no state mutation, same result.
            is_correct = normalized in self.secret_word
            self._refresh_status()
            return is_correct

        self.guessed_letters.add(normalized)
        is_correct = normalized in self.secret_word
        if not is_correct:
            self.wrong_guesses += 1

        self._refresh_status()
        return is_correct

    def start_new_round(
        self,
        secret_word: str,
        max_wrong_guesses: int | None = None,
    ) -> None:
        if self.status is GameStatus.IN_PROGRESS:
            raise RuntimeError("cannot start a new round while current round is in progress")
        normalized_secret_word = self._normalize_secret_word(secret_word)
        next_max_wrong_guesses = self.max_wrong_guesses
        if max_wrong_guesses is not None:
            if max_wrong_guesses <= 0:
                raise ValueError("max_wrong_guesses must be greater than zero")
            next_max_wrong_guesses = max_wrong_guesses

        # Apply state transitions only after all validation passes.
        self.max_wrong_guesses = next_max_wrong_guesses
        self._initialize_round(normalized_secret_word)

    def reset_round(
        self,
        secret_word: str,
        max_wrong_guesses: int | None = None,
    ) -> None:
        # Backward-compatible alias for callers using the previous API.
        self.start_new_round(secret_word, max_wrong_guesses=max_wrong_guesses)

    def _refresh_status(self) -> None:
        if all(char in self.guessed_letters for char in self.secret_word):
            self.status = GameStatus.WON
        elif self.wrong_guesses >= self.max_wrong_guesses:
            self.status = GameStatus.LOST
        else:
            self.status = GameStatus.IN_PROGRESS

    def _initialize_round(self, secret_word: str) -> None:
        self.secret_word = self._normalize_secret_word(secret_word)
        self.guessed_letters = set()
        self.wrong_guesses = 0
        self.status = GameStatus.IN_PROGRESS

    @staticmethod
    def _normalize_secret_word(secret_word: str) -> str:
        if not isinstance(secret_word, str):
            raise ValueError("secret_word must contain only alphabetic characters")
        normalized = secret_word.strip().lower()
        if not normalized or not normalized.isalpha():
            raise ValueError("secret_word must contain only alphabetic characters")
        return normalized
