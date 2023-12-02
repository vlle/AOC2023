package main

import (
	"aoc2023/common"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

/*
The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
*/

/*
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/

func GameFailed(count int, t string) bool {
	m := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	return m[t] < count
}

func solveGame(line string) int {
	gameNum := strings.TrimPrefix(line, "Game ")
	var sb strings.Builder
	var start int
	for idx, letter := range gameNum {
		if letter == ':' {
			start = idx
			break
		}
		sb.WriteRune(letter)
	}
	num, err := strconv.Atoi(sb.String())
	if err != nil {
		log.Fatal(err)
	}
	l := strings.TrimSpace(gameNum)
	v := strings.Split(l[start+1:], "; ")
	allowed_colors := []string{"red", "blue", "green"}
	for _, game := range v {
		g := strings.Split(game, " ")
		for i := range g {
			g[i] = strings.TrimSuffix(g[i], ",")
			if idx := slices.Index(allowed_colors, g[i]); idx != -1 {
				cubes, err := strconv.Atoi(g[i-1])
				if err != nil {
					log.Fatal(err)
				}
				if GameFailed(cubes, g[i]) {
					return 0
				}
			}
		}
	}
	return num
}

func solveGame2(line string) int {
	gameNum := strings.TrimPrefix(line, "Game ")
	var sb strings.Builder
	var start int
	for idx, letter := range gameNum {
		if letter == ':' {
			start = idx
			break
		}
		sb.WriteRune(letter)
	}
	l := strings.TrimSpace(gameNum)
	v := strings.Split(l[start+1:], "; ")
	allowed_colors := []string{"red", "blue", "green"}
	m := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, game := range v {
		g := strings.Split(game, " ")
		for i := range g {
			g[i] = strings.TrimSuffix(g[i], ",")
			if idx := slices.Index(allowed_colors, g[i]); idx != -1 {
				cubes, err := strconv.Atoi(g[i-1])
				if err != nil {
					log.Fatal(err)
				}
				m[g[i]] = max(m[g[i]], cubes)
			}
		}
	}
	return m["red"] * m["green"] * m["blue"]
}

func solve1(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += solveGame(line)
	}
	fmt.Println(sum)
}

func solve2(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += solveGame2(line)
	}
	fmt.Println(sum)
}

func main() {
	lines := common.FileOpener()
	solve1(lines)
	solve2(lines)
}
