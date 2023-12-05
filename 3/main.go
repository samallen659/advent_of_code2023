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

	fmt.Println(total)
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
