package main

import (
	"aoc2023/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
seed  soil
0     0
1     1
...   ...
48    48
49    49
50    52
51    53
...   ...
96    98
97    99
98    50
99    51
*/

type mapped struct {
	to   int
	from int
	r    int
}

func soil_to_fertilizer(inp int) int {
	return 0
}

//	func transfer(inp int) int {
//		if inp < 50 || inp > 98 {
//			return inp
//		} else if inp >= 50 && inp <= 98 {
//			return inp + 2
//		} else {
//			return inp - 98
//		}
//	}
func transfer(inp int, m []mapped) int {
	for _, maps := range m {
		if inp >= maps.from {
			n := inp - maps.from
			if n <= maps.r {
				return maps.to + n
			}
		}
	}
	return inp
}

func fillMapped(lines []string, m *[]mapped, i int) {
	for idx := i; idx < len(lines) && lines[idx] != ""; idx++ {
		s := strings.Split(lines[idx], " ")
		from, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		to, err := strconv.Atoi(s[0])
		if err != nil {
			log.Fatal(err)
		}
		r, err := strconv.Atoi(s[2])
		if err != nil {
			log.Fatal(err)
		}
		mp := mapped{
			from: from,
			to:   to,
			r:    r,
		}
		*m = append(*m, mp)
	}
}

func fillSeeds(s []string) []int {
	seeds := make([]int, 0, len(s)-1)
	for i := 1; i < len(s); i++ {
		var n int
		n, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, n)
	}
	return seeds
}

func solve1(lines []string) {
	var seeds []int
	seeds_to_soil := make([]mapped, 0)
	soil_to_fertilizer := make([]mapped, 0)
	fertilizer_to_water := make([]mapped, 0)
	water_to_light := make([]mapped, 0)
	light_to_temperature := make([]mapped, 0)
	temperature_to_humidity := make([]mapped, 0)
	humidity_to_location := make([]mapped, 0)
	for i := range lines {
		s := strings.Split(lines[i], " ")
		switch s[0] {
		case "seeds:":
			seeds = fillSeeds(s)
		case "seed-to-soil":
			fillMapped(lines, &seeds_to_soil, i+1)
		case "soil-to-fertilizer":
			fillMapped(lines, &soil_to_fertilizer, i+1)
		case "fertilizer-to-water":
			fillMapped(lines, &fertilizer_to_water, i+1)
		case "water-to-light":
			fillMapped(lines, &water_to_light, i+1)
		case "light-to-temperature":
			fillMapped(lines, &light_to_temperature, i+1)
		case "temperature-to-humidity":
			fillMapped(lines, &temperature_to_humidity, i+1)
		case "humidity-to-location":
			fillMapped(lines, &humidity_to_location, i+1)
		}

	}
	for i := range seeds {
		seeds[i] = transfer(seeds[i], seeds_to_soil)
		seeds[i] = transfer(seeds[i], soil_to_fertilizer)
		seeds[i] = transfer(seeds[i], fertilizer_to_water)
		seeds[i] = transfer(seeds[i], water_to_light)
		seeds[i] = transfer(seeds[i], light_to_temperature)
		seeds[i] = transfer(seeds[i], temperature_to_humidity)
		seeds[i] = transfer(seeds[i], humidity_to_location)
	}
	fmt.Println(seeds)
	minN := seeds[0]
	for _, v := range seeds {
		minN = min(v, minN)
	}
	fmt.Println(minN)
}

func main() {
	lines := common.FileOpener()
	test_lines := common.TestFileOpener()
	solve1(test_lines)
	solve1(lines)
}
