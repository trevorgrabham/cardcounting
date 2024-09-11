package lib

import (
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/userdata"
)

type ContextKey string

var UserData = userdata.NewThreadSafeUserData()

var ErrorStrings = map[string]string{
	"training-correct-strategy-format":  `Correct! The proper strategy was to %s!`,
	"training-wrong-strategy-format": `Incorrect! The correct strategy was to %s!`,
	"training-correct-count-format": `Correct! The count is %d!`,
	"training-wronqag-count-format": `Incorrect! The current count is %d!`,
	"training-settings-deck-string": `Please choose a number of decks between 1 and 8 to train with`,
}