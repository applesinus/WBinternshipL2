package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"sort"
	"strings"
)

// FindAnagrams is a function for finding anagrams
func FindAnagrams(words *[]string) map[string][]string {
	anagramSets := make(map[string][]string)

	// iterate over words
	for _, word := range *words {
		// lowering a word
		word = strings.ToLower(word)
		// sorting a word to get the first anagram
		wordRunes := []rune(word)
		sort.Slice(wordRunes, func(i, j int) bool { return wordRunes[i] < wordRunes[j] })
		sortedWord := string(wordRunes)

		// appending a new anagram
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	// getting rid of single object anagrams
	for key, value := range anagramSets {
		if len(value) == 1 {
			delete(anagramSets, key)
		}
	}

	return anagramSets
}
