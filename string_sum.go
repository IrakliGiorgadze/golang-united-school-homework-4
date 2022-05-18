package string_sum

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

const errFormat = "err: %w"

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	errorInvalidInput   = errors.New("invalid input")
)

func StringSum(input string) (output string, err error) {
	in := []rune(strings.ReplaceAll(input, " ", ""))
	if len(in) == 0 {
		return "", fmt.Errorf(errFormat, errorEmptyInput)
	}

	var listNumbers []int
	var num []rune

	isDigitalBefore := false

	for i := 0; i < len(in); i++ {
		var first, AddToLists, isDigitalNow = false, false, false
		if len(num) == 0 {
			first = true
		}
		num = append(num, in[i])

		if unicode.IsDigit(in[i]) {
			isDigitalNow = true
		} else {
			isDigitalNow = false
		}

		if !first && isDigitalNow == false && isDigitalBefore || i == (len(in)-1) && unicode.IsDigit(in[i]) {
			AddToLists = true
		}

		if AddToLists {
			var number int
			var err error

			if i == len(in)-1 {
				number, err = strconv.Atoi(string(num))
			} else {
				number, err = strconv.Atoi(string(num[0 : len(num)-1]))
			}
			if err != nil {
				log.Println(err)
				return "", fmt.Errorf(errFormat, err)
			}

			listNumbers = append(listNumbers, number)
			num = nil
			num = append(num, in[i])
		}

		if unicode.IsDigit(in[i]) {
			isDigitalBefore = true
		} else {
			isDigitalBefore = false
		}
	}
	if len(listNumbers) != 2 {
		return "", fmt.Errorf(errFormat, errorNotTwoOperands)
	}

	return strconv.Itoa(listNumbers[0] + listNumbers[1]), nil
}
