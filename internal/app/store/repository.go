package store

import "github.com/Constantine27K/rest-api-go.git/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
