# Hangman Game Installation Instructions

This document provides step-by-step instructions to install and run the Hangman game on different operating systems.

## Prerequisites

Before installing the Hangman game, ensure you have the following:

- Go (Golang) version 1.19 or higher installed
- A terminal or command prompt

## Installation Steps

### Step 1: Clone the Repository

First, clone the repository to your local machine using Git:

```bash
git clone <repository-url>
cd <repository-name>
```

Replace `<repository-url>` with the actual URL of the repository and `<repository-name>` with the name of the cloned directory.

### Step 2: Build the Game

Navigate to the project directory and build the game using Go:

```bash
go build -o hangman main.go
```

This command compiles the source code into an executable named `hangman`.

### Step 3: Run the Game

After building, run the game with the following command:

```bash
./hangman
```

## Running on Different Operating Systems

### Linux/macOS

1. Open a terminal.
2. Follow the steps above to clone, build, and run the game.

### Windows

1. Open Command Prompt or PowerShell.
2. Follow the steps above to clone, build, and run the game.

## Troubleshooting

If you encounter any issues during installation or execution:

- Ensure Go is correctly installed by running `go version`.
- Make sure the repository URL is correct.
- Check that all required dependencies are installed.

## Conclusion

You have successfully installed and run the Hangman game. Enjoy playing!