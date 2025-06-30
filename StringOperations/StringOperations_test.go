package StringOperations

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestFindDuplicatedChars(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
	}{
		{"hello", false, "l"},
		{"Hello", false, "l"},
		{"", false, ""},
		{"abcde", false, ""},
		{"aabbcc", false, "abc"},
	}

	for _, test := range tests {
		result := FindDuplicatedChars(test.input, test.isCaseSensitive)
		if result != test.expected {
			t.Errorf("FindDuplicatedChars(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func TestFindUniqueChars(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
	}{
		{"hello", false, "heo"},
		{"Google", false, "le"},
		{"", false, ""},
		{"abcde", false, "abcde"},
		{"aabbcc", false, ""},
	}

	for _, test := range tests {
		result := FindUniqueChars(test.input, test.isCaseSensitive)
		if result != test.expected {
			t.Errorf("FindUniqueChars(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func TestRemoveDuplicatedChars(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
	}{
		{"hello", false, "helo"},
		{"Google", false, "gole"},
		{"", false, ""},
		{"abcde", false, "abcde"},
		{"aabbcc", false, "abc"},
	}

	for _, test := range tests {
		result := RemoveDuplicatedChars(test.input, test.isCaseSensitive)
		if result != test.expected {
			t.Errorf("RemoveDuplicatedChars(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
	}{
		{"hello", true, "olleh"},
		{"Hello", true, "olleH"},
		{"", true, ""},
		{"abcde", true, "edcba"},
	}

	for _, test := range tests {
		result := ReverseString(test.input, test.isCaseSensitive)
		if result != test.expected {
			t.Errorf("ReverseString(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func TestReverseStringInSentence(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
	}{
		{"hello world", true, "olleh dlrow"},
		{"Hello World", true, "olleH dlroW"},
		{"", true, ""},
		{"abc def ghi", true, "cba fed ihg"},
	}

	for _, test := range tests {
		result := ReverseStringInSentence(test.input, test.isCaseSensitive)
		if result != test.expected {
			t.Errorf("ReverseStringInSentence(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func TestGetWordCount(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
	}{
		{"hello world", true, "2"},
		{"Hello World", true, "2"},
		{"", true, "0"},
		{"abc def ghi", true, "3"},
	}

	for _, test := range tests {
		result := GetWordCount(test.input, test.isCaseSensitive)
		if result != test.expected {
			t.Errorf("GetWordCount(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
	}{
		{"hello world", false, "false"},
		{"aBCbA", false, "true"},
		{"", false, "true"},
		{"nurses Run", false, "true"},
	}

	for _, test := range tests {
		result := IsPalindrome(test.input, test.isCaseSensitive)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func TestFindMaxOccurrences(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
		count           int
	}{
		{"hello world", false, "l", 3},
		{"abcde", false, "a b c d e", 1},
		{"", false, "", 0},
		{"aaabb ccc dd", false, "a c", 3},
	}

	for _, test := range tests {
		result := FindMaxOccurrences(test.input, test.isCaseSensitive)
		if !CheckUnorderedResultWithOccurence(result, test.expected, test.count) {
			t.Errorf("FindMaxOccurrences(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func TestGetAllSubstrings(t *testing.T) {
	tests := []struct {
		input           string
		isCaseSensitive bool
		expected        string
	}{
		{"india", false, "india indi ind in i ndia ndi nd n dia di d ia a"},
		{"", false, ""},
	}

	for _, test := range tests {
		result := GetAllSubstrings(test.input, test.isCaseSensitive)
		if !CheckUnorderedResult(result, test.expected) {
			t.Errorf("GetAllSubstrings(%q, %v) = %q; want %q", test.input, test.isCaseSensitive, result, test.expected)
		}
	}
}

func CheckUnorderedResult(result string, expected string) bool {
	var resultChars []string = strings.Split(result, "\n")
	var expectedChars []string = strings.Split(expected, " ")

	if result == "" && expected == "" {
		return true
	}

	if len(resultChars)-1 != len(expectedChars) {
		return false
	}

	for _, char := range expectedChars {
		if !slices.Contains(resultChars, char) {
			return false
		}
	}

	return true
}

func CheckUnorderedResultWithOccurence(result string, expected string, occurence int) bool {
	var resultChars []string = strings.Split(result, " ")
	var expectedChars []string = strings.Split(expected, " ")

	if len(resultChars) == 1 && resultChars[0] == "(0)" && occurence == 0 {
		return true
	}

	if !slices.Contains(resultChars, "("+strconv.Itoa(occurence)+")") {
		return false
	}

	if len(resultChars)-1 != len(expectedChars) {
		return false
	}

	for _, char := range expectedChars {
		if !slices.Contains(resultChars, char) {
			return false
		}
	}

	return true
}
