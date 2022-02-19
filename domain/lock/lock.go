package lock

type Lock interface {
	Set(string)
	Check() string
	Unlock() bool
}
