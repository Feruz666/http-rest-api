package store

import "github.com/Feruz666/http-rest-api/internal/app/model"

// UserRepository
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}