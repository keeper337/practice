# practice

Personal practice repository.

## Word Lists

The application uses text files to provide different word lists for different difficulties:
- `easy_words.txt`: Contains easy words
- `medium_words.txt`: Contains medium difficulty words
- `hard_words.txt`: Contains hard words

These files are consumed by the `LoadWords` function in `word_loader.go` which maps difficulty levels to the appropriate file and loads the words.