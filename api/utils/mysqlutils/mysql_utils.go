// mysql_utils.go

package mysqlutils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/ono5/money-boy/api/utils/errors"
)

const (
	noSearchResult = "record not found"
)

// ParseError - parse mysql error to api error
func ParseError(err error) *errors.ApiErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noSearchResult) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062: // duplicate key
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
