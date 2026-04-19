# practice

Personal practice repository.

## Word Lists

The application uses text files to provide different word lists for different difficulties:
- `easy_words.txt`: Contains easy words
- `medium_words.txt`: Contains medium difficulty words
- `hard_words.txt`: Contains hard words

These files are consumed by the `LoadWords` function in `word_loader.go` which maps difficulty levels to the appropriate file and loads the words.

## Usage

To load words of a specific difficulty, call the `LoadWords` function with the difficulty level as a string:

```go
words, err := LoadWords("easy")
if err != nil {
    // handle error
}
// Use the loaded words
fmt.Printf("Loaded %d easy words\n", len(words))
```

The function supports three difficulty levels:
- "easy": loads words from `easy_words.txt`
- "medium": loads words from `medium_words.txt`
- "hard": loads words from `hard_words.txt`