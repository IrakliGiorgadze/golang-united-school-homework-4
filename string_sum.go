package string_sum

import (
	"errors"
	"strconv"
	"unicode"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	runes := []rune(input)

	if len(runes) == 0 {
		return "", errorEmptyInput
	}

	for _, r := range runes {
		if !unicode.IsDigit(r) {
			return "", errorNotTwoOperands
		}
	}

	integers := make([]int, len(runes))
	for i, r := range runes {
		integers[i], _ = strconv.Atoi(string(r))
	}

	sum := 0
	for _, i := range integers {
		sum += i
	}

	return strconv.Itoa(sum), nil
}
