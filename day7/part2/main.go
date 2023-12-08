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
  typeStrength int
}

func main() {
  cardStrings, bids := parseFile("../input.txt")
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
  jokers := 0

  for i:=0; i<len(labels);i++{
    if labels[i] == 'J' {
      jokers += 1
    }
    cardExists := false;
    for j, card := range cards {
      if labels[i] == card.label {
        cards[j].amount += 1
        cardExists = true
        break
      }
    }
    if cardExists {
      continue
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

  sort.Slice(cards, func(i, j int ) bool { return compareCards(cards, i, j)})

  return Hand {
    cards: cards,
    bid: bid,
    labels: []byte(cardString),
    typeStrength: getTypeStrength(cards, jokers),
  }
}

func cardStrength(label byte) int {
  labelOrder := []byte{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
  for i, el := range labelOrder {
    if el == label {
      return i + 1
    }
  }
  return -1
}

func compareHands(hands []Hand, i int, j int) bool {
  typediff := hands[i].typeStrength != hands[j].typeStrength
  if typediff {
    return hands[i].typeStrength < hands[j].typeStrength
  }

  for k:=0; k<len(hands[i].labels); k++ {
    if cardStrength(hands[i].labels[k]) == cardStrength(hands[j].labels[k]) {
      continue
    }
    return cardStrength(hands[i].labels[k]) < cardStrength(hands[j].labels[k])
  }
  return false
}

func compareCards(cards []Card, i int, j int) bool {
  if cards[i].amount != cards[j].amount {
    return cards[i].amount > cards[j].amount
  }
  return cards[i].strength > cards[j].strength
}

func getTypeStrength(cards []Card, jokers int) int {
  if jokers == 0 {
    switch cards[0].amount {
      case 4, 5: return cards[0].amount + 1
      case 3: return 2 + cards[1].amount
      case 2: return cards[1].amount
      case 1: return 0
    }
  }
  switch jokers {
    case 4, 5: return 6
    case 3: return 4 + cards[1].amount
    case 2: {
      if cards[1].amount == 1 {
        return 3
      }
      return 3 + cards[0].amount
    }
    case 1: {
      if cards[0].amount == 1 {
        return 1
      }
      if cards[0].amount == 2 {
        return 2 + cards[1].amount
      }
      return 2 + cards[0].amount
    }
    default: return -1
  }
}
