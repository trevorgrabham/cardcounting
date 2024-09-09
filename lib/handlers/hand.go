package handlers

import "net/http"

func HandleHand(w http.ResponseWriter, r *http.Request) {
	card := cards.DrawCard()
}