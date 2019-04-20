package main

import (
	"github.com/amaya382/go-api-server-template/infrastructure/gorm"
)

func main() {
	com := gorm.NewCommonRepo()
	com.Migrate()
}
