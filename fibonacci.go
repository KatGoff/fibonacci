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
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("Fibonacci: %d\n", FibonacciFinder(uint(*input)))
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
