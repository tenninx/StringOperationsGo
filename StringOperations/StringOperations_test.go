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
