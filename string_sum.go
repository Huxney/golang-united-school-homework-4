package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when the input contains letters
	errorLetterInString = errors.New("Input contains letters")
)

func ReturnNumbersInString(input string) []string {
	re := regexp.MustCompile(`-?\d[\d,]*\.?[\d{2}]*`)
	numbers := re.FindAllString(input, -1)
	return numbers
}

func ConvertStringArrayToIntArray(param []string) []int {
	var arrayAsIntegers []int

	for _, i := range param {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arrayAsIntegers = append(arrayAsIntegers, j)
	}

	return arrayAsIntegers
}

func CalculateSumOfInts(array []int) int {
	res := 0
	for i := 0; i < len(array); i++ {
		res += array[i]
	}
	return res
}

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	if len(input) == 0 {
		return "", fmt.Errorf("input is empty: %w", errorEmptyInput)
	}

	input = strings.ReplaceAll(input, " ", "")

	var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	if isStringAlphabetic(input) == true {
		return "", fmt.Errorf("%w", errorLetterInString)
	}

	arrayOfStrings := ReturnNumbersInString(input)
	arrayOfNumbers := ConvertStringArrayToIntArray(arrayOfStrings)

	if len(arrayOfNumbers) != 2 {
		return "", fmt.Errorf("too many operands %d: %w", len(arrayOfNumbers), errorNotTwoOperands)
	}

	if isStringAlphabetic(input) == true {
		return "", fmt.Errorf("%w", errorLetterInString)
	}

	sum := CalculateSumOfInts(arrayOfNumbers)
	return strconv.Itoa(sum), err
}
