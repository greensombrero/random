package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to non repeating random number generator")
	fmt.Println("Please enter a range of numbers to pull from")
	reader := bufio.NewReader(os.Stdin)
	start := readNumber(reader, "First number")
	end := readNumber(reader, "Last number")

	generator := NewGenerator(start, end)
	for true {
		value, err := generator.Next()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Println(value)
		reader.ReadString('\n')
	}
}

func readNumber(reader *bufio.Reader, prompt string) int64 {
	valid := false
	for !valid {
		fmt.Printf("%v: ", prompt)

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		// convert CRLF to LF
		line = strings.Replace(line, "\n", "", -1)
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Println("Invalid input")
		} else {
			return value
		}
	}

	return -1
}
