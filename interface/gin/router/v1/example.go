package v1

import (
	"github.com/amaya382/go-api-server-template/interface/gin/handler"
	"github.com/gin-gonic/gin"
)

type ExampleRT interface {
	RouteToBankAccount(rg *gin.RouterGroup)
	RouteToTransfer(rg *gin.RouterGroup)
}

type exampleRT struct {
	bankAccountHdlr handler.BankAccountHdlr
	transferHdlr    handler.TransferHdlr
}

func NewExampleRT(
	bankAccountHdlr handler.BankAccountHdlr,
	transferHdlr handler.TransferHdlr,
) ExampleRT {
	return &exampleRT{bankAccountHdlr, transferHdlr}
}

func (self *exampleRT) RouteToTransfer(rg *gin.RouterGroup) {
	rg = rg.Group("/action")
	rg.POST("/transfer", self.transferHdlr.Transfer)
}

func (self *exampleRT) RouteToBankAccount(rg *gin.RouterGroup) {
	// for _, middleware := range middlewares {
	// 	rg.Use(middleware)
	// }

	rg.POST("/account", self.bankAccountHdlr.Post)
	rg.GET("/account/:number", self.bankAccountHdlr.Get)
}
