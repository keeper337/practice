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
        self.secret_word = self._normalize_secret_word(self.secret_word)
        if self.max_wrong_guesses <= 0:
            raise ValueError("max_wrong_guesses must be greater than zero")

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
        if not letter or len(letter) != 1 or not letter.isalpha():
            raise ValueError("guess must be a single alphabetic letter")

        normalized = letter.lower()
        if normalized in self.guessed_letters:
            return normalized in self.secret_word

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
        self.secret_word = self._normalize_secret_word(secret_word)
        if max_wrong_guesses is not None:
            if max_wrong_guesses <= 0:
                raise ValueError("max_wrong_guesses must be greater than zero")
            self.max_wrong_guesses = max_wrong_guesses
        self.guessed_letters.clear()
        self.wrong_guesses = 0
        self.status = GameStatus.IN_PROGRESS

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

    @staticmethod
    def _normalize_secret_word(secret_word: str) -> str:
        normalized = secret_word.strip().lower()
        if not normalized or not normalized.isalpha():
            raise ValueError("secret_word must contain only alphabetic characters")
        return normalized
