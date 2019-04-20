package gorm

import (
	"context"

	"github.com/amaya382/go-api-server-template/domain/model"
	"github.com/amaya382/go-api-server-template/domain/repository"
	"github.com/amaya382/go-api-server-template/infrastructure/gorm/postgres"
	"github.com/amaya382/go-api-server-template/util"
	"github.com/jinzhu/gorm"
)

type BankAccountRepo struct{}

func NewBankAccountRepo() repository.BankAccountRepo {
	return &BankAccountRepo{}
}

func (self *BankAccountRepo) Create(
	ctx context.Context, bankAccount *model.BankAccount) error {
	db, err := postgres.GetDB(ctx)
	if err != nil {
		return err
	}

	if err := db.Create(bankAccount).Error; err != nil {
		return err
	}
	return nil
}

func (self *BankAccountRepo) Get(
	ctx context.Context, accountNumber string) (*model.BankAccount, error) {
	db, err := postgres.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	bankAccount := &model.BankAccount{AccountNumber: accountNumber}
	if err := db.First(bankAccount, bankAccount).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			tableName := GetTableName(db, bankAccount)
			return nil, util.NewRecordNotFoundErr(tableName)
		}
		return nil, err
	}
	return bankAccount, nil
}

func (self *BankAccountRepo) List(
	ctx context.Context) ([]*model.BankAccount, error) {
	db, err := postgres.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	bankAccounts := []*model.BankAccount{}
	if err := db.Find(&bankAccounts, &bankAccounts).Error; err != nil {
		return nil, err
	}
	return bankAccounts, nil
}

func (self *BankAccountRepo) Replace(
	ctx context.Context, account *model.BankAccount) (
	*model.BankAccount, error) {
	db, err := postgres.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	if err := db.Save(account).Error; err != nil {
		return nil, err
	}

	return account, nil
}

func (self *BankAccountRepo) Delete(
	ctx context.Context, accountNumber string) error {
	db, err := postgres.GetDB(ctx)
	if err != nil {
		return err
	}

	bankAccount := &model.BankAccount{AccountNumber: accountNumber}
	return db.Unscoped().Delete(bankAccount, bankAccount).Error
}
