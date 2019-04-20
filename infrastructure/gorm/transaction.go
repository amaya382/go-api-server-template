package gorm

import (
	"context"
	"errors"

	"github.com/amaya382/go-api-server-template/domain/repository"
	"github.com/amaya382/go-api-server-template/infrastructure/gorm/postgres"
	"github.com/jinzhu/gorm"
)

type transactionRepo struct{}

func NewTransactionRepo() repository.TransactionRepo {
	return &transactionRepo{}
}

func (repo *transactionRepo) Begin(ctx context.Context) (context.Context, error) {
	db, err := postgres.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	tx := db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return context.WithValue(ctx, "tx", tx), nil
}

func (repo *transactionRepo) Commit(ctx context.Context) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		return errors.New("hoge")
	}

	return tx.Commit().Error
}

func (repo *transactionRepo) Rollback(ctx context.Context) error {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if !ok {
		return errors.New("hoge")
	}

	return tx.Rollback().Error
}
