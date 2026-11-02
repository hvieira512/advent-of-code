package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hvieira512/aoc2023/utils"
)

// a hand has 5 cards and the bid
type Hand struct {
	cards string
	bid   int
	rank  int
}

func getHands(lines []string) []Hand {
	hands := []Hand{}

	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		strAux := strings.Split(line, " ")

		cards := strAux[0]
		bid, _ := strconv.Atoi(strAux[1])

		hands = append(hands, Hand{cards, bid, 0})
	}

	return hands
}

func getCardsCharNum(hand Hand) map[string]int {
	cardsCharMap := map[string]int{}
	for j := range hand.cards {
		cardsCharStr := string(hand.cards[j])

		charExists := cardsCharMap[cardsCharStr]
		if charExists == 0 {
			cardsCharMap[cardsCharStr] = 1
		} else {
			value := cardsCharMap[cardsCharStr]
			cardsCharMap[cardsCharStr] = value + 1
		}
	}

	return cardsCharMap
}

func getPartOne(hands []Hand) int {
	result := 0

	highCards, onePairCards, twoPairCards, threeKindCards, fullHouseCards, fourKindCards, fiveKindCards := []Hand{}, []Hand{}, []Hand{}, []Hand{}, []Hand{}, []Hand{}, []Hand{}
	for _, hand := range hands {
		cardsCharNum := getCardsCharNum(hand)

		// distribute all cards to their types
		if isHighCard(cardsCharNum) {
			highCards = append(highCards, hand)
			continue
		} else if isOnePairCard(cardsCharNum) {
			onePairCards = append(onePairCards, hand)
			continue
		} else if isTwoPairCard(cardsCharNum) {
			twoPairCards = append(twoPairCards, hand)
			continue
		} else if isThreeKindCard(cardsCharNum) {
			threeKindCards = append(threeKindCards, hand)
			continue
		} else if isFullHouseCard(cardsCharNum) {
			fullHouseCards = append(fullHouseCards, hand)
			continue
		} else if isFourKindCard(cardsCharNum) {
			fourKindCards = append(fourKindCards, hand)
			continue
		} else if isFiveKindCard(cardsCharNum) {
			fiveKindCards = append(fiveKindCards, hand)
			continue
		}

		// append all to the main one
		hands = append(hands)
	}

	return result
}

func isHighCard(cardCharsNum map[string]int) bool {
	return len(cardCharsNum) == 5
}

func isOnePairCard(cardsCharNum map[string]int) bool {
	panic("unimplemented")
}

func isTwoPairCard(cardsCharNum map[string]int) bool {
	panic("unimplemented")
}

func isThreeKindCard(cardsCharNum map[string]int) bool {
	panic("unimplemented")
}

func isFullHouseCard(cardsCharNum map[string]int) bool {
	panic("unimplemented")
}

func isFourKindCard(cardsCharNum map[string]int) bool {
	panic("unimplemented")
}

func isFiveKindCard(cardCharsNum map[string]int) bool {
	return len(cardCharsNum) == 1
}

func main() {
	lines := utils.GetLinesFile("example.txt")
	hands := getHands(lines)

	fmt.Printf("Part 1: %d\n", getPartOne(hands))
}
