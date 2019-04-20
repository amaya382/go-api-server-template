package gorm

import (
	"github.com/amaya382/go-api-server-template/domain/model"
	"github.com/amaya382/go-api-server-template/util"
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (self *UserRepo) Create(user *model.User) (*model.User, error) {
	if err := self.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (self *UserRepo) Get(id uint) (*model.User, error) {
	user := &model.User{ID: id}
	if err := self.db.First(user, user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			tableName := GetTableName(self.db, user)
			return nil, util.NewRecordNotFoundErr(tableName)
		}
		return nil, err
	}
	return user, nil
}

func (self *UserRepo) List() ([]*model.User, error) {
	users := []*model.User{}
	if err := self.db.Find(&users, &users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (self *UserRepo) Delete(id uint) error {
	user := &model.User{ID: id}
	return self.db.Unscoped().Delete(user, user).Error
}
