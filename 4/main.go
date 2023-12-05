package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code2023/pkg/utils"
	"golang.org/x/exp/slices"
)

func main() {
	data := utils.ReadInput("/4/input.txt")
	fmt.Println(part1(data))
}

func part1(data string) int {
	lines := strings.Split(data, "\n")

	total := 0
	for _, line := range lines {
		numbers := strings.Split(strings.Split(line, ": ")[1], " | ")
		winningNumbers := getNumbers(numbers[0])
		cardNumbers := getNumbers(numbers[1])

		matchCount := 0
		for _, win := range winningNumbers {
			if slices.Contains(cardNumbers, win) {
				matchCount++
			}
		}

		total += int(math.Pow(2, float64(matchCount-1)))
	}

	return total
}

func getNumbers(numberStr string) []int {
	var numbers []int
	r, _ := regexp.Compile(`\d+`)
	numStr := r.FindAllStringSubmatch(numberStr, 100)
	for _, n := range numStr {
		num, err := strconv.Atoi(n[0])
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, num)
	}

	return numbers
}
