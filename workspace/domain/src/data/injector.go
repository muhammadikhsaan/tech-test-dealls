package data

import (
	"dealls.test/domain/src/data/interactions"
	"dealls.test/domain/src/data/privilages"
	"dealls.test/domain/src/data/users"
)

type Repository struct {
	User        users.Repository
	Interaction interactions.Repository
	Privilages  privilages.Repository
}

func NewRepository() *Repository {
	return &Repository{
		User:        users.NewRepository(),
		Interaction: interactions.NewRepository(),
		Privilages:  privilages.NewRepository(),
	}
}
