#!/usr/bin/env python3

import sys
import os

# Add the current directory to the path so we can import the game module
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

from hangman.game import start_game

if __name__ == "__main__":
    start_game()