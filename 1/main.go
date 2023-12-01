package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code2023/pkg/utils"
)

var data = utils.ReadInput("/1/input.txt")

var digitStrings = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var digitMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func main() {
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data string) int {
	lines := strings.Split(data, "\n")

	var numbers []int
	for _, line := range lines {
		chars := strings.Split(line, "")
		var numChars []string
		for _, char := range chars {
			_, err := strconv.Atoi(char)
			if err != nil {
				continue
			}
			numChars = append(numChars, char)
		}
		num, err := strconv.Atoi(numChars[0] + numChars[len(numChars)-1])
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	total := 0
	for i := range numbers {
		total += numbers[i]
	}

	return total
}

func part2(data string) int {
	lines := strings.Split(data, "\n")

	var numbers []int
	for _, line := range lines {
		for _, digit := range digitStrings {
			r, _ := regexp.Compile(digit)
			line = string(r.ReplaceAll([]byte(line), []byte(digitMap[digit])))
		}
		chars := strings.Split(line, "")
		var numChars []string
		for _, char := range chars {
			_, err := strconv.Atoi(char)
			if err != nil {
				continue
			}
			numChars = append(numChars, char)
		}
		num, err := strconv.Atoi(numChars[0] + numChars[len(numChars)-1])
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	total := 0
	for i := range numbers {
		total += numbers[i]
	}

	return total
}
