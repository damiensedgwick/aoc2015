package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func LoadInput(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}

func MakePresentSlice(str string) []string {
	return strings.Split(str, "\n")
}

func CalculatePresentArea(str string) int {
	s := strings.Split(str, "x")

	var sides []int
	for _, sideString := range s {
		side, _ := strconv.Atoi(sideString)
		sides = append(sides, side)
	}

	sort.Ints(sides)

	return sides[0]*sides[1]*3 + sides[0]*sides[2]*2 + sides[1]*sides[2]*2
}

func main() {
	input := LoadInput("./input.txt")
	presents := MakePresentSlice(input)

	total := 0
	for _, present := range presents {
		total += CalculatePresentArea(present)
	}

	fmt.Println(total)
}
