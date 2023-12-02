package main

import (
	"aoc23/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {
	var input = utils.Scan("input.txt", "1")
	var result int
	for _, line := range input {
		var first string
		var last string

		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = string(line[i])
				break	
				
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = string(line[i])
				break
			}
		}

		var sum = first + last
		var num, _ = strconv.Atoi(sum)
		result += num
	}
	fmt.Println(result)
}

func part2() {
	var input = utils.Scan("input.txt", "1")
	var result int
	
	for _, line := range input {
		var first string
		var last string

		// Replace all spelled-out numbers with digits
		line = replaceDigitWords(line)

		fmt.Println(line)

		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				first = string(line[i])
				break	
				
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = string(line[i])
				break
			}
		}

		var sum = first + last
		var num, _ = strconv.Atoi(sum)
		result += num
	}
	fmt.Println("Result: " + fmt.Sprint(result))
}

func replaceDigitWords(input string) string {
    numberWords := map[string]string{
        "one":   "1",
        "two":   "2",
        "three": "3",
        "four":  "4",
        "five":  "5",
        "six":   "6",
        "seven": "7",
        "eight": "8",
        "nine":  "9",
    }

    // Sort the keys by length in descending order
    var keys []string
    for k := range numberWords {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(i, j int) bool {
        return len(keys[i]) > len(keys[j])
    })

    var result strings.Builder
    i := 0
    for i < len(input) {
        replaced := false
        for _, key := range keys {
            if strings.HasPrefix(input[i:], key) {
                result.WriteString(numberWords[key])
                i += len(key)
                replaced = true
                break
            }
        }
        if !replaced {
            result.WriteByte(input[i])
            i++
        }
    }

    return result.String()
}