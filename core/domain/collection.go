package domain

import "example.com/m/core/domain/lock"

type Collection struct {
	id       string
	memories []Memory
	lock     lock.Lock
}
