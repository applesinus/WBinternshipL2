package main_test

import (
	"bufio"
	"flag"
	"io"
	"os"
	"reflect"
	"testing"

	main "dev05"
)

func TestManGrep(t *testing.T) {
	testCases := map[string]struct {
		filename string
		args     []string
		expected []string
	}{
		"test_v": {
			filename: "test_v.txt",
			args:     []string{"-v", "hello"},
			expected: []string{"world", "HELLO"},
		},
		"test_i": {
			filename: "test_i.txt",
			args:     []string{"-i", "hello"},
			expected: []string{"hello", "HeLlO", "HELLO"},
		},
		"test_n": {
			filename: "test_n.txt",
			args:     []string{"-n", "hello"},
			expected: []string{"1:hello", "3:hello world"},
		},
		"test_c": {
			filename: "test_c.txt",
			args:     []string{"-c", "hello"},
			expected: []string{"hello", "hello world", "hello exam", "Count of selected lines: 3"},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			result := test(testCase.filename, testCase.args...)

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected: %v, Got: %v", testCase.expected, result)
			}
		})
	}
}

func test(filename string, args ...string) []string {
	tempFile, err := os.CreateTemp("", "output_*.txt")
	if err != nil {
		panic(err)
	}
	defer tempFile.Close()

	originalStdout := os.Stdout
	os.Stdout = tempFile

	defer func() {
		os.Stdout = originalStdout
	}()

	os.Args = append([]string{"task.go"}, args...)
	os.Args = append(os.Args, filename)

	oldArgs := flag.CommandLine
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	defer func() { flag.CommandLine = oldArgs }()

	flag.Usage = func() {}
	flag.CommandLine.SetOutput(io.Discard)
	flag.CommandLine.Usage = func() {}
	flag.CommandLine.Parse(os.Args)
	flag.CommandLine.SetOutput(io.Discard)

	main.ManGrep()

	if _, err := tempFile.Seek(0, 0); err != nil {
		panic(err)
	}
	result := make([]string, 0)
	scanner := bufio.NewScanner(tempFile)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
