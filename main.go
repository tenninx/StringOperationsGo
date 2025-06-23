package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"xyrisz.com/dev/packages/StringOperations/StringOperations"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var input, output, choice string

	for {
		fmt.Println("A collection of string operations to manipulate and analyze strings.")
		fmt.Println("Select a string operation to perform:")
		fmt.Println("1  - Find duplicated characters in a string")
		fmt.Println("2  - Get all unique characters in a string")
		fmt.Println("3  - Remove duplicated characters in a string")
		fmt.Println("4  - Reverse a string")
		fmt.Println("5  - Reverse a string in a sentence but maintain word order")
		fmt.Println("6  - Get word count")
		fmt.Println("7  - Check palindrome")
		fmt.Println("8  - Find character(s) with max occurrences")
		fmt.Println("9  - Get all possible substrings in a string")
		fmt.Println("10 - Get and capitalize the first letter of each word in a sentence")
		fmt.Println("11 - Get a list of words matching the given first letter")
		fmt.Println("12 - Get a list of words matching the given continuous character(s) anywhere")
		fmt.Println("13 - Get a list of words matching the given character(s) in order")
		fmt.Println("X  - Exit")

		scanner.Scan()
		choice = strings.ToLower(scanner.Text())

		if choice == "x" {
			break
		}

		fmt.Println("Enter the string to process:")

		scanner.Scan()
		input = scanner.Text()

		switch choice {
		case "1":
			output = StringOperations.FindDuplicatedChars(input, false)
		case "2":
			output = StringOperations.FindUniqueChars(input, false)
		case "3":
			output = StringOperations.RemoveDuplicatedChars(input, false)
		case "4":
			output = StringOperations.ReverseString(input, true)
		case "5":
			output = StringOperations.ReverseStringInSentence(input, true)
		case "6":
			output = StringOperations.GetWordCount(input, false)
		case "7":
			output = StringOperations.IsPalindrome(input, false)
		case "8":
			output = StringOperations.FindMaxOccurrences(input, false)
		case "9":
			output = StringOperations.GetAllSubstrings(input, false)
		case "10":
			output = StringOperations.GetFirstLetterCapital(input, false)
		case "11":
			output = StringOperations.GetWordsMatchPrefix(input, false)
		case "12":
			output = StringOperations.GetWordsMatchAnywhere(input, false)
		case "13":
			output = StringOperations.GetWordsMatchAnywhereRandom(input, false)
		}

		fmt.Println("The output is: " + output)
	}
}
