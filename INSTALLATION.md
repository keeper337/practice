# Installation Instructions for Hangman Game

This document provides clear instructions on how to install and run the Hangman game.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.19 or higher)
- A terminal or command prompt

## Installation Steps

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

2. **Build the Game**
   ```bash
   go build -o hangman main.go
   ```

3. **Run the Game**
   ```bash
   ./hangman
   ```

## Cross-Platform Compatibility

These instructions work on Windows, macOS, and Linux. Ensure that you have Go installed and properly configured in your system's PATH.

## Troubleshooting

If you encounter any issues:

- Make sure Go is correctly installed by running `go version`.
- Verify that the source files are correctly placed in the project directory.
- Check for any syntax errors in the code using `go vet`.

For further assistance, refer to the [Go documentation](https://golang.org/doc/).