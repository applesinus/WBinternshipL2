package main

import "testing"

// given test 1
// - "a4bc2d5e" => "aaaabccddddde"
func Test1(t *testing.T) {
	str := unpack("a4bc2d5e")
	if str != "aaaabccddddde" {
		t.Fatalf("Test 1: a4bc2d5e\nExpected %s, got %s", "aaaabccddddde", str)
	}
}

// given test 2
// - "abcd" => "abcd"
func Test2(t *testing.T) {
	str := unpack("abcd")
	if str != "abcd" {
		t.Fatalf("Test 2: abcd\nExpected %s, got %s", "abcd", str)
	}
}

// given test 3
// - "45" => ""
func Test3(t *testing.T) {
	str := unpack("45")
	if str != "" {
		t.Fatalf("Test 3: 45\nExpected %s (empty string), got %s", "", str)
	}
}

// given test 4
// - "" => ""
func Test4(t *testing.T) {
	str := unpack("")
	if str != "" {
		t.Fatalf("Test 4: (empty string)\nExpected %s (empty string), got %s", "", str)
	}
}

// given test 5
// - "qwe\4\5" => "qwe45"
func Test5(t *testing.T) {
	str := unpack("qwe\\4\\5")
	if str != "qwe45" {
		t.Fatalf("Test 5: qwe\\4\\5\nExpected %s, got %s", "qwe45", str)
	}
}

// given test 6
// - "qwe\45" => "qwe44444"
func Test6(t *testing.T) {
	str := unpack("qwe\\45")
	if str != "qwe44444" {
		t.Fatalf("Test 6: qwe\\45\nExpected %s, got %s", "qwe44444", str)
	}
}

// given test 7
// - "qwe\5" => "qwe\\\\"
func Test7(t *testing.T) {
	str := unpack("qwe\\\\5")
	if str != "qwe\\\\\\\\\\" {
		t.Fatalf("Test 7: qwe\\\\5\nExpected %s, got %s", "qwe\\\\\\\\\\", str)
	}
}

// custom test to ensure that escaped 'a'-'z' are considered as a wrong expression
// - "a\\b" => ""
func Test8(t *testing.T) {
	str := unpack("a\\b")
	if str != "" {
		t.Fatalf("Test 8: a\\b\nExpected %s (empty string), got %s", "", str)
	}
}
