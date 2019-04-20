package handler

import (
	"github.com/amaya382/go-api-server-template/usecase"
	"github.com/gin-gonic/gin"
)

type TransferHdlr interface {
	Transfer(c *gin.Context)
}

type transferHdlr struct {
	transferUC usecase.TransferUC
}

func NewTransferHdlr(transferUC usecase.TransferUC) TransferHdlr {
	return &transferHdlr{transferUC}
}

type postTransfer struct {
	SenderNumber   string `json:"sender_number" binding:"required,len=7"`
	ReceiverNumber string `json:"receiver_number" binding:"required,len=7"`
	Amount         uint   `json:"amount" binding:"required"`
}

func (self *transferHdlr) Transfer(c *gin.Context) {
	ctx := GinCtxToCtx(c)

	body := &postTransfer{}
	err := c.ShouldBindJSON(body)
	if err != nil {
		// TODO: Error handler wrapper
		c.JSON(400, err)
		return
	}

	err = self.transferUC.Transfer(
		ctx, body.Amount, body.SenderNumber, body.ReceiverNumber)
	if err != nil {
		// TODO: Error handler wrapper
		c.JSON(400, err)
		return
	}

	c.JSON(200, body)
}
