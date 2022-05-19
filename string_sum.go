package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}
func (e MyError) Unwrap() error {
	return fmt.Errorf("error")
}

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

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
	sum := int64(0)
	sumCnt := 0
	if len(strings.TrimSpace(input)) < 1 {
		return "", fmt.Errorf("undef: %w", errorEmptyInput)
	}
	for _, sumItem := range strings.Split(input, "+") {
		for i, item := range strings.Split(sumItem, "-") {
			trim := strings.TrimSpace(item)
			value, err := strconv.ParseInt(trim, 10, 64)
			if err != nil {
				if len(trim) < 1 {
					continue
				}
				var err2 MyError = MyError{Message: err.Error()}
				return "", fmt.Errorf("undef: %w", err2)
			}
			sumCnt += 1
			if i > 0 {
				sum -= value
			} else {
				sum += value
			}
		}
	}
	if sumCnt != 2 {
		return "", fmt.Errorf("undef: %w", errorNotTwoOperands)
	}
	return strconv.FormatInt(sum, 10), nil
}
