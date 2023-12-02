package main

import (
	"aoc23/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {
	input := utils.Scan("input.txt", "2")
	result := 0
	for _, game := range input {
		valid := true
		gameIDPattern := regexp.MustCompile(`Game (\d+)`)
		gameID, _ := strconv.Atoi(gameIDPattern.FindStringSubmatch(game)[1])

		subsetPattern := regexp.MustCompile(`(\d+) (\w+)`)
		subsets := subsetPattern.FindAllString(game, -1)

		for _, subset := range subsets {
			numberPattern := regexp.MustCompile(`^(\d+)`)
			strValue := numberPattern.FindStringSubmatch(subset)[1]
			value, _ := strconv.Atoi(strValue)

			switch {
			case strings.Contains(subset, "red"):
				if value > 12 {
					valid = false
				}
			case strings.Contains(subset, "green"):
				if value > 13 {
					valid = false
				}
			case strings.Contains(subset, "blue"):
				if value > 14 {
					valid = false
				}
			default:
				fmt.Println("Invalid color!")
			}
		}

		if valid {
			result += gameID
		} 
	}
	fmt.Println(result)
}

func part2() {
	input := utils.Scan("input.txt", "2")
	result := 0
	for _, game := range input {
		highestRed := 0
		highestGreen := 0
		highestBlue := 0

		subsetPattern := regexp.MustCompile(`(\d+) (\w+)`)
		subsets := subsetPattern.FindAllString(game, -1)

		for _, subset := range subsets {
			numberPattern := regexp.MustCompile(`^(\d+)`)
			strValue := numberPattern.FindStringSubmatch(subset)[1]
			value, _ := strconv.Atoi(strValue)

			switch {
			case strings.Contains(subset, "red"):
				if value > highestRed {
					highestRed = value
				}
			case strings.Contains(subset, "green"):
				if value > highestGreen {
					highestGreen = value
				}
			case strings.Contains(subset, "blue"):
				if value > highestBlue {
					highestBlue = value
				}
			default:
				fmt.Println("Invalid color!")
			}
		}
		result += highestRed * highestGreen * highestBlue
	}
	fmt.Println(result)
}