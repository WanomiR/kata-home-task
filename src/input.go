package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input a value:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		toNumber, _ := strconv.Atoi(text)
		fmt.Println(toNumber + 8)
	}
}
