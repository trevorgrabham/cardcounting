package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/trevorgrabham/cardcounting/cardcounting/html"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/cards"
)

func HandleDealCard(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user-id").(int64)
	if !ok { panic(fmt.Errorf("couldn't parse 'user-id'")) }
	card, finished := lib.UserData.Draw(userID)
	if finished {
		finishedMessage := template.Must(template.New("finished-message").ParseFiles(html.IncludeFiles["finished-message"]...))
		if err := finishedMessage.Execute(w, nil); err != nil { panic(err) }
	}
	c := template.Must(template.New("card").ParseFiles(html.IncludeFiles["card"]...))
	if err := c.Execute(w, cards.NewCard(card)); err != nil { panic(err) }
}