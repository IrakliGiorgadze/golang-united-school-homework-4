package string_sum

import (
	"errors"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string, input2 string) (output string, err error) {
	//add 2 strings together and return the result as a string  if the input is empty, return an error
	if input == "" {
		err = errorEmptyInput
		return
	}
	if input2 == "" {
		err = errorEmptyInput
		return
	}
	if len(input) != len(input2) {
		err = errorNotTwoOperands
		return
	}
	for i := 0; i < len(input); i++ {
		output += string(input[i] + input2[i])
	}

	return "", nil
}
