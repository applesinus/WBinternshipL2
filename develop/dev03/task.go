package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Отсортировать строки в файле по аналогии с консольной утилитой sort
(man sort — смотрим описание и основные параметры): на входе подается
файл из несортированными строками, на выходе — файл с отсортированными.

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	SortFile()
}

// SortFile sorts the lines in a file
func SortFile() {
	// Getting flags
	column := flag.Int("k", 0, "Указание колонки для сортировки")
	numerical := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	// Opening a file
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Reading the file
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Determining the sorting function based on the flags
	comparator := func(i, j int) bool {
		a, b := lines[i], lines[j]

		if *column > 0 {
			// Splitting the strings into columns if needed
			aCols := strings.Fields(a)
			bCols := strings.Fields(b)

			if *column <= len(aCols) && *column <= len(bCols) {
				a = aCols[*column-1]
				b = bCols[*column-1]
			}
		}

		if *numerical {
			// Converting strings to numbers if needed
			aNum, errA := strconv.ParseFloat(a, 64)
			bNum, errB := strconv.ParseFloat(b, 64)
			if errA == nil && errB == nil {
				return aNum < bNum
			}
		}

		// By default, sort by strings
		return a < b
	}

	// Sorting the lines

	// Removing duplicates if needed
	if *unique {
		lines = removeDuplicates(lines)
	}

	// Reversing the order if needed
	if *reverse {
		sort.SliceStable(lines, func(i, j int) bool {
			return !comparator(i, j)
		})
	} else {
		sort.SliceStable(lines, comparator)
	}

	// Creating the output file
	outputFile, err := os.Create("sorted_output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Ошибка записи строки в файл:", err)
			return
		}
	}
}

// Removes duplicates from an array of strings
func removeDuplicates(lines []string) []string {
	uniqueLines := make(map[string]bool)
	result := make([]string, 0)

	for _, line := range lines {
		if !uniqueLines[line] {
			uniqueLines[line] = true
			result = append(result, line)
		}
	}

	return result
}
