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
	runes := []rune(input)

	chars := make([]string, len(runes))

	for i, r := range runes {
		chars[i] = string(r)
	}

	sum := 0
	for _, c := range chars {
		num, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		sum += num
	}

	return strconv.Itoa(sum), nil

}
