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

func main() {

	// if input.txt != ok {
	input := inputParameter()
	// } else inputFile (loop over an array of ints)

	result := fmt.Sprintf("Input: %d | Fibonacci: %d\n", *input, FibonacciFinder(uint(*input)))
	fmt.Printf(outputFile(result))
}

// FibonacciFinder finds Fib(n)
func FibonacciFinder(n uint) uint {
	if _, ok := Fib[n]; !ok {
		Fib[n] = FibonacciFinder(n-1) + FibonacciFinder(n-2)
	}
	return Fib[n]
}

func inputParameter() *int {
	inputNum := flag.Int("input", 0, "Input integer to find Fibonacci number.")
	flag.Parse()
	if flagSet("input") == false {
		fmt.Printf("Input: ")
		_, err := fmt.Scanf("%d", inputNum)
		if err != nil {
			errorHandler(err)
		}
	}
	return inputNum
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

// func inputFile() {
// read each line in input.txt
// check for int
// put int through FibonacciFinder
// }

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

// errorHandler prints the message to the terminal and exits the program
func errorHandler(err error) {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
}
