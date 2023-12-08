package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type hand struct {
	cards string
	rank  int
}

func parseLine(s string) hand {
	cards := ""
	rankString := ""
	foundSpace := false
	for _, char := range s {
		if foundSpace {
			rankString += string(char)
		} else if char == ' ' {
			foundSpace = true
		} else {
			cards += string(char)
		}
	}
	rank, _ := strconv.Atoi(rankString)
	return hand{cards: cards, rank: rank}
}

func getHandMap(h hand) map[rune]int {
	m := make(map[rune]int)
	for _, char := range h.cards {
		m[char] += 1
	}
	return m
}

func getHandVal(h hand, handleJokers bool) int {
	hMap := getHandMap(h)
	maxCount := 0
	var valHighCount rune
	for k, v := range hMap {
		if v > maxCount && (k != 'J' || !handleJokers) {
			maxCount = v
			valHighCount = k
		}
	}
	if handleJokers {
		if hMap['J'] != 0 {
			hMap[valHighCount] += hMap['J']
			hMap['J'] = 0
		}
		for k, v := range hMap {
			if v > maxCount {
				maxCount = v
				valHighCount = k
			}
		}
	}
	if maxCount == 5 {
		// five of a kind
		return 7
	} else if maxCount == 4 {
		// four of a kind
		return 6
	} else if maxCount == 1 {
		// high card
		return 1
	}
	secondHighest := 0
	for k, v := range hMap {
		if k != valHighCount && v > secondHighest {
			secondHighest = v
		}
	}
	if maxCount == 3 && secondHighest == 2 {
		// full house
		return 5
	} else if maxCount == 3 && secondHighest == 1 {
		// three of a kind
		return 4
	} else if maxCount == 2 && secondHighest == 2 {
		// two pair
		return 3
	} else if maxCount == 2 && secondHighest == 1 {
		return 2
	}
	log.Fatal("invalid hand", h, "with map", hMap)
	return 0
}

func convertCardToInt(c rune) int {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 0
	case 'T':
		return 10
	default:
		return int(c - '0')
	}
}

func sortHandsForJokers(handleJokers bool) func(a hand, b hand) int {
	return func(a hand, b hand) int {
		return sortHands(a, b, handleJokers)
	}
}

func sortHands(a hand, b hand, handleJokers bool) int {
	aVal := getHandVal(a, handleJokers)
	bVal := getHandVal(b, handleJokers)
	if aVal == bVal {
		for idx := range a.cards {
			if a.cards[idx] != b.cards[idx] {
				return convertCardToInt([]rune(a.cards)[idx]) - convertCardToInt([]rune(b.cards)[idx])
			}
		}
		return 0
	} else {
		return aVal - bVal
	}

}

func solvePart(lines []string, handleJokers bool) {
	var hands []hand
	for _, line := range lines {
		hands = append(hands, parseLine(line))
	}
	slices.SortFunc(hands, sortHandsForJokers(handleJokers))
	var num int
	res := 0
	for idx, hand := range hands {
		fmt.Println("cards ", hand.cards, "at rank ", idx+1, " w bid ", hand.rank)
		num = hand.rank * (idx + 1)
		fmt.Println("results incremented by ", num)
		res += num
	}
	fmt.Println("Result: ", res)
}

func main() {
	// thanks https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	solvePart(lines, false)
	solvePart(lines, true)
}
