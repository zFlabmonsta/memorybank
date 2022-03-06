package collection

import (
	"example.com/m/core/domain/lock"
	"example.com/m/core/domain/memory"
)

type Collection struct {
	id       string
	memories []memory.Memory
	lock     lock.Lock
}
