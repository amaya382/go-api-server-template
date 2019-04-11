package gormpostgres

import (
	"github.com/amaya382/go-api-server-template/domain/model"
	"github.com/amaya382/go-api-server-template/util"
	"github.com/jinzhu/gorm"
)

type UserRepo struct{}

func (self *UserRepo) Create(user *model.User) (*model.User, error) {
	if err := DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (self *UserRepo) Get(id uint) (*model.User, error) {
	user := &model.User{ID: id}
	if err := DB.First(user, user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			tableName := GetTableName(DB, user)
			return nil, util.NewRecordNotFoundErr(tableName)
		}
		return nil, err
	}
	return user, nil
}

func (self *UserRepo) List() ([]*model.User, error) {
	users := []*model.User{}
	if err := DB.Find(&users, &users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (self *UserRepo) Delete(id uint) error {
	user := &model.User{ID: id}
	return DB.Unscoped().Delete(user, user).Error
}
