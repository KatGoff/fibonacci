package main

import (
	"fmt"
	"os"
)

// Output automatically calls the outputToTerminal function, but checks that an output file has been specified before calling the outputToFile function
func Output(input, fib uint, outputFile string) {
	outputToTerminal(input, fib)
	if flagSet("outputFile") == true {
		outputToFile(input, fib, outputFile)
	}
}

func outputToTerminal(input, fib uint) {
	fmt.Printf("Input: %d | Fibonacci: %d\n", input, fib)
}

func outputToFile(input, fib uint, outputFile string) {
	output, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	message := fmt.Sprintf("Input: %d | Fibonacci: %d\n", input, fib)
	if err != nil {
		errorHandler(err)
	}
	if _, err = output.WriteString(message); err != nil {
		errorHandler(err)
	}
	if err := output.Close(); err != nil {
		errorHandler(err)
	}
}
