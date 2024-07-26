package main

import (
	"fmt"
	"strings"

	"github.com/justynhunter/aoc2015_go/util"
)

func main() {
	input := util.ReadFile("./input.txt")

	solution1 := solvePart1(input)
	fmt.Println("Day 1, Part 1:", solution1)

	solution2 := solvePart2(input)
	fmt.Println("Day 2, Part 2:", solution2)
}

func solvePart1(input string) int {
	ups := strings.Count(input, "(")
	downs := strings.Count(input, ")")

	return ups - downs
}

func solvePart2(input string) int {
	level := 0
	for i, c := range input {
		if c == '(' {
			level++
		} else if level == 0 {
			return i + 1
		} else {
			level--
		}
	}

	return -1
}
