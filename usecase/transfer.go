package usecase

import (
	"context"

	"github.com/amaya382/go-api-server-template/domain/service"
)

type TransferUC interface {
	Transfer(ctx context.Context,
		amount uint, senderNumber string, receiverNumber string) error
}

type transferUC struct {
	transferSvc service.TransferSvc
}

func NewTransferUC(
	transferSvc service.TransferSvc,
) TransferUC {
	return &transferUC{transferSvc}
}

func (self *transferUC) Transfer(
	ctx context.Context, amount uint, senderNumber string, receiverNumber string) error {
	return self.transferSvc.Transfer(
		ctx, amount, senderNumber, receiverNumber)
}
