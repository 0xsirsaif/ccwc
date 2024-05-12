ccwc
------

ccwc is a command-line utility written in Go that mimics the functionality of the Unix `wc` command.
It provides the number of lines, words, and characters in a given file.

## Usage

You can use ccwc with the following flags:

- `-l`: Number of lines
- `-w`: Number of words
- `-m`: Number of characters
- `-c`: Number of bytes

If no flags are provided, ccwc will default to `-l`, `-w`, and `-c`.
