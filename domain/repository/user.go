package repository

import "github.com/amaya382/go-api-server-template/domain/model"

type UserRepo interface {
	Create(*model.User) (*model.User, error)
	Get(id uint) (*model.User, error)
	List() ([]*model.User, error)
	Delete(id uint) error
}
