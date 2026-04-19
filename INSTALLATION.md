# Installation Guide

This document provides clear installation steps for the project.

## Prerequisites

Before installing, ensure you have the following:

- Go 1.21 or higher
- Git
- A Unix-like environment (Linux/macOS) or Windows with WSL

## Installation Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/practice.git
   cd practice
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build -o practice .
   ```

4. Run the application:
   ```bash
   ./practice
   ```

## Platform Compatibility

The installation steps are compatible with:

- Linux (Ubuntu, CentOS)
- macOS (Intel and Apple Silicon)
- Windows (with WSL)

Ensure that you have the required tools installed for your platform.