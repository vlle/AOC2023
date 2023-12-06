package main

import (
	"aoc2023/common"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type mapped struct {
	to   int
	from int
	r    int
}

func transfer(inp int, m []mapped) int {
	for _, maps := range m {
		if inp >= maps.from {
			n := inp - maps.from
			if n < maps.r {
				return maps.to + n
			}
		}
	}
	return inp
}

func mergeSlice(sl []seedRange) []seedRange {
	sort.Slice(sl, func(i int, j int) bool { return sl[i].from < sl[j].from })
	fmt.Println(sl)
	new_sl := []seedRange{sl[0]}
	for i := 1; i < len(sl); i++ {
		if (sl[i].from + sl[i].count) < (new_sl[len(new_sl)-1].count + new_sl[len(new_sl)-1].from) {
			new_sl[len(new_sl)-1].count = max(new_sl[len(new_sl)-1].count, sl[i].count)
		} else {
			new_sl = append(new_sl, sl[i])
		}
	}
	return new_sl
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

func fillSeeds2(s []string) []seedRange {
	seeds := make([]seedRange, 0, len(s)-1)
	for i := 1; i < len(s)-1; i++ {
		var n int
		n, err := strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
		var n2 int
		n2, err = strconv.Atoi(s[i+1])
		if err != nil {
			log.Fatal(err)
		}
		sr := seedRange{
			from:  n,
			count: n2,
		}
		seeds = append(seeds, sr)
		i++
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
	minN := seeds[0]
	for _, v := range seeds {
		minN = min(v, minN)
	}
	fmt.Println(minN)
}

var memo = map[int]int{}

type seedRange struct {
	from  int
	count int
}

func solve2(lines []string) {
	var seeds []seedRange
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
			seeds = fillSeeds2(s)
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
	// fmt.Println(seeds)
	// seeds = mergeSlice(seeds)
	fmt.Println(len(seeds))
	var wg sync.WaitGroup
	ch := make(chan int)
	for i := range seeds {
		wg.Add(1)
		go func(seeds []seedRange) {
			for _, seed := range seeds {
				mpn := 9223372036854775807
				for j := seed.from; j < seed.from+seed.count; j++ {
					pn := transfer(j, seeds_to_soil)
					pn = transfer(pn, soil_to_fertilizer)
					pn = transfer(pn, fertilizer_to_water)
					pn = transfer(pn, water_to_light)
					pn = transfer(pn, light_to_temperature)
					pn = transfer(pn, temperature_to_humidity)
					pn = transfer(pn, humidity_to_location)
					if pn < mpn {
						mpn = pn
					}
				}
				ch <- mpn
			}
			wg.Done()
		}(seeds[i : i+2])
	}
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		pn := 9223372036854775807
		for n := range ch {
			pn = min(pn, n)
		}
		fmt.Println(pn)
		wg2.Done()
	}()
	wg.Wait()
	close(ch)
	wg2.Wait()
}

func main() {
	lines := common.FileOpener()
	test_lines := common.TestFileOpener()
	solve1(test_lines)
	solve1(lines)

	solve2(test_lines)
	solve2(lines)
}
