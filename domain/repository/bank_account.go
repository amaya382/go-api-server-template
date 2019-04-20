package repository

import (
	"context"

	"github.com/amaya382/go-api-server-template/domain/model"
)

type BankAccountRepo interface {
	Create(
		ctx context.Context, bankAccount *model.BankAccount) error
	Get(ctx context.Context, accountNumber string) (*model.BankAccount, error)
	List(ctx context.Context) ([]*model.BankAccount, error)
	Replace(
		ctx context.Context, account *model.BankAccount) (
		*model.BankAccount, error)
	Delete(ctx context.Context, accountNumber string) error
}
