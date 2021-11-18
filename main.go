package main

import (
	"fmt"
	"os"
	"strings"
)

var x, y int
var input string

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func main() {
	fmt.Println("Calculator: Add, Subtract, Multiply, Divide")
	fmt.Println("Enter the numbers you would like to calculate: ")

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	fmt.Print("Enter the first letter of the operand to perform: ")

	fmt.Scan(&input)

	if strings.ToLower(input) == "a" {
		fmt.Println(add(x, y))
	} else if strings.ToLower(input) == "s" {
		fmt.Println(subtract(x, y))
	} else if strings.ToLower(input) == "m" {
		fmt.Println(multiply(x, y))
	} else if strings.ToLower(input) == "d" {
		fmt.Println(divide(x, y))
	} else if strings.ToLower(input) == "exit" {
		os.Exit(3)
	}
}
