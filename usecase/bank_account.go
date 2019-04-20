package usecase

import (
	"context"

	"github.com/amaya382/go-api-server-template/domain/model"
	"github.com/amaya382/go-api-server-template/domain/repository"
	"github.com/amaya382/go-api-server-template/domain/service"
)

type BankAccountUC interface {
	Create(
		ctx context.Context, accountHolder string) (
		*model.BankAccount, error)
	Get(
		ctx context.Context, accountNumber string) (
		*model.BankAccount, error)
}

type bankAccountUC struct {
	bankAccountRepo repository.BankAccountRepo
	bankAccountSvc  service.BankAccountSvc
}

func NewBankAccountUC(
	bankAccountRepo repository.BankAccountRepo,
	bankAccountSvc service.BankAccountSvc,
) BankAccountUC {
	return &bankAccountUC{bankAccountRepo, bankAccountSvc}
}

func (self *bankAccountUC) Create(
	ctx context.Context, accountHolder string) (
	*model.BankAccount, error) {
	return self.bankAccountSvc.Create(ctx, accountHolder)
}

func (self *bankAccountUC) Get(
	ctx context.Context, accountNumber string) (*model.BankAccount, error) {
	return self.bankAccountRepo.Get(ctx, accountNumber)
}
