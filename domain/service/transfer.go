package service

import (
	"context"

	"github.com/amaya382/go-api-server-template/domain/repository"
)

type TransferSvc interface {
	Transfer(ctx context.Context,
		amount uint, senderNumber string, receiverNumber string) error
}

type transferSvc struct {
	bankAccountRepo repository.BankAccountRepo
	transactionRepo repository.TransactionRepo
}

func NewTransferSvc(
	bankAccountRepo repository.BankAccountRepo,
	transactionRepo repository.TransactionRepo) TransferSvc {
	return &transferSvc{bankAccountRepo, transactionRepo}
}

func (self *transferSvc) Transfer(ctx context.Context,
	amount uint, senderNumber string, receiverNumber string) error {
	ctx, err := self.transactionRepo.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			self.transactionRepo.Rollback(ctx)
		}
	}()

	sender, err := self.bankAccountRepo.Get(ctx, senderNumber)
	if err != nil {
		return err
	}
	receiver, err := self.bankAccountRepo.Get(ctx, receiverNumber)
	if err != nil {
		return err
	}

	if err := sender.Withdraw(amount); err != nil {
		return err
	}
	if err := receiver.Deposit(amount); err != nil {
		return err
	}

	if _, err := self.bankAccountRepo.Replace(ctx, sender); err != nil {
		return err
	}
	if _, err := self.bankAccountRepo.Replace(ctx, receiver); err != nil {
		return err
	}

	return self.transactionRepo.Commit(ctx)
}
