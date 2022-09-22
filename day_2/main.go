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

func MakeSortedIntSlice(str string) []int {
	s := strings.Split(str, "x")

	var sides []int
	for _, sideString := range s {
		side, _ := strconv.Atoi(sideString)
		sides = append(sides, side)
	}

	sort.Ints(sides)

	return sides
}

func CalculatePresentArea(str string) int {
	s := MakeSortedIntSlice(str)

	return s[0]*s[1]*3 + s[0]*s[2]*2 + s[1]*s[2]*2
}

func CalculateRibbonLength(str string) int {
	s := MakeSortedIntSlice(str)

	return s[0] + s[0] + s[1] + s[1] + s[0]*s[1]*s[2]
}

func main() {
	input := LoadInput("./input.txt")
	presents := MakePresentSlice(input)

	total := 0
	ribbon := 0
	for _, present := range presents {
		total += CalculatePresentArea(present)
		ribbon += CalculateRibbonLength(present)
	}

	fmt.Println(total, "Feet of wrapping paper needed.")
	fmt.Println(ribbon, "Feet of ribbon needed.")
}
