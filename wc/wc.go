package wc

import (
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"
)

func checkErr(e error, msg string) {
	if e != nil {
		fmt.Println(msg, e)
		os.Exit(1)
	}
}

func readFile(filePath string) []byte {
	var fileData []byte
	var err error

	if filePath == "-" {
		fileData, err = io.ReadAll(os.Stdin)
	} else {
		fileData, err = os.ReadFile("./" + filePath)
	}

	checkErr(err, "Error loading File:")

	return fileData
}

func Ccwc() {
	lFlag := flag.Bool("l", false, "Num Of Lines")
	wFlag := flag.Bool("w", false, "Num Of Words")
	mFlag := flag.Bool("m", false, "Num Of Characters")
	cFlag := flag.Bool("c", false, "Num Of bytes")

	flag.Parse()

	var byteCount, lineCount, wordCount, charCount int
	var printByte, printLine, printWord, printChars bool
	var input string

	// With no FILE, or when FILE is -, read standard input.
	if len(flag.Args()) < 1 || flag.Args()[0] == "-" {
		input = "-"
	} else {
		input = flag.Args()[0]
	}

	printLine = *lFlag == true
	printWord = *wFlag == true
	printChars = *mFlag == true
	printByte = *cFlag == true

	// check the default option, i.e. no options are provided, -c -l -w
	allFalse := !printLine && !printWord && !printByte && !printChars
	if allFalse {
		printByte = true
		printLine = true
		printWord = true
	}

	fileData := readFile(input)

	if printLine {
		for _, data := range fileData {
			if data == '\n' {
				lineCount++
			}
		}
		fmt.Printf("  %d", lineCount)
	}

	if printWord {
		inWord := false
		for _, data := range fileData {
			if unicode.IsSpace(rune(data)) {
				if inWord {
					wordCount++
					inWord = false
				}
			} else {
				inWord = true
			}
		}
		if inWord {
			wordCount++
		}

		fmt.Printf("  %d", wordCount)
	}

	if printChars {
		for range string(fileData) {
			charCount++
		}

		fmt.Printf("  %d", charCount)
	}

	if printByte {
		if input == "-" {
			byteCount = len(fileData)
		} else {
			fileInfo, err := os.Stat(input)
			checkErr(err, "Error reading FileInfo:")
			byteCount = int(fileInfo.Size())
		}
		fmt.Printf("  %d", byteCount)
	}

	if input != "-" {
		fmt.Printf("  %s\n", input)
	}
}
