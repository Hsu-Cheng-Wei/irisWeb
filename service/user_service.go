package service

import "irisWeb/domain"

type UserService interface {
	GetAll() []domain.User
	Insert(u domain.User) ([]byte, error)
	Delete(id string) error
}
