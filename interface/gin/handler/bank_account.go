package handler

import (
	"github.com/amaya382/go-api-server-template/usecase"
	"github.com/gin-gonic/gin"
)

type BankAccountHdlr interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
}

type bankAccountHdlr struct {
	banckAccountUC usecase.BankAccountUC
}

func NewBankAccountHdlr(
	bankAccountUC usecase.BankAccountUC,
) BankAccountHdlr {
	return &bankAccountHdlr{bankAccountUC}
}

type postBankAccount struct {
	AccountHolder string `json:"account_holder" binding:"required,min=3,max=20"`
}

func (self *bankAccountHdlr) Post(c *gin.Context) {
	ctx := GinCtxToCtx(c)

	body := &postBankAccount{}
	err := c.ShouldBindJSON(body)
	if err != nil {
		c.JSON(400, err)
		return
	}

	bankAccount, err := self.banckAccountUC.Create(ctx, body.AccountHolder)
	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(201, bankAccount)
}

type getBankAccount struct {
	AccountNumber string `uri:"account_number" binding:"required,len=7"`
}

func (self *bankAccountHdlr) Get(c *gin.Context) {
	// TODO: Auth middleware

	ctx := GinCtxToCtx(c)

	pathParam := &postBankAccount{}
	err := c.ShouldBindJSON(pathParam)
	if err != nil {
		// TODO: Error handler wrapper
		c.JSON(400, err)
		return
	}

	bankAccount, err := self.banckAccountUC.Create(
		ctx, pathParam.AccountHolder)
	if err != nil {
		c.JSON(400, err)
		return
	}

	// TODO: sheriff wrapper
	c.JSON(200, bankAccount)
}
