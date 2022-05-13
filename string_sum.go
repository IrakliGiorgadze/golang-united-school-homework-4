package string_sum

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// func TestStringSum(t *testing.T) {
// 	testCases := map[string]struct {
// 		input    string
// 		output   string
// 		expErr   error
// 		numError bool
// 	}{
// 		"both operands positive":    {input: "24+55", output: "79", expErr: nil},
// 		"first operand negative":    {input: "-24+55", output: "31", expErr: nil},
// 		"second operand negative":   {input: "24-55", output: "-31", expErr: nil},
// 		"both operands negative":    {input: "-24-55", output: "-79", expErr: nil},
// 		"with whitespace":           {input: " -24 - 55 ", output: "-79", expErr: nil},
// 		"empty input":               {input: "", output: "", expErr: errorEmptyInput},
// 		"three operands":            {input: "11+23+43", output: "", expErr: errorNotTwoOperands},
// 		"one operand":               {input: "42", output: "", expErr: errorNotTwoOperands},
// 		"letters in first operand":  {input: "24c+55", output: "", expErr: &strconv.NumError{Func: "Atoi", Num: "24c", Err: strconv.ErrSyntax}, numError: true},
// 		"letters in second operand": {input: "24+55f", output: "", expErr: &strconv.NumError{Func: "Atoi", Num: "55f", Err: strconv.ErrSyntax}, numError: true},
// 	}

// 	for name, tt := range testCases {
// 		t.Run(name, func(t *testing.T) {
// 			output, err := StringSum(tt.input)
// 			if tt.expErr != nil {
// 				if tt.numError {
// 					e := errors.Unwrap(err)
// 					if !reflect.DeepEqual(e, tt.expErr) {
// 						t.Errorf("%s:\n wrong error is wrapped into the returned error: got %s, want %s", name, e, tt.expErr)
// 					}
// 					if !errors.As(err, &tt.expErr) {
// 						t.Errorf("%s:\n wrong error type is used in the return: got %T, want %T", name, err, tt.expErr)
// 					}

// 				} else {
// 					if err == tt.expErr {
// 						t.Errorf("%s:\n returned error must be wrapped", name)
// 					}
// 					if !errors.Is(err, tt.expErr) {
// 						t.Errorf("%s:\n wrong error is used in the return: got %s, want %s", name, err.Error(), tt.expErr.Error())
// 					}
// 				}
// 			}
// 			if err != nil {
// 				t.Errorf("error should be nil: got %s", err)
// 			}

// 			if output != tt.output {
// 				t.Errorf("error in the sum output: got %s, want %s", output, tt.output)
// 			}
// 		})
// 	}
// }

func StringSum(input string) (output string, err error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errorEmptyInput
	}
	operands := strings.Split(input, "+")
	if len(operands) != 2 {
		return "", errorNotTwoOperands
	}
	first, err := strconv.Atoi(operands[0])
	if err != nil {
		return "", err
	}
	second, err := strconv.Atoi(operands[1])
	if err != nil {
		return "", err
	}

	sum := first + second
	//convert sum to string
	output = strconv.Itoa(sum)

	return output, nil

}
