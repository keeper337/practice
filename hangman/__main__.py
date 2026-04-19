#!/usr/bin/env python3

import sys
import os

# Add the directory containing hangman.py to the Python path
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

from hangman.main import main

if __name__ == "__main__":
    main()