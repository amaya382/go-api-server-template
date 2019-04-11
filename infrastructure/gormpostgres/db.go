package gormpostgres

import (
	"fmt"

	"github.com/amaya382/go-api-server-template/config"
	"github.com/amaya382/go-api-server-template/domain/model"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
	dbConf := config.Config.DB
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.Name, dbConf.User, dbConf.Password, dbConf.SSL)
	db, err := gorm.Open("postgres", connStr)
	// db, err := gorm.Open("sqlite3", "example.db")
	if err != nil {
		panic(err.Error())
	}

	DB = db
}

func Migrate() error {
	if err := DB.AutoMigrate(&model.User{}).Error; err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.Task{}).Error; err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.TaskList{}).Error; err != nil {
		return err
	}

	return nil
}
