package main

import (
	"fmt"
	"strings"
)

func main() {
	const message1 = "Hello,"
	var message2 = "Welcome!"
	message3 := `I love Go`

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)

	name := "Wanomir"
	fmt.Printf("%c\n", name[0])

	nameLength := len(name)
	fmt.Println("Length of my name is:", nameLength)

	// concatenation
	concat := message1 + " " + message2
	fmt.Println(concat)

	// create three strings
	string1 := "Wanomir"
	string2 := "Wanomir Pro"
	string3 := "Wanomir"

	// compare strings
	fmt.Println(strings.Compare(string1, string2)) // -1
	fmt.Println(strings.Compare(string2, string3)) // 1
	fmt.Println(strings.Compare(string1, string3)) // 0

	text := "Go Programming"
	substring1 := "Go"
	substring2 := "Golang"

	// check if Go is present in Go Programming
	result := strings.Contains(text, substring1)
	fmt.Println(result)

	// check if Golang is present in Go Programming
	result = strings.Contains(text, substring2)
	fmt.Println(result)

	text1 := "car"
	fmt.Println("Old String:", text)

	// replace r with t
	replacedText := strings.Replace(text1, "r", "t", 1)
	fmt.Println("New String:", replacedText)

	// replace 2 r with 2 a
	strings.Replace("Programiz", "r", "R", 2)
	// Output: PRogRamiz

	var message = "I love Golang"
	// split string from space " "
	splittedString := strings.Split(message, " ")

	fmt.Println(splittedString)
}
