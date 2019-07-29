package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// InputFromFile reads the contents of a specified file and returns it as an array
func InputFromFile(inputFile string) []int {
	var inputArray []int

	// input input.txt
	fileByte, err := ioutil.ReadFile(inputFile)
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

// InputFromParameter takes an integer from a flag
func InputFromParameter(inputNum int) int {
	if flagSet("inputInt") == false {
		fmt.Printf("Input: ")
		_, err := fmt.Scanf("%d", inputNum)
		if err != nil {
			errorHandler(err)
		}
	}
	return inputNum
}

// InputFromPrompt prompts for input and accepts an integer
func InputFromPrompt() int {
	var inputNum int
	fmt.Printf("Input: ")
	_, err := fmt.Scanf("%d", &inputNum)
	if err != nil {
		errorHandler(err)
	}
	return inputNum
}
