package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Params struct {
	System   string
	X, Y     int
	Operator rune
}

func main() {
	var params Params
	var result int

	description := "Arabic/Roman calculator.\n\n" +
		"Pass in two numbers in the range between 1 and 10\n" +
		"(in either Arabic or Roman system, but not together)\n" +
		"with one of the following operators in between: '-', '+', '*', or '/'.\n" +
		"All elements should be separated with a blank space.\n" +
		"Examples: 'a + b', 'a / b', 'IV * III', 'X - V'.\n" +
		"To stop the program press CTRL + C at the same time.\n\n"

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(description)
	for {
		fmt.Print("Input:\n")
		inputString, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		tokens, err := parseString(inputString)
		if err != nil {
			panic(err)
		}

		operand1, operator, operand2 := tokens[0], tokens[1], tokens[2]
		err = params.ValidateInput(operand1, operand2, operator)
		if err != nil {
			panic(err)
		}

		result = params.CalculateResult()

		if params.System == "arabic" {
			fmt.Printf("Output:\n%d\n\n", result)
		} else if params.System == "roman" {
			fmt.Printf("Output:\n%s\n\n", intToRoman(result))
		}
	}

}

func (p *Params) CalculateResult() int {
	result := 0
	switch p.Operator {
	case '+':
		result = p.X + p.Y
	case '-':
		result = p.X - p.Y
	case '*':
		result = p.X * p.Y
	case '/':
		result = p.X / p.Y
	}
	return result
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

func (p *Params) ValidateInput(operand1 string, operand2 string, operator string) error {
	var err error = nil

	romanToInt := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}

	if !isValidOperator(operator) {
		err = fmt.Errorf("operators allowed: +, -, *, /; got: %s", operator)
	} else {
		p.Operator = rune(operator[0])
	}

	if isValidRoman(operand1, romanToInt) && isValidRoman(operand2, romanToInt) {
		p.X = romanToInt[operand1]
		p.Y = romanToInt[operand2]
		p.System = "roman"
	} else if isValidInt(operand1) && isValidInt(operand2) {
		p.X, _ = strconv.Atoi(operand1)
		p.Y, _ = strconv.Atoi(operand2)
		p.System = "arabic"
	} else {
		err = fmt.Errorf("input numbers should be either both integers or roman, got: %s and %s", operand1, operand2)
	}

	return err

}

func intToRoman(numInt int) string {
	convertTable := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	if numInt < 1 {
		err := fmt.Errorf("roman system doesn't have zero and negative numbers, got: %d", numInt)
		panic(err)
	}

	numRoman := ""
	for _, pair := range convertTable {
		for numInt >= pair.value {
			numRoman += pair.digit
			numInt -= pair.value
		}
	}
	return numRoman
}
