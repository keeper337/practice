# Hangman Game Examples

This document provides examples of gameplay scenarios to help understand how the Hangman game works.

## Example 1: Winning the Game

```
Welcome to Hangman!
Try to guess the word by suggesting letters.

Word: _ _ _ _ _ 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: e
Correct guess!

Word: _ e _ _ _ 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: a
Correct guess!

Word: _ e _ a _ 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: n
Incorrect guess!

Word: _ e _ a _ 
Attempts left: 5/6

  +---+
  |   |
  O   |
      |
      |
      |
=========

Enter a letter: t
Correct guess!

Word: _ e _ a t 
Attempts left: 5/6

  +---+
  |   |
  O   |
      |
      |
      |
=========

Enter a letter: r
Correct guess!

Word: _ e _ a t 
Attempts left: 5/6

  +---+
  |   |
  O   |
      |
      |
      |
=========

Enter a letter: s
Incorrect guess!

Word: _ e _ a t 
Attempts left: 4/6

  +---+
  |   |
  O   |
  |   |
      |
      |
=========

Enter a letter: l
Correct guess!

Word: l e _ a t 
Attempts left: 4/6

  +---+
  |   |
  O   |
  |   |
      |
      |
=========

Enter a letter: c
Incorrect guess!

Word: l e _ a t 
Attempts left: 3/6

  +---+
  |   |
  O   |
 /|   |
      |
      |
=========

Enter a letter: u
Incorrect guess!

Word: l e _ a t 
Attempts left: 2/6

  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========

Enter a letter: i
Incorrect guess!

Word: l e _ a t 
Attempts left: 1/6

  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========

Enter a letter: o
Correct guess!

Word: l e o a t 
Attempts left: 1/6

  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========

Enter a letter: p
Incorrect guess!

Word: l e o a t 
Attempts left: 0/6

  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========

Game over! You lost.
The word was: lemon
```

## Example 2: Losing the Game

```
Welcome to Hangman!
Try to guess the word by suggesting letters.

Word: _ _ _ _ _ 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: x
Incorrect guess!

Word: _ _ _ _ _ 
Attempts left: 5/6

  +---+
  |   |
  O   |
      |
      |
      |
=========

Enter a letter: y
Incorrect guess!

Word: _ _ _ _ _ 
Attempts left: 4/6

  +---+
  |   |
  O   |
  |   |
      |
      |
=========

Enter a letter: z
Incorrect guess!

Word: _ _ _ _ _ 
Attempts left: 3/6

  +---+
  |   |
  O   |
 /|   |
      |
      |
=========

Enter a letter: w
Incorrect guess!

Word: _ _ _ _ _ 
Attempts left: 2/6

  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========

Enter a letter: v
Incorrect guess!

Word: _ _ _ _ _ 
Attempts left: 1/6

  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========

Enter a letter: u
Incorrect guess!

Word: _ _ _ _ _ 
Attempts left: 0/6

  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========

Game over! You lost.
The word was: grape
```

## Example 3: Winning with Fewer Attempts

```
Welcome to Hangman!
Try to guess the word by suggesting letters.

Word: _ _ _ _ _ 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: a
Correct guess!

Word: _ a _ _ _ 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: p
Correct guess!

Word: _ a _ _ _ 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: l
Correct guess!

Word: _ a _ l _ 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: e
Correct guess!

Word: _ a _ l e 
Attempts left: 6/6

  +---+
  |   |
      |
      |
      |
      |
=========

Enter a letter: n
Incorrect guess!

Word: _ a _ l e 
Attempts left: 5/6

  +---+
  |   |
  O   |
      |
      |
      |
=========

Enter a letter: t
Correct guess!

Word: _ a _ l e 
Attempts left: 5/6

  +---+
  |   |
  O   |
      |
      |
      |
=========

Congratulations! You won!
```