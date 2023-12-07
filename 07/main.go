package main

import (
	"aoc2023/common"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

/*
In Camel Cards, you get a list of hands, and your goal is to order them based on the strength of each hand. A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order, where A is the highest and 2 is the lowest.

Every hand is exactly one type. From strongest to weakest, they are:

Five of a kind, where all five cards have the same label: AAAAA
Four of a kind, where four cards have the same label and one card has a different label: AA8AA
Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
High card, where all cards' labels are distinct: 23456

*/

/*
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
6440
*/

var card_rank = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var card_rank2 = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 1,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type play_hand struct {
	Hand     [16]int
	HandLet  []string
	Bid      int
	Rank     int
	HandType int
	Points   int
}

const fivek = 7
const fourk = 6
const fullhouse = 5
const threek = 4
const twopair = 3
const onepair = 2
const highcard = 1
const no = 0

func numerousKind(h [16]int, cap int, reward int) (int, int) {
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] == cap {
			return reward, h[i] * i
		}
	}
	return no, 0
}

func fullHouse(h [16]int) (int, int) {
	r := slices.Contains(h[:], 3) && slices.Contains(h[:], 2)
	if r {
		points := 0
		for i := range h {
			points += i * h[i]
		}
		return fullhouse, points
	}
	return no, 0
}

func numerousPair(h [16]int) (int, int) {
	var paircount int
	var points int
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] == 2 {
			points += i * 2
			paircount++
		}
	}
	if paircount == 2 {
		return twopair, points
	}
	if paircount == 1 {
		return onepair, points
	}
	return no, 0
}

func highCard(h [16]int) (int, int) {
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] != 0 {
			return highcard, i
		}
	}
	return highcard, 0
}

func highCard2(h [16]int) int {
	for i := len(h) - 1; i >= 0; i-- {
		if h[i] != 0 {
			return highcard
		}
	}
	return highcard
}

func returnHandRank(h play_hand) (int, int) {
	if v, points := numerousKind(h.Hand, 5, fivek); v != 0 {
		return v, points
	}
	if v, points := numerousKind(h.Hand, 4, fourk); v != 0 {
		return v, points
	}
	if v, points := fullHouse(h.Hand); v != 0 {
		return v, points
	}
	if v, points := numerousKind(h.Hand, 3, threek); v != 0 {
		return v, points
	}
	if v, points := numerousPair(h.Hand); v != 0 {
		return v, points
	}
	return highCard(h.Hand)
}
func returnHandRankClassic(h [16]int) int {
	if v, _ := numerousKind(h, 5, fivek); v != 0 {
		return v
	}
	if v, _ := numerousKind(h, 4, fourk); v != 0 {
		return v
	}
	if v, _ := fullHouse(h); v != 0 {
		return v
	}
	if v, _ := numerousKind(h, 3, threek); v != 0 {
		return v
	}
	if v, _ := numerousPair(h); v != 0 {
		return v
	}
	return highCard2(h)
}

// KTJJT 220
// QQQJA 483

func rr2(h play_hand, idx int, v *[]int) {
	if idx >= len(h.HandLet) {
		var a [16]int
		vs := h.Hand
		for _, card := range h.HandLet {
			a[card_rank[card]]++
		}
		h.Hand = a
		*v = append(*v, returnHandRankClassic(h.Hand))
		h.Hand = vs
	}
	for i := idx; i < len(h.HandLet); i++ {
		if h.HandLet[i] == "J" {
			for k := range card_rank2 {
				if k == "J" {
					continue
				}
				h.HandLet[i] = k
				rr2(h, i+1, v)
			}
			h.HandLet[i] = "J"
		} else {
			rr2(h, i+1, v)
		}
	}
}

func returnHandRank2(h play_hand) int {
	m := 0
	js := strings.Count(strings.Join(h.HandLet, ""), "J")
	for i := 0; i < js; i++ {
		h.Hand[card_rank2["J"]]--
		for k, val := range card_rank2 {
			if k == "J" {
				continue
			}
			h.Hand[val]++
			if v, _ := numerousKind(h.Hand, 5, fivek); v != 0 {
				m = max(v, m)
			}
			if v, _ := numerousKind(h.Hand, 4, fourk); v != 0 {
				m = max(v, m)
			}
			if v, _ := fullHouse(h.Hand); v != 0 {
				m = max(v, m)
			}
			if v, _ := numerousKind(h.Hand, 3, threek); v != 0 {
				m = max(v, m)
			}
			if v, _ := numerousPair(h.Hand); v != 0 {
				m = max(v, m)
			}
			h.Hand[val]--
		}
		h.Hand[card_rank2["J"]]++
	}
	return max(highCard2(h.Hand), m)
}

func cmprHands(a play_hand, b play_hand) int {
	for i := range a.HandLet {
		if card_rank[a.HandLet[i]] < card_rank[b.HandLet[i]] {
			return -1
		} else if card_rank[a.HandLet[i]] > card_rank[b.HandLet[i]] {
			return 1
		}
	}
	return 0
}

func cmprHands2(a play_hand, b play_hand) int {
	for i := range a.HandLet {
		if card_rank2[a.HandLet[i]] < card_rank2[b.HandLet[i]] {
			return -1
		} else if card_rank2[a.HandLet[i]] > card_rank2[b.HandLet[i]] {
			return 1
		}
	}
	return 0
}

// 1 [K K J J 3] [K K J T T]
func twoPairCompare(a play_hand, b play_hand) int {
	for i := range a.Hand {
		if a.Hand[i] != 2 && a.Hand[i] != 0 {
			a.Hand[i] = 0
		}
	}
	for i := range b.Hand {
		if b.Hand[i] != 2 && b.Hand[i] != 0 {
			b.Hand[i] = 0
		}
	}
	for i := len(a.Hand) - 1; i >= 0; i-- {
		if a.Hand[i] < b.Hand[i] {
			return -1
		} else if a.Hand[i] > b.Hand[i] {
			return 1
		}
	}
	return 0
}

func cmprHandTypes(a play_hand, b play_hand) int {
	if a.HandType == twopair {
		v := twoPairCompare(a, b)
		return v
	}
	if a.Points < b.Points {
		return -1
	} else if a.Points > b.Points {
		return 1
	}
	return 0
}

func c(a, b string) int {
	if card_rank[a] < card_rank[b] {
		return 1
	} else if card_rank[a] > card_rank[b] {
		return -1
	}
	return 0
}

func solve1(lines []string) {
	hands := []play_hand{}
	for _, line := range lines {
		var h play_hand
		var a [16]int
		l := strings.Split(line, " ")
		l1 := strings.Split(l[0], "")
		// slices.SortFunc(l1, c)

		r, _ := strconv.Atoi(l[1])
		h.Bid = r
		h.HandLet = l1
		for _, card := range l1 {
			a[card_rank[card]]++
		}
		h.Hand = a
		hands = append(hands, h)
	}
	for i := range hands {
		t, points := returnHandRank(hands[i])
		hands[i].HandType = t
		hands[i].Points = points
	}
	slices.SortFunc(hands, func(a play_hand, b play_hand) int {
		if a.HandType < b.HandType {
			return -1
		} else if a.HandType > b.HandType {
			return 1
		}
		return cmprHands(a, b)
	})
	var bids int
	for i := range hands {
		hands[i].Rank = i + 1
		bids += hands[i].Rank * hands[i].Bid
	}
	fmt.Println(bids)
}

func solve2(lines []string) {
	hands := []play_hand{}
	for _, line := range lines {
		var h play_hand
		var a [16]int
		l := strings.Split(line, " ")
		l1 := strings.Split(l[0], "")
		// slices.SortFunc(l1, c)

		r, _ := strconv.Atoi(l[1])
		h.Bid = r
		h.HandLet = l1
		for _, card := range l1 {
			a[card_rank2[card]]++
		}
		h.Hand = a
		hands = append(hands, h)
	}
	for i := range hands {
		n := []int{}
		rr2(hands[i], 0, &n)
		hands[i].HandType = slices.Max(n)
	}
	slices.SortFunc(hands, func(a play_hand, b play_hand) int {
		if a.HandType < b.HandType {
			return -1
		} else if a.HandType > b.HandType {
			return 1
		}
		return cmprHands2(a, b)
	})
	var bids int
	for i := range hands {
		hands[i].Rank = i + 1
		bids += hands[i].Rank * hands[i].Bid
		fmt.Println(hands[i].HandLet, i+1)
	}
	fmt.Println(bids)
}

func main() {
	t := common.TestFileOpener()
	solve1(t)
	l := common.FileOpener()
	solve1(l)
	solve2(t)
	solve2(l)
}
