package account

import (
	"example.com/m/core/domain/collection"
	"example.com/m/core/domain/user"
)

type Account struct {
	id          string
	user        user.User
	collections []collection.Collection
	password    string
}
