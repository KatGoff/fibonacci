package main

import (
	"flag"
	"fmt"
	"os"
)

// Fib allows the stored map keys and values to be accessed in FibonacciFinder
var Fib = map[uint]uint{
	0: 0,
	1: 1,
}

// Results stores all inputs and results from FibonacciFinder
var Results = map[uint]uint{}

func main() {
	var inputArray []int

	// Flags
	inputNum := flag.Int("inputInt", 0, "Input integer to find Fibonacci number.")
	inputFile := flag.String("inputFile", "input.txt", "Input filename to find Fibonacci numbers.")
	outputFile := flag.String("outputFile", "output.txt", "Input filename to print Fibonacci numbers to.")
	flag.Parse()

	if flagSet("inputFile") == true {
		inputArray = InputFromFile(*inputFile)
		calculateFibonacci(inputArray)
		Output(Results, *outputFile)
	} else if flagSet("inputNum") == true {
		inputArray = []int{InputFromParameter(*inputNum)}
		calculateFibonacci(inputArray)
		Output(Results, *outputFile)
	} else {
		inputArray = []int{InputFromPrompt()}
		calculateFibonacci(inputArray)
		Output(Results, *outputFile)
	}
}

// FibFinder finds Fib(n)
func FibFinder(n uint) uint {
	if _, ok := Fib[n]; !ok {
		Fib[n] = FibFinder(n-1) + FibFinder(n-2)
	}
	return Fib[n]
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

func calculateFibonacci(inputArray []int) {
	for _, num := range inputArray {
		Results[uint(num)] = FibFinder(uint(num))
	}
}

// errorHandler prints the message to the terminal and exits the program
func errorHandler(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}
