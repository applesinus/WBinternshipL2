package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// УТОЧНЕНИЕ НА СДАЧУ:
// в тз не содержится информации про обработку прописных букв
// тем не менее есть закомменченая строка с их обработкой

import (
	"strings"
)

// writeFinal writes character count times to given builder
func writeFinal(final *strings.Builder, c rune, count int) {
	final.Grow(count)
	for i := 0; i < count; i++ {
		final.WriteRune(c)
	}
}

func unpack(str string) string {
	// declaring variables
	var final strings.Builder
	prev := '-'
	escaped := false

	for _, c := range str {
		switch {

		// if a character is a back slash
		case c == '\\':
			if escaped {
				// if it's escaped, write it one time to final and set prev to it
				writeFinal(&final, c, 1)
				prev = c
			}
			escaped = !escaped

		// if a character is a digit
		case int(c) >= '0' && int(c) <= '9':
			if escaped {
				// if it's escaped, write it one time to final and set prev to it
				writeFinal(&final, c, 1)
				prev = c
				// also set escaped to false
				escaped = false
			} else {
				if prev == '-' {
					// if prev is an invalid character, return an empty string
					return ""
				}
				// if prev is a valid character, write it count-1 times to final
				// count-1 because one time it's already written
				writeFinal(&final, prev, int(c)-'0'-1)
				// set prev to - to indicate it's not a valid character to write or use again
				prev = '-'
			}

		// if a character is a letter
		// use below line to process ALL CHARACTERS:
		// case (int(c) >= 'a' && int(c) <= 'z') || (int(c) >= 'A' && int(c) <= 'Z'):
		case int(c) >= 'a' && int(c) <= 'z':
			if escaped {
				// if it's escaped, return an empty string
				return ""
			}
			// if prev is a valid character, write it one time to final  and set prev to it
			writeFinal(&final, c, 1)
			prev = c
		}
	}

	return final.String()
}
