package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	fmt.Println(tokens)

}

func parseString(str string) ([]string, error) {
	str = strings.TrimSpace(str)
	tokens := strings.Split(str, " ")

	if len(tokens) != 3 {
		return nil, fmt.Errorf("Expected 3 intput elements, got %d", len(tokens))
	}

	return tokens, nil
}
