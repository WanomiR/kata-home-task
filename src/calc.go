package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var intToRoman = map[int]string{
// 	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X", 100: "C"
// }

func main() {
	var x, y int
	var numSystem string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your input goes here:\n")
	inputString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	tokens, err := parseString(inputString)
	if err != nil {
		panic(err)
	}

	operand1, operator, operand2 := tokens[0], tokens[1], tokens[2]
	validateInput(operand1, operand2, &numSystem, &operator, &x, &y)

	fmt.Println(x, operator, y, numSystem)

}

func parseString(str string) ([]string, error) {
	str = strings.TrimSpace(str)
	tokens := strings.Split(str, " ")

	if len(tokens) != 3 {
		return nil, fmt.Errorf("expected 3 intput elements, got: %d", len(tokens))
	}

	return tokens, nil
}

func isValidRoman(operand string, romanToInt map[string]int) bool {
	for key, _ := range romanToInt {
		if operand == key {
			return true
		}
	}
	return false
}

func isValidInt(operand string) bool {
	num, err := strconv.Atoi(operand)
	if err == nil {
		if num < 1 || num > 10 {
			rangeErr := fmt.Errorf("input numbers should be from 1 to 10, got: %d", num)
			panic(rangeErr)
		}
		return true
	}
	return false
}

func isValidOperator(operator string) bool {
	validOperators := []string{"+", "-", "*", "/"}
	for _, value := range validOperators {
		if operator == value {
			return true
		}
	}
	return false
}

func validateInput(operand1 string, operand2 string, numSystem *string, operator *string, x *int, y *int) {
	romanToInt := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if !isValidOperator(*operator) {
		err := fmt.Errorf("operators allowed: +, -, *, /; got: %s", *operator)
		panic(err)
	}

	if isValidRoman(operand1, romanToInt) && isValidRoman(operand2, romanToInt) {
		*x = romanToInt[operand1]
		*y = romanToInt[operand2]
		*numSystem = "roman"
	} else if isValidInt(operand1) && isValidInt(operand2) {
		*x, _ = strconv.Atoi(operand1)
		*y, _ = strconv.Atoi(operand2)
		*numSystem = "arabic"
	} else {
		err := fmt.Errorf("input numbers should be either both integers or roman, got: %s and %s", operand1, operand2)
		panic(err)
	}

}
