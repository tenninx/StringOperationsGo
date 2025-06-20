package StringOperations

import (
	"bytes"
	"slices"
	"strconv"
	"strings"
)

func FindDuplicatedChars(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessCaseSensitivity(p_strInput, p_isCaseSensitive)
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
	p_strInput = ProcessCaseSensitivity(p_strInput, p_isCaseSensitive)
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

func RemoveDuplicatedChars(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessCaseSensitivity(p_strInput, p_isCaseSensitive)
	var chars map[string]int = make(map[string]int)

	for i := 0; i < len(p_strInput); i++ {
		_, ok := chars[string(p_strInput[i])]
		if !ok {
			chars[string(p_strInput[i])] = 0
		}
	}

	return ConvertStringMapToString(chars)
}

func ReverseString(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessCaseSensitivity(p_strInput, p_isCaseSensitive)
	var reversed string

	for i := len(p_strInput) - 1; i >= 0; i-- {
		reversed += string(p_strInput[i])
	}

	return reversed
}

func ReverseStringInSentence(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessCaseSensitivity(p_strInput, p_isCaseSensitive)
	var reversed bytes.Buffer

	words := strings.Split(p_strInput, " ")

	for _, element := range words {
		for i := len(element) - 1; i >= 0; i-- {
			reversed.WriteString(string(element[i]))
		}

		reversed.WriteString(" ")
	}

	if reversed.Len() > 0 {
		reversed.Truncate(reversed.Len() - 1)
	}

	return reversed.String()
}

func GetWordCount(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessCaseSensitivity(p_strInput, p_isCaseSensitive)

	p_strInput = strings.TrimSpace(p_strInput)

	return strconv.Itoa(len(strings.Split(p_strInput, " ")))
}

func IsPalindrome(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessCaseSensitivity(p_strInput, p_isCaseSensitive)

	p_strInput = strings.TrimSpace(p_strInput)
	p_strInput = strings.Replace(p_strInput, " ", "", -1)

	for i := 0; i < len(p_strInput)/2; i++ {
		if string(p_strInput[i]) != string(p_strInput[len(p_strInput)-1-i]) {
			return "false"
		}
	}

	return "true"
}

func FindMaxOccurrences(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessCaseSensitivity(p_strInput, p_isCaseSensitive)
	var chars map[string]int = make(map[string]int)

	max := 0
	for i := 0; i < len(p_strInput); i++ {
		chars[string(p_strInput[i])] = chars[string(p_strInput[i])] + 1

		if chars[string(p_strInput[i])] > max {
			max = chars[string(p_strInput[i])]
		}
	}

	var result bytes.Buffer

	for key, value := range chars {
		if value == max {
			result.WriteString(key + " ")
		}
	}

	if result.Len() > 0 {
		result.Truncate(result.Len() - 1)
	}

	return result.String()
}

func ProcessCaseSensitivity(p_strInput string, p_isCaseSensitive bool) string {
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
