package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/trevorgrabham/cardcounting/cardcounting/html"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/cards"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/strategy"
)

const (
	CorrectStrategyString = `Correct! The proper strategy was to %s!`
	WrongStrategyString = `Incorrect! The correct strategy was to %s!`
	CorrectCountString = `Correct! The count is %d!`
	WrongCountString = `Incorrect! The current count is %d!`
)

func HandleStartPractice(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64) 
	if !ok { panic(fmt.Errorf("couldn't parse out user-id")) }
	if err := r.ParseForm(); err != nil { panic(err) }
	res := r.Form.Get("decks")
	if res == "" { panic(fmt.Errorf("no 'decks' parameter provided to /startPractice")) }
	numDecks, err := strconv.Atoi(res)
	if err != nil { panic(err) }
	if numDecks < 1 || numDecks > 8 { panic(fmt.Errorf("got: decks(%d)\twant: 0 < decks < 9", numDecks))}
	deck := cards.NewDeck(numDecks)
	lib.UserData.AddDeck(userID, deck)
	practice := template.Must(template.New("practice").ParseFiles(html.IncludeFiles["practice"]...))
	if err = practice.Execute(w, nil); err != nil { panic(err) }
}

func SetHandContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil { panic(fmt.Errorf("sethandcontext(): unable to parse form")) }
		playerValueString := r.Form.Get("player")
		if playerValueString == "" { panic(fmt.Errorf("sethandcontext(): no 'player' value provided"))}
		dealerValueString := r.Form.Get("dealer")
		if playerValueString == "" { panic(fmt.Errorf("sethandcontext(): no 'dealer' value provided"))}
		playerValue, err := strconv.Atoi(playerValueString)
		if err != nil { panic(fmt.Errorf("sethandcontext(): %v", err))}
		dealerValue, err := strconv.Atoi(dealerValueString)
		if err != nil { panic(fmt.Errorf("sethandcontext(): %v", err))}
		hand := strategy.Hand{
			Player: strategy.ValueToTotal[playerValue],
			Dealer: strategy.ValueToTotal[dealerValue]}
		ctx := context.WithValue(r.Context(), lib.ContextKey("hand"), hand)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HandleHit(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("handlehit(): couldn't parse 'user-id'")) }
	hand, ok := r.Context().Value("hand").(strategy.Hand)
	if !ok { panic(fmt.Errorf("handlehit(): couldn't parse 'hand'"))}
	properPlay, err := hand.Strategy()
	if err != nil { panic(fmt.Errorf("handlehit(): %v", err)) }
	practiceResponse := template.Must(template.New("practice-response").ParseFiles(html.IncludeFiles["practice-response"]...))
	if properPlay == strategy.Hit {
		if err := practiceResponse.Execute(w, fmt.Sprintf(CorrectStrategyString, "hit")); err != nil { panic(fmt.Errorf("handlehit(): %v", err)) }
		return
	}
	lib.UserData.IncStrategyErrors(userID)
	if err := practiceResponse.Execute(w, fmt.Sprintf(WrongStrategyString, properPlay)); err != nil { panic(fmt.Errorf("handlehit(): %v", err)) }
}

func HandleStand(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("handlestand(): couldn't parse 'user-id'")) }
	hand, ok := r.Context().Value("hand").(strategy.Hand)
	if !ok { panic(fmt.Errorf("handlestand(): couldn't parse 'hand'")) }
	properPlay, err := hand.Strategy()
	if err != nil { panic(fmt.Errorf("handlestand(): %v", err)) }
	practiceResponse := template.Must(template.New("practice-response").ParseFiles(html.IncludeFiles["practice-response"]...))
	if properPlay == strategy.Stand {
		if err := practiceResponse.Execute(w, fmt.Sprintf(CorrectStrategyString, "stand")); err != nil { panic(fmt.Errorf("handlestand(): %v", err)) }
		return
	}
	lib.UserData.IncStrategyErrors(userID)
	if err := practiceResponse.Execute(w, fmt.Sprintf(WrongStrategyString, properPlay)); err != nil { panic(fmt.Errorf("handlestand(): %v", err)) }
}

func HandleDouble(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("handledouble(): couldn't parse 'user-id'")) }
	hand, ok := r.Context().Value("hand").(strategy.Hand)
	if !ok { panic(fmt.Errorf("handledouble(): couldn't parse 'hand'")) }
	properPlay, err := hand.Strategy()
	if err != nil { panic(fmt.Errorf("handledouble(): %v", err)) }
	practiceResponse := template.Must(template.New("practice-response").ParseFiles(html.IncludeFiles["practice-response"]...))
	if properPlay == strategy.Double {
		if err := practiceResponse.Execute(w, fmt.Sprintf(CorrectStrategyString, "double")); err != nil { panic(fmt.Errorf("handledouble(): %v", err)) }
		return
	}
	lib.UserData.IncStrategyErrors(userID)
	if err := practiceResponse.Execute(w, fmt.Sprintf(WrongStrategyString, properPlay)); err != nil { panic(fmt.Errorf("handledouble(): %v", err)) }
}

func HandleSplit(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("handlesplit(): couldn't parse 'user-id'")) }
	hand, ok := r.Context().Value("hand").(strategy.Hand)
	if !ok { panic(fmt.Errorf("handlesplit(): couldn't parse 'hand'")) }
	properPlay, err := hand.Strategy()
	if err != nil { panic(fmt.Errorf("handlesplit(): %v", err)) }
	practiceResponse := template.Must(template.New("practice-response").ParseFiles(html.IncludeFiles["practice-response"]...))
	if properPlay == strategy.Split {
		if err := practiceResponse.Execute(w, fmt.Sprintf(CorrectStrategyString, "split")); err != nil { panic(fmt.Errorf("handlesplit(): %v", err)) }
		return
	}
	lib.UserData.IncStrategyErrors(userID)
	if err := practiceResponse.Execute(w, fmt.Sprintf(WrongStrategyString, properPlay)); err != nil { panic(fmt.Errorf("handlesplit(): %v", err)) }
}

func HandleCountCheck(w http.ResponseWriter, r *http.Request) {

}