package main

import (
	"fmt"
	"os"
)

// Output automatically calls the outputToTerminal function, but checks that an output file has been specified before calling the outputToFile function
func Output(Results map[uint]uint, outputFile string) {
	outputToTerminal(Results)
	if flagSet("outputFile") == true {
		outputToFile(Results, outputFile)
	}
}

func outputToTerminal(Results map[uint]uint) {
	for input, output := range Results {
		fmt.Printf("Input: %d | Fibonacci: %d\n", input, output)
	}
}

func outputToFile(Results map[uint]uint, outputFile string) {
	output, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	for integer, fib := range Results {
		message := fmt.Sprintf("Input: %d | Fibonacci: %d\n", integer, fib)
		if err != nil {
			errorHandler(err)
		}
		if _, err = output.WriteString(message); err != nil {
			errorHandler(err)
		}
	}
	if err := output.Close(); err != nil {
		errorHandler(err)
	}
}
