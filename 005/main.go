package main

import (
	"aoc2023/common"
	"fmt"
	"strings"
)

type seedRange struct {
	from  int
	count int
}

type mapped struct {
	destinationStart int
	sourceStart      int
	length           int
}

func solve2(lines []string) {
	var seeds []seedRange
	seedsToSoil := make([]mapped, 0)
	soilToFertilizer := make([]mapped, 0)
	fertilizerToWater := make([]mapped, 0)
	waterToLight := make([]mapped, 0)
	lightToTemperature := make([]mapped, 0)
	temperatureToHumidity := make([]mapped, 0)
	humidityToLocation := make([]mapped, 0)

	// Parse the almanac data
	for i := range lines {
		s := strings.Split(lines[i], " ")
		switch s[0] {
		case "seeds:":
			seeds = fillSeeds(s[1:])
		case "seed-to-soil":
			fillMapped(lines, &seedsToSoil, i+1)
		case "soil-to-fertilizer":
			fillMapped(lines, &soilToFertilizer, i+1)
		case "fertilizer-to-water":
			fillMapped(lines, &fertilizerToWater, i+1)
		case "water-to-light":
			fillMapped(lines, &waterToLight, i+1)
		case "light-to-temperature":
			fillMapped(lines, &lightToTemperature, i+1)
		case "temperature-to-humidity":
			fillMapped(lines, &temperatureToHumidity, i+1)
		case "humidity-to-location":
			fillMapped(lines, &humidityToLocation, i+1)
		}
	}

	minLocation := 9223372036854775807
	for _, seed := range seeds {
		for j := seed.from; j < seed.from+seed.count; j++ {
			location := transfer(j, seedsToSoil)
			location = transfer(location, soilToFertilizer)
			location = transfer(location, fertilizerToWater)
			location = transfer(location, waterToLight)
			location = transfer(location, lightToTemperature)
			location = transfer(location, temperatureToHumidity)
			location = transfer(location, humidityToLocation)
			if location < minLocation {
				minLocation = location
			}
		}
	}

	fmt.Println(minLocation)
}

func fillSeeds(data []string) []seedRange {
	seeds := make([]seedRange, 0)
	for i := 0; i < len(data); i += 2 {
		from := parseInt(data[i])
		count := parseInt(data[i+1])
		seeds = append(seeds, seedRange{from, count})
	}
	return seeds
}

func fillMapped(lines []string, mappedSlice *[]mapped, startIndex int) {
	for i := startIndex; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")
		if len(line) != 3 {
			break
		}
		destinationStart := parseInt(line[0])
		sourceStart := parseInt(line[1])
		length := parseInt(line[2])
		*mappedSlice = append(*mappedSlice, mapped{destinationStart, sourceStart, length})
	}
}

func transfer(value int, mappedSlice []mapped) int {
	for _, m := range mappedSlice {
		if value >= m.sourceStart && value < m.sourceStart+m.length {
			return value - m.sourceStart + m.destinationStart
		}
	}
	return value
}

func parseInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}

func main() {
	lines := []string{
		"seeds: 79 14 55 13",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}

	solve2(lines)
	lines2 := common.FileOpener()
	// test_lines := common.TestFileOpener()

	//solve2(test_lines)
	solve2(lines2)
}
