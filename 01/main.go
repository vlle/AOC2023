package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func FileOpener() []string {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func solve1() {
	lines := FileOpener()
	var res int
	for _, line := range lines {
		var num string
		for i := range line {
			if line[i] >= '0' && line[i] <= '9' {
				num += string(line[i])
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				num += string(line[i])
				break
			}
		}
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		res += n

	}
	fmt.Println(res)
}

func floatingWindowForward(i int, s string, m map[string]string) string {
	var res string
	for k := i; k < len(s); k++ {
		if v, ok := m[s[i:k]]; ok {
			res += v
			break
		}
		if len(s[i:k]) > len("seven") {
			i = k
		}
	}
	return res
}

func floatingWindowBackwards(i int, s string, m map[string]string) string {
	var res string
	for k := i; k >= 0; k-- {
		fmt.Println(s[k:i])
		if v, ok := m[s[k:i]]; ok {
			res += v
			break
		}
		if len(s[k:i]) >= len("seven") {
			i = k
		}
	}
	return res
}

func getNum(s string, m map[string]string) string {
	for i := 0; i < len(s); i++ {
		n := 0
		if i+1 > len(s) {
			n = len(s)
		} else {
			n = i + 1
		}
		fmt.Println("try", s[:n])
		if v, ok := m[s[:n]]; ok {
			fmt.Println("v", s[:n])
			return v
		}
	}
	return ""
}

func solve2_line(line string) int {

	m := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var num string
	fmt.Println("line", line)
	for i := range line {
		to := len("seven") + 1
		if (len(line) - i) < to {
			to = len(line) - i
		}
		r := getNum(line[i:i+to], m)
		if r != "" {
			num += string(r)
			break
		}
		if line[i] >= '0' && line[i] <= '9' {
			num += string(line[i])
			break
		}
	}
	fmt.Println(num)
	for i := len(line) - 1; i >= 0; i-- {
		to := len("seven") + 1
		if (len(line) - i) < to {
			to = len(line) - i
		}
		r := getNum(line[i:i+to], m)
		if r != "" {
			num += r
			break
		}
		if line[i] >= '0' && line[i] <= '9' {
			num += string(line[i])
			break
		}
	}
	fmt.Println(num)
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func solve2() {
	lines := FileOpener()
	var res int
	for _, line := range lines {
		res += solve2_line(line)
	}
	fmt.Println(res)
}

func main() {
	solve1()
	solve2()
}
