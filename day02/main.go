package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/justynhunter/aoc2015_go/util"
)

func main() {
	input := util.ReadFile("./input.txt")

	boxDimensions := getBoxDimensions(input)

	solution1 := solvePart1(boxDimensions)
	fmt.Println("Day 2 Part 1:", solution1)

	solution2 := solvePart2(boxDimensions)
	fmt.Println("Day 2 Part 2:", solution2)
}

type BoxDimensions struct {
	width  int
	height int
	length int
}

func getBoxDimensions(input string) []BoxDimensions {
	lines := strings.Split(input, "\n")

	var boxDimensions []BoxDimensions
	for _, line := range lines {
		var boxDimension BoxDimensions
		_, err := fmt.Scanf(line, "%dx%dx%d", &boxDimension.width, &boxDimension.height, &boxDimension.length)
		if err != nil {
			panic(err)
		}

		boxDimensions = append(boxDimensions, boxDimension)
	}

	return boxDimensions
}

func solvePart1(boxDimensions []BoxDimensions) int {
	totalPaperNeeded := 0
	for _, dim := range boxDimensions {
		wh := dim.width * dim.height
		wl := dim.width * dim.length
		hl := dim.height * dim.length

		totalPaperNeeded += 2*wh + 2*wl + 2*hl + min(wh, min(wl, hl))
	}

	return totalPaperNeeded
}

func solvePart2(boxDimensions []BoxDimensions) int {
	totalRibbon := 0
	for _, dim := range boxDimensions {
		bow := dim.width * dim.height * dim.length
		ribbon := 2 * slices.Min([]int{dim.width + dim.height, dim.width + dim.length, dim.height + dim.length})

		totalRibbon += ribbon + bow
	}

	return totalRibbon
}
