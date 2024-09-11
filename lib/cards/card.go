package cards

import (
	"fmt"
	"math/rand/v2"
	"os"
)

type Card struct {
	Rank    CardRank
	Value   uint8
	SVG 		[]byte
}

func (c *Card) GetSVG() (err error) {
	c.SVG, err = os.ReadFile(CardImgPaths[c.Rank])
	if err != nil { return fmt.Errorf("getsvg from %s: %v", CardImgPaths[c.Rank], err) }
	return nil 
}

func NewCard(c CardRank) (card Card) {
	res := Card{Rank: c, Value: ConvertToValue(c)}
	res.GetSVG()
	return res
}

type Deck []CardRank

func Sum(cards []Card) uint8 {
	var sum uint8
	for _, card := range cards {
		sum += card.Value
	}
	return sum
}

func NewDeck(numDecks int) (deck Deck) {
	for i := range numDecks * 52 {
		deck = append(deck, CardRank(i%52+2))
	}
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return
}

func ConvertToValue(rank CardRank) uint8 {
	switch {
	case rank == Unknown:
		return 0
	case rank >= TwoOfClubs && rank <= TwoOfHearts:
		return 2
	case rank >= ThreeOfClubs && rank <= ThreeOfHearts:
		return 3
	case rank >= FourOfClubs && rank <= FourOfHearts:
		return 4
	case rank >= FiveOfClubs && rank <= FiveOfHearts:
		return 5
	case rank >= SixOfClubs && rank <= SixOfHearts:
		return 6
	case rank >= SevenOfClubs && rank <= SevenOfHearts:
		return 7
	case rank >= EightOfClubs && rank <= EightOfHearts:
		return 8
	case rank >= NineOfClubs && rank <= NineOfHearts:
		return 9
	case rank >= TenOfClubs && rank <= KingOfHearts:
		return 10
	case rank >= AceOfClubs && rank <= AceOfHearts:
		return 1
	}
	return 255
}

type CardRank uint8

const (
	Nil CardRank = iota
	Unknown 
	TwoOfClubs
	TwoOfDiamonds
	TwoOfSpades
	TwoOfHearts
	ThreeOfClubs
	ThreeOfDiamonds
	ThreeOfSpades
	ThreeOfHearts
	FourOfClubs
	FourOfDiamonds
	FourOfSpades
	FourOfHearts
	FiveOfClubs
	FiveOfDiamonds
	FiveOfSpades
	FiveOfHearts
	SixOfClubs
	SixOfDiamonds
	SixOfSpades
	SixOfHearts
	SevenOfClubs
	SevenOfDiamonds
	SevenOfSpades
	SevenOfHearts
	EightOfClubs
	EightOfDiamonds
	EightOfSpades
	EightOfHearts
	NineOfClubs
	NineOfDiamonds
	NineOfSpades
	NineOfHearts
	TenOfClubs
	TenOfDiamonds
	TenOfSpades
	TenOfHearts
	JackOfClubs
	JackOfDiamonds
	JackOfSpades
	JackOfHearts
	QueenOfClubs
	QueenOfDiamonds
	QueenOfSpades
	QueenOfHearts
	KingOfClubs
	KingOfDiamonds
	KingOfSpades
	KingOfHearts
	AceOfClubs
	AceOfDiamonds
	AceOfSpades
	AceOfHearts
)

var CardImgPaths = map[CardRank]string{
	Unknown:         "html/svgs/card_back.svg",
	TwoOfClubs:      "html/svgs/2_of_clubs.svg",
	TwoOfDiamonds:   "html/svgs/2_of_diamonds.svg",
	TwoOfSpades:     "html/svgs/2_of_spades.svg",
	TwoOfHearts:     "html/svgs/2_of_hearts.svg",
	ThreeOfClubs:    "html/svgs/3_of_clubs.svg",
	ThreeOfDiamonds: "html/svgs/3_of_diamonds.svg",
	ThreeOfSpades:   "html/svgs/3_of_spades.svg",
	ThreeOfHearts:   "html/svgs/3_of_hearts.svg",
	FourOfClubs:     "html/svgs/4_of_clubs.svg",
	FourOfDiamonds:  "html/svgs/4_of_diamonds.svg",
	FourOfSpades:    "html/svgs/4_of_spades.svg",
	FourOfHearts:    "html/svgs/4_of_hearts.svg",
	FiveOfClubs:     "html/svgs/5_of_clubs.svg",
	FiveOfDiamonds:  "html/svgs/5_of_diamonds.svg",
	FiveOfSpades:    "html/svgs/5_of_spades.svg",
	FiveOfHearts:    "html/svgs/5_of_hearts.svg",
	SixOfClubs:      "html/svgs/6_of_clubs.svg",
	SixOfDiamonds:   "html/svgs/6_of_diamonds.svg",
	SixOfSpades:     "html/svgs/6_of_spades.svg",
	SixOfHearts:     "html/svgs/6_of_hearts.svg",
	SevenOfClubs:    "html/svgs/7_of_clubs.svg",
	SevenOfDiamonds: "html/svgs/7_of_diamonds.svg",
	SevenOfSpades:   "html/svgs/7_of_spades.svg",
	SevenOfHearts:   "html/svgs/7_of_hearts.svg",
	EightOfClubs:    "html/svgs/8_of_clubs.svg",
	EightOfDiamonds: "html/svgs/8_of_diamonds.svg",
	EightOfSpades:   "html/svgs/8_of_spades.svg",
	EightOfHearts:   "html/svgs/8_of_hearts.svg",
	NineOfClubs:     "html/svgs/9_of_clubs.svg",
	NineOfDiamonds:  "html/svgs/9_of_diamonds.svg",
	NineOfSpades:    "html/svgs/9_of_spades.svg",
	NineOfHearts:    "html/svgs/9_of_hearts.svg",
	TenOfClubs:      "html/svgs/10_of_clubs.svg",
	TenOfDiamonds:   "html/svgs/10_of_diamonds.svg",
	TenOfSpades:     "html/svgs/10_of_spades.svg",
	TenOfHearts:     "html/svgs/10_of_hearts.svg",
	JackOfClubs:     "html/svgs/jack_of_clubs.svg",
	JackOfDiamonds:  "html/svgs/jack_of_diamonds.svg",
	JackOfSpades:    "html/svgs/jack_of_spades.svg",
	JackOfHearts:    "html/svgs/jack_of_hearts.svg",
	QueenOfClubs:    "html/svgs/queen_of_clubs.svg",
	QueenOfDiamonds: "html/svgs/queen_of_diamonds.svg",
	QueenOfSpades:   "html/svgs/queen_of_spades.svg",
	QueenOfHearts:   "html/svgs/queen_of_hearts.svg",
	KingOfClubs:     "html/svgs/king_of_clubs.svg",
	KingOfDiamonds:  "html/svgs/king_of_diamonds.svg",
	KingOfSpades:    "html/svgs/king_of_spades.svg",
	KingOfHearts:    "html/svgs/king_of_hearts.svg",
	AceOfClubs:      "html/svgs/ace_of_clubs.svg",
	AceOfDiamonds:   "html/svgs/ace_of_diamonds.svg",
	AceOfSpades:     "html/svgs/ace_of_spades.svg",
	AceOfHearts:     "html/svgs/ace_of_hearts.svg",
}