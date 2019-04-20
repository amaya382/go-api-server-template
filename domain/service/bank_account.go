package service

import (
	"context"

	"github.com/amaya382/go-api-server-template/domain/model"
	"github.com/amaya382/go-api-server-template/domain/repository"
	"github.com/amaya382/go-api-server-template/util"
)

type BankAccountSvc interface {
	Create(
		ctx context.Context, accountHolder string) (
		*model.BankAccount, error)
}

type bankAccountSvc struct {
	bankAccountRepo repository.BankAccountRepo
}

func NewBankAccountSvc(
	bankAccountRepo repository.BankAccountRepo,
) BankAccountSvc {
	return &bankAccountSvc{bankAccountRepo}
}

func (self *bankAccountSvc) Create(
	ctx context.Context, accountHolder string) (
	*model.BankAccount, error) {
	accountNumber := util.RandString(7)
	bankAccount := &model.BankAccount{
		AccountHolder: accountHolder,
		AccountNumber: accountNumber}
	err := self.bankAccountRepo.Create(ctx, bankAccount)
	if err != nil {
		return nil, err
	}

	return bankAccount, nil
}
