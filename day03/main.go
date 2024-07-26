package main

import (
	"fmt"
	"slices"

	"github.com/justynhunter/aoc2015_go/util"
)

func main() {
	input := util.ReadFile("./input.txt")

	solution1 := solvePart1(input)
	fmt.Println("Day 1, Part 1:", solution1)

	solution2 := solvePart2(input)
	fmt.Println("Day 2, Part 2:", solution2)
}

type Position struct {
	x int
	y int
}

func solvePart1(input string) int {
	curr := Position{x: 0, y: 0}
	var visited []Position

	for _, c := range input {
		if !slices.Contains(visited, curr) {
			visited = append(visited, Position{x: curr.x, y: curr.y})
		}

		if c == '>' {
			curr.x++
		} else if c == 'v' {
			curr.y--
		} else if c == '<' {
			curr.x--
		} else {
			curr.y++
		}
	}

	return len(visited)
}

func solvePart2(input string) int {
	santa := Position{x: 0, y: 0}
	robo := Position{x: 0, y: 0}
	var visited []Position

	for i, c := range input {
		if !slices.Contains(visited, santa) {
			visited = append(visited, Position{x: santa.x, y: santa.y})
		}

		if !slices.Contains(visited, robo) {
			visited = append(visited, Position{x: robo.x, y: robo.y})
		}

		if c == '>' {
			if i%2 == 0 {
				robo.x++
			} else {
				santa.x++
			}
		} else if c == 'v' {
			if i%2 == 0 {
				robo.y--
			} else {
				santa.y--
			}
		} else if c == '<' {
			if i%2 == 0 {
				robo.x--
			} else {
				santa.x--
			}
		} else {
			if i%2 == 0 {
				robo.y++
			} else {
				santa.y++
			}
		}
	}

	return len(visited)
}
