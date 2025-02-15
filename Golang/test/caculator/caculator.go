package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Printf("%v must be a number only.\n", prompt)
		os.Exit(1)
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	if value2 == 0 {
		fmt.Println("Error: Division by zero is not allowed.")
		os.Exit(1)
	}
	return value1 / value2
}

func getOperator() string {
	fmt.Println("Operator is (+, -, *, /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput("Value1 = ")
	value2 := getInput("Value2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /.")
		os.Exit(1)
	}

	fmt.Printf("Result is: %v\n", result)
}
