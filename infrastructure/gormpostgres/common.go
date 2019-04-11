package gormpostgres

import (
	"regexp"

	"github.com/amaya382/go-api-server-template/util"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

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
