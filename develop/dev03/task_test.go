package main_test

import (
	"bufio"
	"flag"
	"io"
	"os"
	"reflect"
	"testing"

	main "dev03"
)

func TestSortByKey(t *testing.T) {
	testCases := map[string]struct {
		filename string
		column   string
		key      string
		expected []string
	}{
		"test_k": {
			filename: "test_k.txt",
			column:   "2",
			key:      "-k",
			expected: []string{"apple 10 red", "banana 5 yellow", "orange 8 orange"},
		},
		"test_r": {
			filename: "test_r.txt",
			key:      "-r",
			column:   "",
			expected: []string{"owl", "duck", "dog", "cat", "beaver"},
		},
		"test_u": {
			filename: "test_u.txt",
			key:      "-u",
			column:   "",
			expected: []string{"_no", "no_repeat", "repeat", "still_no_repeat"},
		},
		"test_n": {
			filename: "test_n.txt",
			key:      "-n",
			column:   "",
			expected: []string{"0", "1", "2", "3", "4", "7", "10"},
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			expected := testCase.expected

			var result []string
			if testCase.column == "" {
				result = sortFile(testCase.filename, testCase.key)
			} else {
				result = sortFile(testCase.filename, testCase.key, testCase.column)
			}

			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Ожидаемый результат %v, полученный %v", expected, result)
			}
		})
	}
}

func sortFile(filename string, args ...string) []string {
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

	main.SortFile()

	file, err := os.Open("sorted_output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}
