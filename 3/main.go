package main

import (
	"aoc23/utils"
	"fmt"
	"image"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	input := utils.Scan("input.txt", "3")

	grid := buildGrid(input)
	part1, part2 := calculateParts(input, grid)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func buildGrid(input []string) map[image.Point]rune {
	grid := make(map[image.Point]rune)
	for y, line := range input {
		for x, char := range line {
			if char != '.' && !unicode.IsDigit(char) {
				grid[image.Point{X: x, Y: y}] = char
			}
		}
	}
	return grid
}

func calculateParts(input []string, grid map[image.Point]rune) (int, int) {
	part1, part2 := 0, 0
	parts := make(map[image.Point][]int)

	for y, line := range input {
		for _, match := range regexp.MustCompile(`\d+`).FindAllStringIndex(line, -1) {
			bounds := calculateBounds(match, y)

			number, _ := strconv.Atoi(line[match[0]:match[1]])
			for point := range bounds {
				if _, exists := grid[point]; exists {
					parts[point] = append(parts[point], number)
					part1 += number
				}
			}
		}
	}

	for point, numbers := range parts {
		if grid[point] == '*' && len(numbers) == 2 {
			part2 += numbers[0] * numbers[1]
		}
	}
	return part1, part2
}

func calculateBounds(match []int, y int) map[image.Point]struct{} {
	bounds := make(map[image.Point]struct{})
	for x := match[0]; x < match[1]; x++ {
		for _, delta := range []image.Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
			bounds[image.Point{X: x, Y: y}.Add(delta)] = struct{}{}
		}
	}
	return bounds
}
