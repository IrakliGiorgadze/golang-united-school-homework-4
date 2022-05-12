package string_sum

import (
	"errors"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	if input == "" {
		return "", errorEmptyInput
	}

	if len(input)%2 != 0 {
		return "", errorNotTwoOperands
	}

	output = ""
	for i := 0; i < len(input); i += 2 {
		output += input[i : i+2]
	}

	return output, nil

}
