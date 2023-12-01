package main

import (
	"fmt"
	"github.com/samallen659/advent_of_code2023/pkg/utils"
	"log"
	"strconv"
	"strings"
)

var data = utils.ReadInput("/1/input.txt")

func main() {
	fmt.Println(part1(data))
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
