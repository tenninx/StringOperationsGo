package StringOperations

import (
	"slices"
	"strings"
)

func FindDuplicatedChars(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = CheckCaseSensitivity(p_strInput, p_isCaseSensitive)
	var chars map[string]int = make(map[string]int)

	for i := 0; i < len(p_strInput); i++ {
		substring := p_strInput[i+1:]
		if slices.Contains([]rune(substring), rune(p_strInput[i])) {
			_, ok := chars[string(p_strInput[i])]
			if !ok {
				chars[string(p_strInput[i])] = 0
			}
		}
	}

	return ConvertStringMapToString(chars)
}

func FindUniqueChars(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = CheckCaseSensitivity(p_strInput, p_isCaseSensitive)
	var chars map[string]int = make(map[string]int)

	for i := 0; i < len(p_strInput); i++ {
		chars[string(p_strInput[i])] = chars[string(p_strInput[i])] + 1
	}

	for r, count := range chars {
		if count != 1 {
			delete(chars, r)
		}
	}

	return ConvertStringMapToString(chars)
}

func CheckCaseSensitivity(p_strInput string, p_isCaseSensitive bool) string {
	if !p_isCaseSensitive {
		return strings.ToLower(p_strInput)
	}

	return p_strInput
}

func ConvertStringMapToString(p_strInput map[string]int) string {
	var dupChars string
	for r := range p_strInput {
		dupChars += r
	}

	return dupChars
}
