package handlers

import (
	"context"
	"crypto/rand"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/securecookie"
	"github.com/trevorgrabham/cardcounting/cardcounting/lib"
)

var s *securecookie.SecureCookie
var nextUserID *atomic.Int64

func init() {
	var hash, block = make([]byte, 64), make([]byte, 32)
	if _, err := rand.Read(hash); err != nil { panic(err) }
	if _, err := rand.Read(block); err != nil { panic(err) }
	s = securecookie.New(hash, block)
}

func SetCookieContext(next http.Handler) http.Handler {
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