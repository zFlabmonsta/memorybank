package transaction

import (
	"example.com/m/core/domain/collection"
	"example.com/m/core/domain/memory"
	"example.com/m/core/domain/user"
)

type transaction struct {
	collection collection.Collection
	memory     memory.Memory
	user       user.User
}
