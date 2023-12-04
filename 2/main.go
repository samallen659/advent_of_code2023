package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/samallen659/advent_of_code2023/pkg/utils"
)

var data = utils.ReadInput("/2/input.txt")

func main() {
	fmt.Println(part1(data))
}

type Round struct {
	blue  int
	red   int
	green int
}

type Game struct {
	id     int
	rounds []Round
}

func part1(data string) int {
	lines := strings.Split(data, "\n")

	var games []Game
	for i, line := range lines {
		var game Game
		game.id = i + 1
		roundsInput := strings.Split(strings.Split(line, ": ")[1], ";")
		for _, roundInput := range roundsInput {
			var round Round
			r, _ := regexp.Compile("\\d{1,2} blue")
			if r.Match([]byte(roundInput)) {
				blue := r.FindStringSubmatch(roundInput)
				blueCount, _ := strconv.Atoi(strings.Split(blue[0], " ")[0])
				round.blue += blueCount
			}
			r, _ = regexp.Compile("\\d{1,2} red")
			if r.Match([]byte(roundInput)) {
				red := r.FindStringSubmatch(roundInput)
				redCount, _ := strconv.Atoi(strings.Split(red[0], " ")[0])
				round.red += redCount
			}
			r, _ = regexp.Compile("\\d{1,2} green")
			if r.Match([]byte(roundInput)) {
				green := r.FindStringSubmatch(roundInput)
				greenCount, _ := strconv.Atoi(strings.Split(green[0], " ")[0])
				round.green += greenCount
			}
			game.rounds = append(game.rounds, round)
		}

		games = append(games, game)
	}

	var count int
loop:
	for _, g := range games {
		for _, r := range g.rounds {
			if r.blue > 14 || r.red > 12 || r.green > 13 {
				continue loop
			}

		}
		count += g.id
	}

	return count
}
