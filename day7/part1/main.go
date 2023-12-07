package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
  label byte
  amount int
  strength int
}

type Hand struct {
  labels []byte
  cards []Card
  bid int
}

func main() {
  cardStrings, bids := parseFile("example.txt")
  hands := make([]Hand, len(cardStrings))
  for i:=0; i<len(cardStrings); i++ {
    hands[i] = getHand(cardStrings[i], bids[i])
  }
  sort.Slice(hands, func(i, j int) bool{ return compareHands(hands, i,j)})

  res := 0
  for i, hand := range hands {
    res += (i + 1) * hand.bid
  }
  fmt.Println(res)
}

// takes a file path and returns the cards and bids in that file
func parseFile(filePath string) ([]string, []string) {
  var cards []string
  var bids []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
    card, bid, _ := strings.Cut(line, " ")
    cards = append(cards, card)
    bids = append(bids, bid)
	}
	return cards, bids
}

// takes card string and returns a representation of the card as an array of CardSymbols
func getHand(cardString string, bidString string) Hand {
  labels := []byte(cardString)
  var cards []Card
  for i:=0; i<len(labels);i++{
    for j, card := range cards {
      if labels[i] == card.label {
        cards[j].amount += 1
        break
      }
    }
    newCard := Card {
      label: labels[i],
      amount: 1,
      strength: cardStrength(labels[i]),
    }
    cards = append(cards, newCard)
  }
  bid, err := strconv.Atoi(bidString)
  if err != nil {
    log.Fatalln(err)
  }
  sort.Slice(cards, func(i, j int ) bool { return cards[i].amount > cards[j].amount})

  return Hand {
    cards: cards,
    bid: bid,
    labels: []byte(cardString),
  }
}

func cardStrength(label byte) int {
  labelOrder := []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
  for i, el := range labelOrder {
    if el == label {
      return i + 1
    }
  }
  return -1
}

func compareHands(hands []Hand, i int, j int) bool {
  cards1 := hands[i].cards
  cards2 := hands[j].cards
  if cards1[0].amount != cards2[0].amount {
    return cards1[0].amount < cards2[0].amount
  }
  if cards1[0].amount + cards1[1].amount != cards2[0].amount + cards2[1].amount {
    return cards1[0].amount + cards1[1].amount < cards2[0].amount + cards2[1].amount
  }
  for k:=0; k<len(hands[i].labels); k++ {
    if cardStrength(hands[i].labels[k]) == cardStrength(hands[j].labels[k]) {
      continue
    }
    return cardStrength(hands[i].labels[k]) < cardStrength(hands[j].labels[k])
  }
  return false
}
