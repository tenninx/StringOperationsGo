package StringOperations

import (
	"bytes"
	"slices"
	"strconv"
	"strings"
)

var toMatch = [...]string{"hello", "world", "hi", "my", "home", "we", "are", "the", "champion", "c#", "is", "a", "great", "languague", "champ", "campaign", "champagne", "challenge", "chameleon", "chipotle", "castlevania"}

func FindDuplicatedChars(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)
	var dupChars []string

	for i := 0; i < len(p_strInput); i++ {
		substring := p_strInput[i+1:]

		if strings.Contains(substring, string(p_strInput[i])) {
			if !slices.Contains(dupChars, string(p_strInput[i])) {
				dupChars = append(dupChars, string(p_strInput[i]))
			}
		}
	}

	return strings.Join(dupChars, "")
}

func FindUniqueChars(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)
	var uniqueChars []string

	for i := 0; i < len(p_strInput); i++ {
		if strings.Count(p_strInput, string(p_strInput[i])) == 1 {
			uniqueChars = append(uniqueChars, string(p_strInput[i]))
		}
	}

	return strings.Join(uniqueChars, "")
}

func RemoveDuplicatedChars(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)
	var uniqueChars []string

	for i := 0; i < len(p_strInput); i++ {
		if !slices.Contains(uniqueChars, string(p_strInput[i])) {
			uniqueChars = append(uniqueChars, string(p_strInput[i]))
		}
	}

	return strings.Join(uniqueChars, "")
}

func ReverseString(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)
	var reversed bytes.Buffer

	for i := len(p_strInput) - 1; i >= 0; i-- {
		reversed.WriteString(string(p_strInput[i]))
	}

	return reversed.String()
}

func ReverseStringInSentence(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)
	var reversed bytes.Buffer

	words := strings.Fields(p_strInput)

	for idx, element := range words {
		for i := len(element) - 1; i >= 0; i-- {
			reversed.WriteString(string(element[i]))
		}

		if idx < len(words)-1 {
			reversed.WriteString(" ")
		}
	}

	return reversed.String()
}

func GetWordCount(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)

	return strconv.Itoa(len(strings.Fields(p_strInput)))
}

func IsPalindrome(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)

	p_strInput = strings.Replace(p_strInput, " ", "", -1)

	for i := 0; i < len(p_strInput)/2; i++ {
		if string(p_strInput[i]) != string(p_strInput[len(p_strInput)-1-i]) {
			return "false"
		}
	}

	return "true"
}

func FindMaxOccurrences(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)
	var chars map[string]int = make(map[string]int)

	max := 0
	for i := 0; i < len(p_strInput); i++ {
		chars[string(p_strInput[i])] = chars[string(p_strInput[i])] + 1

		if chars[string(p_strInput[i])] > max {
			max = chars[string(p_strInput[i])]
		}
	}

	var result strings.Builder

	for key, value := range chars {
		if value == max {
			result.WriteString(key + " ")
		}
	}

	return result.String() + "(" + strconv.Itoa(max) + ")"
}

func GetAllSubstrings(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)
	var substrings map[string]bool = make(map[string]bool)

	for i := 0; i < len(p_strInput); i++ {
		for j := 0; j < len(p_strInput)-i; j++ {
			substring := p_strInput[i : i+j+1]

			_, ok := substrings[substring]
			if !ok {
				substrings[substring] = false
			}
		}
	}

	var result bytes.Buffer

	for key := range substrings {
		result.WriteString("\n" + key)
	}

	return result.String()
}

func GetFirstLetterCapital(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = ProcessInput(p_strInput, p_isCaseSensitive)

	var words []string = strings.Fields(p_strInput)

	var result bytes.Buffer

	for _, word := range words {
		if len(word) > 0 {
			result.WriteString(strings.ToUpper(word[:1]) + " ")
		}
	}

	return result.String()
}

func GetWordsMatchPrefix(p_prefix string, p_isCaseSensitive bool) string {
	p_prefix = ProcessInput(p_prefix, p_isCaseSensitive)

	var result bytes.Buffer

	for _, word := range toMatch {
		if strings.HasPrefix(word, p_prefix) {
			result.WriteString(word + " ")
		}
	}

	return result.String()
}

func GetWordsMatchAnywhere(p_anyLetters string, p_isCaseSensitive bool) string {
	p_anyLetters = ProcessInput(p_anyLetters, p_isCaseSensitive)

	var result bytes.Buffer

	for _, word := range toMatch {
		if strings.Contains(word, p_anyLetters) {
			result.WriteString(word + " ")
		}
	}

	return result.String()
}

func GetWordsMatchAnywhereRandom(p_anyLetters string, p_isCaseSensitive bool) string {
	p_anyLetters = ProcessInput(p_anyLetters, p_isCaseSensitive)

	var result bytes.Buffer

	for _, word := range toMatch {
		var currentIdx = -1
		for i := 0; i < len(p_anyLetters); i++ {
			var idx = strings.Index(word, string(p_anyLetters[i]))
			if idx == -1 || idx < currentIdx {
				break
			}

			if i+1 < len(p_anyLetters) {
				currentIdx = idx
				continue
			}

			result.WriteString(word + " ")
		}
	}

	return result.String()
}

func ProcessInput(p_strInput string, p_isCaseSensitive bool) string {
	p_strInput = strings.TrimSpace(p_strInput)
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
