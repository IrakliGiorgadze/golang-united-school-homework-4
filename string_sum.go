package string_sum

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)



func StringSum(input string) (output string, err error) {
	if input == "" {
		return "", errorEmptyInput
	}

	// if len(input)%2 != 0 {
	// 	return "", errorNotTwoOperands
	// }

	if input[0] == '-' {
		return "", errorNotTwoOperands
	}

	if input[1] == ' ' {
		return "", errorNotTwoOperands
	}

	runes := []rune(input)

	fmt.Println("runes:", runes)

	chars := make([]string, len(runes))

	fmt.Println("chars:", chars)

	for i, r := range runes {
		chars[i] = string(r)
	}

	fmt.Println("chars:", chars)

	//add the first two numbers 
	firstNumber, err := strconv.Atoi(chars[0])
	if err != nil {
		return "", err
	}
	secondNumber, err := strconv.Atoi(chars[2])
	if err != nil {
		return "", err
	}

	sum := firstNumber + secondNumber
	fmt.Println("sum:", sum)

	//return the sum
	return "", nil

}

