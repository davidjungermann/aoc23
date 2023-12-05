package main

import (
	"aoc23/utils"
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	lines := utils.Scan("input.txt", "4")
	total := 0
	for _, l := range lines {
		count := getWinningCountForLine(l)
		score := convertCountToScore(count)
		total += score
	}
	fmt.Printf("Part 1: %d\n", total)
}

func convertCountToScore(count int) int {
	score := 0
	for n := 1; n <= count; n++ {
		if n == 1 {
			score = 1
		} else {
			score *= 2
		}
	}
	return score
}

func getWinningCountForLine(l string) int {
	parts := strings.Split(l, "|")
	partOne := strings.Split(strings.TrimSpace(parts[0]), ":")
	winningNumbers := strings.Split(strings.TrimSpace(partOne[1]), " ")
	ourNumbers := strings.Split(parts[1], " ")

	count := 0
	for _, winning := range winningNumbers {
		win := strings.TrimSpace(winning)
		if win == "" {
			continue
		}
		for _, ours := range ourNumbers {
			thisNumber := strings.TrimSpace(ours)
			if thisNumber == "" {
				continue
			}
			if thisNumber == win {
				count += 1
			}
		}
	}
	return count
}

func part2() {
	lines := utils.Scan("input.txt", "4")
	cards := make([]int, len(lines))

	for idx := range lines {
		cards[idx] = 1
	}

	for idx, l := range lines {
		wincount := getWinningCountForLine(l)
		counttil := idx + wincount
		if counttil >= len(lines) {
			counttil = len(lines) - 1
		}
		for n := idx + 1; n <= counttil; n++ {
			cards[n] += cards[idx]
		}
	}

	total := 0
	for idx := range lines {
		total += cards[idx]
	}
	fmt.Printf("Part 2: %d\n", total)
}
