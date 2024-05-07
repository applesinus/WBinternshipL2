package main_test

import (
	"bufio"
	"flag"
	"io"
	"os"
	"reflect"
	"testing"

	main "dev06"
)

func TestManGrep(t *testing.T) {
	testCases := map[string]struct {
		filename string
		args     []string
		expected []string
	}{
		"test_d": {
			filename: "test_d.txt",
			args:     []string{"-f", "1,2,3", "-d", ","},
			expected: []string{"1,2,3\n", "4,5,6\n"},
		},
		"test_f": {
			filename: "test_f.txt",
			args:     []string{"-f", "1,3"},
			expected: []string{"1\t3\n", "4\t6\n"},
		},
		"test_s": {
			filename: "test_s.txt",
			args:     []string{"-f", "1,2,3", "-s"},
			expected: []string{},
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

	inputFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	tempInputFile, err := os.CreateTemp("", "input_*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempInputFile.Name())

	if _, err := io.Copy(tempInputFile, inputFile); err != nil {
		panic(err)
	}
	tempInputFile.Close()

	inputFile, err = os.Open(tempInputFile.Name())
	if err != nil {
		panic(err)
	}

	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()
	os.Stdin = inputFile
	os.Args = append([]string{"task.go"}, args...)

	oldArgs := flag.CommandLine
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	defer func() { flag.CommandLine = oldArgs }()

	flag.Usage = func() {}
	flag.CommandLine.SetOutput(io.Discard)
	flag.CommandLine.Usage = func() {}
	flag.CommandLine.Parse(os.Args)
	flag.CommandLine.SetOutput(io.Discard)

	main.ManCut()

	if _, err := tempFile.Seek(0, 0); err != nil {
		panic(err)
	}
	result := make([]string, 0)
	scanner := bufio.NewScanner(tempFile)
	for scanner.Scan() {
		result = append(result, scanner.Text()+"\n")
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
