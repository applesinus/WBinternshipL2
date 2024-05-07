package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	ManGrep()
}

// ManGrep - task
func ManGrep() {
	// Parsing flags
	after := flag.Int("A", 0, "Print +N lines after each match")
	before := flag.Int("B", 0, "Print +N lines before each match")
	context := flag.Int("C", 0, "Print ±N lines around each match")
	count := flag.Bool("c", false, "Print only a count of selected lines")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions")
	invert := flag.Bool("v", false, "Selected lines are those not matching any of the specified patterns")
	fixed := flag.Bool("F", false, "Interpret patterns as fixed strings, not regular expressions")
	lineNum := flag.Bool("n", false, "Prefix each line of output with the 1-based line number within its input file")

	flag.Parse()

	// Getting pattern
	pattern := flag.Arg(0)

	// If no pattern is provided, print usage and exit
	if pattern == "" {
		fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Function for matching
	matcher := func(line string) bool {
		if *fixed {
			if *ignoreCase {
				return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			}
			return strings.Contains(line, pattern)
		}
		if *ignoreCase {
			return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		}
		return strings.Contains(line, pattern)
	}

	// Function for printing lines
	printLine := func(line string, lineNumber int) {
		if *lineNum {
			fmt.Printf("%d:%s\n", lineNumber, line)
		} else {
			fmt.Println(line)
		}
	}

	var selectedCount int

	// If no files are provided, read from stdin
	files := flag.Args()[1:]
	if len(files) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for lineNumber := 1; scanner.Scan(); lineNumber++ {
			line := scanner.Text()

			// Adding line to context
			if len(lines) > 0 {
				lines = append(lines, line)
			}

			// After, before and context flags
			if (matcher(line) && !*invert) || (!matcher(line) && *invert) {
				// before
				for _, prevLine := range lines {
					printLine(prevLine, lineNumber-len(lines))
				}

				printLine(line, lineNumber)

				// after
				for i := 1; i <= *after; i++ {
					if scanner.Scan() {
						nextLine := scanner.Text()
						printLine(nextLine, lineNumber+i)
					}
				}

				lines = nil

				selectedCount++
			} else if *before > 0 || *context > 0 {
				// before
				lines = append(lines, line)
				if len(lines) > *before+*context {
					lines = lines[1:]
				}
			}
		}
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
				continue
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			var lines []string
			for lineNumber := 1; scanner.Scan(); lineNumber++ {
				line := scanner.Text()

				if len(lines) > 0 {
					lines = append(lines, line)
				}

				// After, before and context flags
				if (matcher(line) && !*invert) || (!matcher(line) && *invert) {
					// before
					for _, prevLine := range lines {
						printLine(prevLine, lineNumber-len(lines))
					}

					printLine(line, lineNumber)

					// after
					for i := 1; i <= *after; i++ {
						if scanner.Scan() {
							nextLine := scanner.Text()
							printLine(nextLine, lineNumber+i)
						}
					}

					lines = nil

					selectedCount++
				} else if *before > 0 || *context > 0 {
					// before
					lines = append(lines, line)
					if len(lines) > *before+*context {
						lines = lines[1:]
					}
				}
			}
		}
	}

	// Count flag
	if *count {
		fmt.Printf("Count of selected lines: %d\n", selectedCount)
	}
}
