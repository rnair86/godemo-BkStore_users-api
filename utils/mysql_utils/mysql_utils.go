package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/rnair86/godemo-BkStore_users-api/logger"
	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	//fmt.Println(err)
	logger.Error("Error while performing Db operation", err)

	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundRequestError("no records matching given query")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error parsing database response %s", err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid Data")
	}

	return errors.NewInternalServerError("error processing request")

}
