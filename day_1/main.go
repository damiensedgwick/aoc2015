package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func LoadInput(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}

func IncrementByOne(n int) int {
	return n + 1
}

func DecrementByOne(n int) int {
	return n - 1
}

func CreateStringSlice(str string) []string {
	return strings.Split(str, "")
}

func main() {
	var result = 0
	var basementHit = false

	input := LoadInput("./input.txt")
	slice := strings.Split(input, "")

	for i, v := range slice {
		switch v {
		case "(":
			result = IncrementByOne(result)
		case ")":
			result = DecrementByOne(result)
		}

		if result == -1 && basementHit == false {
			basementHit = true
			// we `+1` because the problem states that
			// the first char is at index 1
			fmt.Println(i + 1)
		}
	}

	fmt.Println(result)
}
