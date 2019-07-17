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
	input := flag.Int("input", 0, "Input integer to find Fibonacci number.")
	flag.Parse()
	if flagSet("input") == false {
		fmt.Printf("Input: ")
		_, err := fmt.Scanf("%d", input)
		if err != nil {
			Error(err)
		}
	}

	// Print result to output.txt
	result := fmt.Sprintf("Input: %d | Fibonacci: %d\n", *input, FibonacciFinder(uint(*input)))
	output, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Error(err)
	}
	if _, err = output.WriteString(result); err != nil {
		Error(err)
	}
	if err := output.Close(); err != nil {
		Error(err)
	}

	fmt.Printf(result)
}

// Error prints the message to the terminal and exits the program
func Error(err error) bool {
	fmt.Printf("Error: %v\n", err)
	os.Exit(1)
	return true
}

// FibonacciFinder finds Fib(n)
func FibonacciFinder(n uint) uint {
	if _, ok := Fib[n]; !ok {
		Fib[n] = FibonacciFinder(n-1) + FibonacciFinder(n-2)
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
