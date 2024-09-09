package html

import (
	"html/template"
	"os"

	"github.com/trevorgrabham/cardcounting/cardcounting/lib/cards"
)

var IncludeFiles = map[string][]string{
	"index": {"html/index.html"},
}

type IndexData struct {
}

func NewIndexData() IndexData {
	return IndexData{}
}

type CardData struct {
	SVG template.HTMLAttr
	Value int
}

func NewCardData(card cards.Card) (CardData, error) {
	svg, err := os.ReadFile(card.ImgPath)
	if err != nil { return CardData{}, err }
	value := cards.ConvertToValue(card.Rank)
	return CardData{ SVG: template.HTMLAttr(svg), Value: value }, nil 
}

type CountValidationData struct {
	Correct bool 
	Count int16
}

func NewCountValidationData(correct bool, count int16) CountValidationData {
	return CountValidationData{
		Correct: correct,
		Count: count}
}