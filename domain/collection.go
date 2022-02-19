package domain

import "example.com/m/domain/lock"

type Collection struct {
	id       string
	memories []Memory
	lock     lock.Lock
}
