package domain

type Account struct {
	id          string
	user        User
	collections []Collection
	password    string
}
