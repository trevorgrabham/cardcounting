package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/trevorgrabham/cardcounting/cardcounting/html"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/cards"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/strategy"
)

var trainingButtons = html.TrainingData{
	Split: html.ButtonData{
		Button: html.NewElementData(
			html.WithID("split-button"),
			html.WithClasses("training-button"),
			html.WithHTMX("get", "/training/split", "target", "#message-box")),
		Text: "Split"},
	Stand: html.ButtonData{
		Button: html.NewElementData(
			html.WithID("stand-button"),
			html.WithClasses("training-button"),
			html.WithHTMX("get", "/training/stand", "target", "#message-box")),
		Text: "Stand"},
	Hit: html.ButtonData{
		Button: html.NewElementData(
			html.WithID("hit-button"),
			html.WithClasses("training-button"),
			html.WithHTMX("get", "/training/hit", "target", "#message-box")),
		Text: "Hit"},
	Double: html.ButtonData{
		Button: html.NewElementData(
			html.WithID("double-button"),
			html.WithClasses("training-button"),
			html.WithHTMX("get", "/training/double", "target", "#message-box")),
		Text: "Double"}}

var trainingMessage = html.TrainingMessageData{
	ButtonData: html.ButtonData{
		Button: html.NewElementData(
			html.WithID("close-message"),
			html.WithClasses("training-message-button"),
			html.WithHyperscript(`on click add .hidden to #message-box`)),
		Text: "Close"}}

func HandleStartTraining(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64) 
	if !ok { panic(fmt.Errorf("couldn't parse out user-id")) }

	if err := r.ParseForm(); err != nil { panic(err) }
	res := r.Form.Get("decks")
	var errors []string
	if res == "" {
		errors = append(errors, lib.ErrorStrings["training-settings-deck-string"])
	}

	numDecks, err := strconv.Atoi(res)
	if err != nil { panic(err) }
	if numDecks < 1 || numDecks > 8 {
		errors = append(errors, lib.ErrorStrings["training-settings-deck-string"])
	}

	var dealer, player []cards.Card
	deck := cards.NewDeck(numDecks)
	player = append(player, cards.NewCard(deck[0]))
	dealer = append(dealer, cards.NewCard(deck[1]))
	player = append(player, cards.NewCard(deck[2]))
	lib.UserData.AddDeck(userID, deck[3:])

	data := trainingButtons
	data.Dealer = html.HandData{
		Cards: dealer,
		Value: dealer[0].Value}
	data.Player = html.HandData{
		Cards: player,
		Value: cards.Sum(player)}

	training := template.Must(template.New("training").ParseFiles(html.IncludeFiles["training"]...))
	if err = training.Execute(w, data); err != nil { panic(err) }
}

func HandleHit(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("handlehit(): couldn't parse 'user-id'")) }
	hand, ok := r.Context().Value("hand").(strategy.Hand)
	if !ok { panic(fmt.Errorf("handlehit(): couldn't parse 'hand'"))}
	properPlay, err := hand.Strategy()
	if err != nil { panic(fmt.Errorf("handlehit(): %v", err)) }

	data := trainingMessage

	message := template.Must(template.New("training-message").ParseFiles(html.IncludeFiles["training-message"]...))
	
	if properPlay == strategy.Hit {
		data.Message = fmt.Sprintf(lib.ErrorStrings["training-correct-strategy-format"], properPlay)
		if err := message.Execute(w, data); err != nil { panic(fmt.Errorf("handlehit(): %v", err)) }
		return
	}
	data.Message = fmt.Sprintf(lib.ErrorStrings["training-wrong-strategy-format"], properPlay)
	lib.UserData.IncStrategyErrors(userID)
	if err := message.Execute(w, data); err != nil { panic(fmt.Errorf("handlehit(): %v", err)) }
}

func HandleStand(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("handlestand(): couldn't parse 'user-id'")) }
	hand, ok := r.Context().Value("hand").(strategy.Hand)
	if !ok { panic(fmt.Errorf("handlestand(): couldn't parse 'hand'")) }

	properPlay, err := hand.Strategy()
	if err != nil { panic(fmt.Errorf("handlestand(): %v", err)) }

	data := trainingMessage
	message := template.Must(template.New("practice-response").ParseFiles(html.IncludeFiles["practice-response"]...))

	if properPlay == strategy.Stand {
		data.Message = fmt.Sprintf(lib.ErrorStrings["training-correct-strategy-format"], properPlay)
		if err := message.Execute(w, data); err != nil { panic(fmt.Errorf("handlestand(): %v", err)) }
		return
	}
	data.Message = fmt.Sprintf(lib.ErrorStrings["training-correct-strategy-format"], properPlay)
	lib.UserData.IncStrategyErrors(userID)
	if err := message.Execute(w, data); err != nil { panic(fmt.Errorf("handlestand(): %v", err)) }
}

func HandleDouble(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("handledouble(): couldn't parse 'user-id'")) }
	hand, ok := r.Context().Value("hand").(strategy.Hand)
	if !ok { panic(fmt.Errorf("handledouble(): couldn't parse 'hand'")) }

	properPlay, err := hand.Strategy()
	if err != nil { panic(fmt.Errorf("handledouble(): %v", err)) }

	data := trainingMessage
	message := template.Must(template.New("practice-response").ParseFiles(html.IncludeFiles["practice-response"]...))

	if properPlay == strategy.Double {
		data.Message = fmt.Sprintf(lib.ErrorStrings["training-correct-strategy-format"], properPlay)
		if err := message.Execute(w, data); err != nil { panic(fmt.Errorf("handledouble(): %v", err)) }
		return
	}
	data.Message = fmt.Sprintf(lib.ErrorStrings["training-wrong-strategy-format"], properPlay)
	lib.UserData.IncStrategyErrors(userID)
	if err := message.Execute(w, data); err != nil { panic(fmt.Errorf("handledouble(): %v", err)) }
}

func HandleSplit(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("handlesplit(): couldn't parse 'user-id'")) }
	hand, ok := r.Context().Value("hand").(strategy.Hand)
	if !ok { panic(fmt.Errorf("handlesplit(): couldn't parse 'hand'")) }

	properPlay, err := hand.Strategy()
	if err != nil { panic(fmt.Errorf("handlesplit(): %v", err)) }

	data := trainingMessage
	message := template.Must(template.New("practice-response").ParseFiles(html.IncludeFiles["practice-response"]...))
	if properPlay == strategy.Split {
		data.Message = fmt.Sprintf(lib.ErrorStrings["training-correct-strategy-format"], properPlay)
		if err := message.Execute(w, data); err != nil { panic(fmt.Errorf("handlesplit(): %v", err)) }
		return
	}
	data.Message = fmt.Sprintf(lib.ErrorStrings["training-correct-strategy-format"], properPlay)
	lib.UserData.IncStrategyErrors(userID)
	if err := message.Execute(w, data); err != nil { panic(fmt.Errorf("handlesplit(): %v", err)) }
}

func HandleCountCheck(w http.ResponseWriter, r *http.Request) {

}