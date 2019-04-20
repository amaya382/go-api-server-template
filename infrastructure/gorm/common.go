package gorm

import (
	"context"
	"regexp"

	"github.com/amaya382/go-api-server-template/domain/model"
	"github.com/amaya382/go-api-server-template/infrastructure/gorm/postgres"
	"github.com/amaya382/go-api-server-template/util"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type CommonRepo struct{}

func NewCommonRepo() *CommonRepo {
	return &CommonRepo{}
}

func (self *CommonRepo) Migrate() error {
	db, err := postgres.GetDB(context.Background())
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.AutoMigrate(&model.BankAccount{}).Error; err != nil {
		return err
	}
	if err := tx.AutoMigrate(&model.User{}).Error; err != nil {
		return err
	}
	if err := tx.AutoMigrate(&model.Task{}).Error; err != nil {
		return err
	}
	if err := tx.AutoMigrate(&model.TaskList{}).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func CheckUniqueViolationErr(err error) (bool, *util.UniqueViolationErr) {
	switch typedErr := err.(type) {
	case *pq.Error:
		if typedErr.Code == "23505" {
			reColInConstraint := regexp.MustCompile("^" + typedErr.Table + "_(.+)_key$")
			cands := reColInConstraint.FindStringSubmatch(typedErr.Constraint)
			col := ""
			if len(cands) > 1 {
				col = cands[1]
			}
			return true, util.NewUniqueViolationErr(typedErr.Table, col, "")
		}
		return false, nil
	case *util.UniqueViolationErr:
		return true, typedErr
	default:
		return false, nil
	}
}

func GetTableName(db *gorm.DB, model interface{}) string {
	return db.NewScope(model).GetModelStruct().TableName(db)
}
