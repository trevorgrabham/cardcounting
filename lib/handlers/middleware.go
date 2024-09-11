package handlers

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"

	"github.com/gorilla/securecookie"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/strategy"
)

var s *securecookie.SecureCookie
var nextUserID *atomic.Int64

func init() {
	hashString := os.Getenv("SCHASH")
	blockString := os.Getenv("SCBLOCK")
	hash, err := base64.StdEncoding.DecodeString(hashString)
	if err != nil { panic(err) }
	block, err := base64.StdEncoding.DecodeString(blockString)
	if err != nil { panic(err) }
	s = securecookie.New(hash, block)
}

func NewMiddleWare(opts ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for _, opt := range opts {
			next = opt(next)
		}
		return next 
	}
}

func WithCookies(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userID int64
		if cookie, err := r.Cookie("user-id"); err != nil {
			for {
				userID = nextUserID.Load()
				if nextUserID.CompareAndSwap(userID, userID+1) { break }
			}
			encoded, err := s.Encode("user-id", userID)
			if err != nil { panic(err) }
			cookie = &http.Cookie{
				Name: "user-id",
				Value: encoded, 
				Path: "/"}
			http.SetCookie(w, cookie)
		} else if err = s.Decode("user-id", cookie.Value, &userID); err != nil { panic(err) }
		ctx := context.WithValue(r.Context(), lib.ContextKey("user-id"), userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WithHands(next http.Handler) http.Handler {
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