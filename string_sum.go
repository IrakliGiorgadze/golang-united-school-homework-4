package string_sum

import (
	"errors"
	"strconv"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	//convert input to runes
	runes := []rune(input)

	//convert runes to array of characters
	chars := make([]string, len(runes))

	for i, r := range runes {
		chars[i] = string(r)
	}

	//use isDigit to check if character is a digit
	isDigit := func(c string) bool {
		_, err := strconv.Atoi(c)
		return err == nil
	}

	//use isOperator to check if character is an operator
	isOperator := func(c string) bool {
		return c == "+" || c == "-"
	}

	//use isSpace to check if character is a space
	isSpace := func(c string) bool {
		return c == " "
	}

	//calculate sum of digits
	sum := func(digits []string) (int, error) {
		var sum int
		for _, d := range digits {
			i, err := strconv.Atoi(d)
			if err != nil {
				return 0, err
			}
			sum += i
		}
		return sum, nil
	}

	//check if input is empty
	if len(chars) == 0 {
		return "", errorEmptyInput
	}

	//check if input contains only digits and operators
	for _, c := range chars {
		if !isDigit(c) && !isOperator(c) && !isSpace(c) {
			return "", errorNotTwoOperands
		}
	}

	return output, nil
}
