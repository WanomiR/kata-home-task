package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var romanToInt = map[string]int{
// 	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
// }

// var intToRoman = map[int]string{
// 	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
// }

func main() {
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

	var x, operand, y interface{} = tokens[0], tokens[1], tokens[2]

	fmt.Println(x, operand, y)
	// fmt.Println(romanToInt)
	// fmt.Println(intToRoman)

}

func parseString(str string) ([]string, error) {
	str = strings.TrimSpace(str)
	tokens := strings.Split(str, " ")

	if len(tokens) != 3 {
		return nil, fmt.Errorf("expected 3 intput elements, got %d", len(tokens))
	}

	return tokens, nil
}
