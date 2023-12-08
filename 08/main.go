package main

import (
	"aoc2023/common"
	"fmt"
	"strings"
)

// 		RL
//
// AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)

type ways struct {
	left  string
	right string
}

func solve1(lines []string) {
	instr := lines[0]
	m := make(map[string]ways)
	for i := range lines {
		if i == 0 || i == 1 {
			continue
		}
		lines[i] = strings.ReplaceAll(lines[i], "=", "")
		lines[i] = strings.ReplaceAll(lines[i], "(", "")
		lines[i] = strings.ReplaceAll(lines[i], ")", "")
		lines[i] = strings.ReplaceAll(lines[i], ",", "")
		var s []string
		for _, w := range strings.Split(lines[i], " ") {
			if w != "" {
				s = append(s, w)
			}
		}
		m[s[0]] = ways{left: s[1], right: s[2]}
		fmt.Println(m[s[0]])
	}
	p := "AAA"
	point := m[p]
	path := 0
	for i := 0; p != "ZZZ"; i++ {
		if instr[i%len(instr)] == 'L' {
			p = point.left
		} else if instr[i%len(instr)] == 'R' {
			p = point.right
		}
		point = m[p]
		path++
	}
	fmt.Println(path)

}

func solve2(lines []string) {
	instr := lines[0]
	m := make(map[string]ways)
	starts := []string{}
	for i := range lines {
		if i == 0 || i == 1 {
			continue
		}
		lines[i] = strings.ReplaceAll(lines[i], "=", "")
		lines[i] = strings.ReplaceAll(lines[i], "(", "")
		lines[i] = strings.ReplaceAll(lines[i], ")", "")
		lines[i] = strings.ReplaceAll(lines[i], ",", "")
		var s []string
		for _, w := range strings.Split(lines[i], " ") {
			if w != "" {
				s = append(s, w)
			}
		}
		m[s[0]] = ways{left: s[1], right: s[2]}
		if strings.HasSuffix(s[0], "A") {
			starts = append(starts, s[0])
		}
	}
	points := make([]ways, 0, len(starts))
	for _, s := range starts {
		points = append(points, m[s])
	}
	path := 0
	not_ends_z := true
	fmt.Println(starts)
	fmt.Println(points)
	one := 0
	for i := 0; not_ends_z; i++ {
		if instr[i%len(instr)] == 'L' {
			for i := range starts {
				starts[i] = points[i].left
			}
		} else if instr[i%len(instr)] == 'R' {
			for i := range starts {
				starts[i] = points[i].right
			}
		}
		for i := range points {
			points[i] = m[starts[i]]
		}
		if one < 10 {
			fmt.Println(starts)
			one++
		}
		path++
		not_ends_z = false
		for i := range starts {
			if !strings.HasSuffix(starts[i], "Z") {
				not_ends_z = true
				break
			}
		}
		// fmt.Println(p1, p2)
	}
	fmt.Println(path)

}

func solve3(lines []string) {
	instr := lines[0]
	m := make(map[string]ways)
	starts := []string{}
	for i := range lines {
		if i == 0 || i == 1 {
			continue
		}
		lines[i] = strings.ReplaceAll(lines[i], "=", "")
		lines[i] = strings.ReplaceAll(lines[i], "(", "")
		lines[i] = strings.ReplaceAll(lines[i], ")", "")
		lines[i] = strings.ReplaceAll(lines[i], ",", "")
		var s []string
		for _, w := range strings.Split(lines[i], " ") {
			if w != "" {
				s = append(s, w)
			}
		}
		m[s[0]] = ways{left: s[1], right: s[2]}
		if strings.HasSuffix(s[0], "A") {
			starts = append(starts, s[0])
		}
	}
	ms := []int{}
	for _, p := range starts {
		ms = append(ms, (solveCase(instr, p, m)))
	}
	lcm := 1
	for _, num := range ms {
		lcm = lcm * num / gcd(lcm, num)
	}
	fmt.Println(lcm)
}

func gcd(n, m int) int {
	if m == 0 {
		return n
	}
	return gcd(m, n%m)
}

// def gcd(n, m):
//     if m == 0:
//         return n
//     return gcd(m, n % m)
//
// A = [10, 25, 37, 15, 75, 12]
//
// lcm = 1
// for i in A:
//     lcm = lcm * i // gcd(lcm, i)
//
// print(lcm)

func solveCase(instr, p string, m map[string]ways) int {
	point := m[p]
	path := 0
	for i := 0; !strings.HasSuffix(p, "Z"); i++ {
		if instr[i%len(instr)] == 'L' {
			p = point.left
		} else if instr[i%len(instr)] == 'R' {
			p = point.right
		}
		point = m[p]
		path++
	}
	return path
}

func main() {
	t := common.FileOpener()
	test := common.TestFileOpener()
	// solve2(test)
	// solve2(t)
	solve3(test)
	solve3(t)
}
