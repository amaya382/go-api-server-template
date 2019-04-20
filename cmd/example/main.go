package main

import (
	"github.com/amaya382/go-api-server-template/config"
	"github.com/amaya382/go-api-server-template/domain/service"
	"github.com/amaya382/go-api-server-template/infrastructure/gorm"
	"github.com/amaya382/go-api-server-template/interface/gin/handler"
	v1 "github.com/amaya382/go-api-server-template/interface/gin/router/v1"
	"github.com/amaya382/go-api-server-template/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	rV1 := r.Group("/v1")

	bankAccountRepo := gorm.NewBankAccountRepo()
	transactionRepo := gorm.NewTransactionRepo()
	bankAccountSvc := service.NewBankAccountSvc(bankAccountRepo)
	transferSvc := service.NewTransferSvc(bankAccountRepo, transactionRepo)
	bankAccountUC := usecase.NewBankAccountUC(
		bankAccountRepo, bankAccountSvc)
	transferUC := usecase.NewTransferUC(transferSvc)
	bankAccountHdlr := handler.NewBankAccountHdlr(bankAccountUC)
	transferHdlr := handler.NewTransferHdlr(transferUC)
	rt := v1.NewExampleRT(bankAccountHdlr, transferHdlr)
	rt.RouteToTransfer(rV1)
	rt.RouteToBankAccount(rV1)
	r.Run(":" + config.Config.General.Port)
}
