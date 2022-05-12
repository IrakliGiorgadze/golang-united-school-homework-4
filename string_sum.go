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
	var sum int
	for i := 0; i < len(input); i += 2 {
		sum += int(input[i]) + int(input[i+1])
	}
	return string(sum), nil
}
