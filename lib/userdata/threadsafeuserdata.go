package userdata

import (
	"fmt"
	"sync"

	"github.com/trevorgrabham/cardcounting/cardcounting/lib/cards"
)

type UserData struct {
	deck cards.Deck 
	currentCount int16
	numDecks uint8
	countErrors uint16
	strategyErrors uint16
}

type ThreadSafeUserData struct {
	data map[int64]UserData
	lock sync.Mutex
}

func NewThreadSafeUserData() *ThreadSafeUserData {
	return &ThreadSafeUserData{ data: make(map[int64]UserData) }
}

func (d *ThreadSafeUserData) AddDeck(userID int64, deck cards.Deck) {
	d.lock.Lock()
	user, ok := d.data[userID]
	if !ok {
		user = UserData{ deck: deck, numDecks: uint8(len(deck)/52) }
	} else { user.deck = deck }
	d.data[userID] = user
	d.lock.Unlock()
}

func (d *ThreadSafeUserData) Draw(userID int64) (card cards.Card, finished bool) {
	d.lock.Lock()
	defer d.lock.Unlock()
	user, ok := d.data[userID]

	// not sure how they got this far without having any user data
	if !ok { return cards.Card{}, true }
	// We will never have to worry about the case where ok = true, but len(deck) = 0, since we check before returning the last card and would have deleted data[userID] before returning the last card

	deck := user.deck
	// clearly they forgot to check value of 'finished'
	if deck == nil { return cards.Card{}, true }

	c := deck[0]
	switch {
	case c <= cards.Unknown:
		panic(fmt.Errorf("drew an 'Uknowkn' or 'Nil' card"))
	case c <= cards.SixOfHearts:
		user.currentCount++
	case c >= cards.TenOfClubs:
		user.currentCount--
	}
	card = cards.NewCard(c)
	if err := card.GetSVG(); err != nil { panic(err) }
	if len(deck) <= 1 {
		user.deck = nil
		d.data[userID] = user
		return card, true
	}
	user.deck = deck[1:]
	d.data[userID] = user
	return card, false
}

func (d *ThreadSafeUserData) CheckCount(userID int64, count int16) (correct bool, correctCount int16) {
	d.lock.Lock()
	defer d.lock.Unlock()
	user, ok := d.data[userID]
	if !ok { panic(fmt.Errorf("called getcount() without any user data")) }
	if count == user.currentCount {
		return true, count
	}
	user.countErrors++
	d.data[userID] = user 
	return false, user.currentCount
}

func (d *ThreadSafeUserData) GetErrors(userID int64) (countingErrors, strategyErrors uint16) {
	d.lock.Lock()
	defer d.lock.Unlock()
	user, ok := d.data[userID]
	if !ok { panic(fmt.Errorf("called geterrors() without any user data")) }
	return user.countErrors, user.strategyErrors
}

func (d *ThreadSafeUserData) IncStrategyErrors(userID int64) {
	d.lock.Lock()
	defer d.lock.Unlock()
	user, ok := d.data[userID]
	if !ok { panic(fmt.Errorf("called incstrategyerrors() without any user data")) }
	user.strategyErrors++ 
	d.data[userID] = user
}