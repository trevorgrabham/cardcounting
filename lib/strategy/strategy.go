package strategy

import "fmt"

type Option uint8
type Total uint8
type Hand struct {
	Dealer Total
	Player Total
}

func (h Hand) Strategy() (Option, error) {
	switch {
	case h.Player <= Seven:
		return Hit, nil
	case h.Player == Eight:
		if h.Dealer == Five || h.Dealer == Six {
			return Double, nil
		}
		return Hit, nil
	case h.Player == Nine:
		if h.Dealer <= Six {
			return Double, nil
		}
		return Hit, nil
	case h.Player == Ten:
		if h.Dealer <= Nine {
			return Double, nil
		}
		return Hit, nil
	case h.Player == Eleven:
		return Double, nil
	case h.Player == Twelve:
		if h.Dealer >= Four && h.Dealer <= Six {
			return Stand, nil
		}
		return Hit, nil
	case h.Player >= Thirteen && h.Player <= Sixteen:
		if h.Dealer <= Six {
			return Stand, nil
		}
		return Hit, nil
	case h.Player >= Seventeen && h.Player <= Twentyone:
		return Stand, nil
	case h.Player >= SoftThirteen && h.Player <= SoftSixteen:
		if h.Dealer >= Four && h.Dealer <= Six {
			return Double, nil
		}
		return Hit, nil
	case h.Player == SoftSeventeen:
		if h.Dealer <= Six {
			return Double, nil
		}
		return Hit, nil
	case h.Player == SoftEighteen:
		if h.Dealer == Nine || h.Dealer == Ten {
			return Hit, nil
		}
		if h.Dealer >= Three && h.Dealer <= Six {
			return Double, nil
		}
		return Stand, nil
	case h.Player == SoftNineteen:
		if h.Dealer == Four || h.Dealer == Five {
			return Double, nil
		}
		return Stand, nil
	case h.Player == SoftTwenty:
		return Stand, nil
	case h.Player == DoubleTwos:
		if h.Dealer <= Eight {
			return Split, nil
		}
		return Hit, nil
	case h.Player == DoubleThrees:
		if h.Dealer >= Four && h.Dealer <= Seven {
			return Split, nil
		}
		return Hit, nil
	case h.Player == DoubleFours:
		if h.Dealer == Five || h.Dealer == Six {
			return Double, nil
		}
		return Hit, nil
	case h.Player == DoubleFives:
		if h.Dealer >= Ten {
			return Hit, nil
		}
		return Double, nil
	case h.Player == DoubleSixes:
		if h.Dealer <= Six {
			return Split, nil
		}
		return Hit, nil
	case h.Player == DoubleSevens:
		if h.Dealer <= Seven {
			return Split, nil
		}
		if h.Dealer == Ten {
			return Stand, nil
		}
		return Hit, nil
	case h.Player == DoubleEights:
		return Split, nil
	case h.Player == DoubleNines:
		if h.Dealer == Seven || h.Dealer >= Ten {
			return Stand, nil
		}
		return Split, nil
	case h.Player == DoubleTens:
		return Stand, nil
	case h.Player == DoubleAces:
		return Split, nil
	}
	return Surrender, fmt.Errorf("unknown match for hand %v", h)
}

const (
	Hit Option = iota
	Stand
	Double
	Split
	Surrender
)

func (o Option) String() string {
	switch o {
	case Hit:
		return "hit"
	case Stand:
		return "stand"
	case Double: 
		return "double"
	case Split:
		return "split"
	case Surrender:
		return "surrender"
	}
	return ""
}

const (
	Two Total = iota
	Three
	// User possible values start here
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Eleven
	// Dealer possible values end here
	Twelve
	Thirteen
	Fourteen
	Fifteen
	Sixteen
	Seventeen
	Eighteen
	Nineteen
	Twenty
	Twentyone
	SoftThirteen
	SoftFourteen
	SoftFifteen
	SoftSixteen
	SoftSeventeen
	SoftEighteen
	SoftNineteen
	SoftTwenty
	DoubleTwos
	DoubleThrees
	DoubleFours
	DoubleFives
	DoubleSixes
	DoubleSevens
	DoubleEights
	DoubleNines
	DoubleTens
	DoubleAces
	Blackjacks
)

var ValueToTotal = map[int]Total{
	1:  Eleven,
	2:  Two,
	3:  Three,
	4:  Four,
	5:  Five,
	6:  Six,
	7:  Seven,
	8:  Eight,
	9:  Nine,
	10: Ten,
	11: Eleven,
	12: Twelve,
	13: Thirteen,
	14: Fourteen,
	15: Fifteen,
	16: Sixteen,
	17: Seventeen,
	18: Eighteen,
	19: Nineteen,
	20: Twenty,
	21: Twentyone,
}