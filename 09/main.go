package main

import (
	"aoc2023/common"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func seqs(lines []string) [][]int {
	seqs := make([][]int, 0, len(lines))
	for _, line := range lines {
		L := strings.Split(line, " ")
		v := make([]int, 0, len(L))
		for _, l := range L {
			n, err := strconv.Atoi(l)
			if err != nil {
				log.Fatal()
			}
			v = append(v, n)
		}
		seqs = append(seqs, v)
	}
	return seqs
}

func solve2(lines []string) {
	seqs := seqs(lines)
	zeros := false
	sums := 0
	for _, seq := range seqs {
		var s []int
		s = append(s, seq[0])
		for !zeros {
			seq, zeros = createDiffSeq(seq)
			s = append(s, seq[0])
		}
		for i := len(s) - 2; i >= 0; i-- {
			s[i] = s[i] - s[i+1]
		}
		var t int
		t += s[0]
		sums += t
		zeros = false
	}
	fmt.Println(sums)
}

func createDiffSeq(seq []int) ([]int, bool) {
	zeros := true
	diff := make([]int, len(seq)-1)
	for i := 0; i < len(diff); i++ {
		diff[i] = seq[i+1] - seq[i]
		if diff[i] != 0 {
			zeros = false
		}
	}
	return diff, zeros
}

func solve1(lines []string) {
	seqs := seqs(lines)
	zeros := false
	sums := 0
	for _, seq := range seqs {
		var t int
		for !zeros {
			t += seq[len(seq)-1]
			seq, zeros = createDiffSeq(seq)
		}
		zeros = false
		sums += t
	}
	fmt.Println(sums)
}

func main() {
	test := common.TestFileOpener()
	solve1(test)
	t := common.FileOpener()
	solve1(t)
	solve2(test)
	solve2(t)
}
