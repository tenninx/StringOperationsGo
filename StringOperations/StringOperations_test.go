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
