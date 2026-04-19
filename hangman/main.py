#!/usr/bin/env python3

import random

def get_random_word():
    words = ["python", "hangman", "computer", "programming"]
    return random.choice(words)

def display_word(word, guessed_letters):
    return ''.join([letter if letter in guessed_letters else '_' for letter in word])

def main():
    print("Welcome to Hangman!")
    word = get_random_word()
    guessed_letters = set()
    attempts = 6

    while attempts > 0:
        print(f"Word: {display_word(word, guessed_letters)}")
        guess = input("Guess a letter: ").lower()

        if guess in guessed_letters:
            print("You already guessed that letter.")
            continue

        guessed_letters.add(guess)

        if guess not in word:
            attempts -= 1
            print(f"Wrong guess! Attempts left: {attempts}")
        else:
            print("Good guess!")

        if all(letter in guessed_letters for letter in word):
            print("Congratulations! You've guessed the word!")
            break

    if attempts == 0:
        print(f"Game over! The word was: {word}")

if __name__ == "__main__":
    main()