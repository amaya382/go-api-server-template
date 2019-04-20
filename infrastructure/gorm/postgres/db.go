package postgres

import (
	"context"
	"fmt"

	"github.com/amaya382/go-api-server-template/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	dbConf := config.Config.DB
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.Name, dbConf.User, dbConf.Password, dbConf.SSL)
	pg, err := gorm.Open("postgres", connStr)
	// db, err := gorm.Open("sqlite3", "example.db")
	if err != nil {
		panic(err.Error())
	}

	db = pg
	db.LogMode(true) // TMP
}

func GetDB(ctx context.Context) (*gorm.DB, error) {
	tx, ok := ctx.Value("tx").(*gorm.DB)
	if ok {
		return tx, nil
	}

	return db, nil
}
