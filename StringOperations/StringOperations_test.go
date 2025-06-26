package StringOperations

import "testing"

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
