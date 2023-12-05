package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code2023/pkg/utils"
)

type index struct {
	y int
	x int
}

type numberIndex struct {
	y      int
	xStart int
	xEnd   int
}

func main() {
	data := utils.ReadInput("/3/input.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func part1(data string) int {
	lines := strings.Split(data, "\n")
	symIdxs := getSymbolIndexes(lines)
	numIdxs := getNumberIndexes(lines)

	numSet := make(map[numberIndex]bool)
	for _, symIdx := range symIdxs {
		for _, numIdx := range numIdxs {
			if numIdx.y >= symIdx.y-1 && numIdx.y <= symIdx.y+1 {
				if (numIdx.xEnd >= symIdx.x && numIdx.xEnd <= symIdx.x+2) || (numIdx.xStart >= symIdx.x-1 && numIdx.xStart <= symIdx.x+1) {
					numSet[numIdx] = true
				}
			}
		}
	}

	total := 0
	for key := range numSet {
		num, err := strconv.Atoi(lines[key.y][key.xStart:key.xEnd])
		if err != nil {
			log.Fatal(err)
		}
		total += num
	}

	return total
}

func part2(data string) int {
	lines := strings.Split(data, "\n")
	starIdxs := getStarSymbolIndexes(lines)
	numIdxs := getNumberIndexes(lines)

	total := 0
	for _, star := range starIdxs {
		count := 0
		ratio := 1
		for _, numIdx := range numIdxs {
			if numIdx.y >= star.y-1 && numIdx.y <= star.y+1 {
				if (numIdx.xEnd >= star.x && numIdx.xEnd <= star.x+2) || (numIdx.xStart >= star.x-1 && numIdx.xStart <= star.x+1) {
					count++
					num, err := strconv.Atoi(lines[numIdx.y][numIdx.xStart:numIdx.xEnd])
					if err != nil {
						log.Fatal(err)
					}
					ratio *= num
				}
			}
		}
		if count == 2 {
			total += ratio
		}
	}

	return total
}

func getStarSymbolIndexes(lines []string) []index {
	var indexes []index
	for i, line := range lines {
		r, err := regexp.Compile(`\*`)
		if err != nil {
			log.Fatal(err)
		}
		idxs := r.FindAllStringSubmatchIndex(line, 100)
		for _, idx := range idxs {
			indexes = append(indexes, index{x: idx[0], y: i})
		}
	}

	return indexes
}

func getNumberIndexes(lines []string) []numberIndex {
	var indexes []numberIndex
	for i, line := range lines {
		r, err := regexp.Compile(`\d+`)
		if err != nil {
			log.Fatal(err)
		}
		idxs := r.FindAllStringSubmatchIndex(line, 100)
		for _, idx := range idxs {
			indexes = append(indexes, numberIndex{y: i, xStart: idx[0], xEnd: idx[1]})
		}
	}
	return indexes
}

func getSymbolIndexes(lines []string) []index {
	var indexes []index
	for i, line := range lines {
		r, err := regexp.Compile(`(\*|#|\+|\$|!|@|£|\^|%|&|>|<|-|_|=|~|"|/|\\)`)
		if err != nil {
			log.Fatal(err)
		}
		idxs := r.FindAllStringSubmatchIndex(line, 100)
		for _, idx := range idxs {
			indexes = append(indexes, index{x: idx[0], y: i})
		}
	}

	return indexes
}
