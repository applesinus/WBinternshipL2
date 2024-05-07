package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testData := map[*[]string]map[string][]string{
		&[]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}: {
			"акптя":  {"пятак", "пятка", "тяпка"},
			"иклост": {"листок", "слиток", "столик"},
		},
		&[]string{"я", "хочу", "работать", "в", "WB"}: {},
		&[]string{}:               {},
		&[]string{"мама", "амма"}: {"аамм": {"мама", "амма"}},
	}

	for input, expected := range testData {
		result := FindAnagrams(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	}
}
