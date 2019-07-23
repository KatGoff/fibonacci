package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Fib allows the stored map keys and values to be accessed in FibonacciFinder
var Fib = map[uint]uint{
	0: 0,
	1: 1,
}

func main() {
	var inputArray []int

	// if no input.txt file exists, take input as a parameter
	if _, err := os.Stat("input.txt"); err != nil {
		inputArray = []int{inputParameter()}
	} else {
		inputArray = inputFile()
	}

	outputResult(inputArray)
}

// FibonacciFinder finds Fib(n)
func FibonacciFinder(n uint) uint {
	if _, ok := Fib[n]; !ok {
		Fib[n] = FibonacciFinder(n-1) + FibonacciFinder(n-2)
	}
	return Fib[n]
}

// rename function e.g. readIntegersFromInputFile(fileName string)
func inputFile() []int {
	var inputArray []int

	// can we change this so that you pass in the file as a parameter and if there isn't one, you prompt?
	// input input.txt
	fileByte, err := ioutil.ReadFile("input.txt")
	if err != nil {
		errorHandler(err)
	}

	// convert input and append to inputArray
	fileStr := strings.Split(string(fileByte), "\n")
	for _, str := range fileStr {
		integer, err := strconv.Atoi(str)
		inputArray = append(inputArray, integer)
		if err != nil {
			errorHandler(err)
		}
	}

	return inputArray
}

func inputParameter() int {
	inputNum := flag.Int("input", 0, "Input integer to find Fibonacci number.")
	flag.Parse()
	if flagSet("input") == false {
		fmt.Printf("Input: ")
		_, err := fmt.Scanf("%d", inputNum)
		if err != nil {
			errorHandler(err)
		}
	}
	return *inputNum
}

// flagSet checks that a flag has been provided
func flagSet(name string) bool {
	var flagFound bool

	flag.Visit(func(flag *flag.Flag) {
		if flag.Name == name {
			flagFound = true
		}
	})

	return flagFound
}

func outputFile(result string) string {
	output, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		errorHandler(err)
	}
	if _, err = output.WriteString(result); err != nil {
		errorHandler(err)
	}
	if err := output.Close(); err != nil {
		errorHandler(err)
	}
	return result
}

func outputResult(inputArray []int) {
	for _, num := range inputArray {
		result := fmt.Sprintf("Input: %d | Fibonacci: %d\n", num, FibonacciFinder(uint(num)))
		fmt.Printf(outputFile(result))
	}
}

func calculateFibonacci(numbers ...int) map[int]int {
	// calculate the numbers and then return the map
}
// if output file is specified as a parameter:
func printResultsToFile(numToFibMap map[int]int) error {
	...
}
// in all cases:
func printResultsToScreen(numToFibMap map[int]int) {
	...
}

// errorHandler prints the message to the terminal and exits the program
func errorHandler(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}
