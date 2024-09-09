package lib

import (
	"github.com/trevorgrabham/cardcounting/cardcounting/lib/userdata"
)

type ContextKey string

var UserData = userdata.NewThreadSafeUserData()
