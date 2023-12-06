package main

import (
	"aoc2023/common"
	"fmt"
	"strconv"
	"strings"
)

type race struct {
	Time     int
	Distance int
}

// So, the example from before:
//
// Time:      7  15   30
// Distance:  9  40  200
// ...now instead means this:
//
// Time:      71530
// Distance:  940200
// Now, you have to figure out how many ways there are to win this single race. In this example, the race lasts for 71530 milliseconds and the record distance you need to beat is 940200 millimeters. You could hold the button anywhere from 14 to 71516 milliseconds and beat the record, a total of 71503 ways!

// time * speed = distance
//

func solve2(lines []string) {
	times := strings.Split(lines[0], " ")[1:]
	distances := strings.Split(lines[1], " ")[1:]
	t, _ := strconv.Atoi(strings.Join(times, ""))
	d, _ := strconv.Atoi(strings.Join(distances, ""))
	r := race{
		Time:     t,
		Distance: d,
	}
	time := r.Time
	var w int
	var fl int
	var sl int
	for button_hold := 0; button_hold <= r.Time; button_hold++ {
		speed := button_hold
		if speed*time > r.Distance {
			fl = speed
			fmt.Println("fl", fl)
			break
		}
		time--
	}
	time = 0
	for button_hold := r.Time; button_hold > 0; button_hold-- {
		speed := button_hold
		if speed*time > r.Distance {
			sl = speed
			fmt.Println("sl", sl)
			break
		}
		time++
	}
	fmt.Println(sl - fl + 1)
	fmt.Println(w)
}

func solve1(lines []string) {
	times := strings.Split(lines[0], " ")[1:]
	new_times := []string{}
	for _, t := range times {
		if t != "" {
			new_times = append(new_times, t)
		}
	}
	times = new_times
	distances := strings.Split(lines[1], " ")[1:]
	new_distances := []string{}
	for _, t := range distances {
		if t != "" {
			new_distances = append(new_distances, t)
		}
	}
	distances = new_distances
	races := make([]race, 0, len(times))
	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		races = append(races, race{t, d})
	}
	wins := []int{}
	for _, race := range races {
		t := race.Time
		var w int
		for button_hold := 0; button_hold <= race.Distance; button_hold++ {
			speed := button_hold
			if speed*t > race.Distance {
				w++
				fmt.Println(speed)
			}
			t--
		}
		wins = append(wins, w)
	}
	res := 1
	for _, w := range wins {
		res *= w
	}
	fmt.Println(wins)
	fmt.Println(res)
}

func main() {
	lines := common.FileOpener()
	test_lines := common.TestFileOpener()
	solve1(lines)
	solve2(test_lines)
	solve2(lines)
}
