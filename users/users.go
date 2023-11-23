package users

import (
	"slices"
)

type User struct {
	Name      string
	interests []string
}

func NewUser(name string, interests ...string) *User {
	return &User{
		Name: name, interests: interests,
	}
}

func (u *User) IsInterested(subject string) bool {
	return slices.Contains(u.interests, subject)
}
