// product_dao.go

package products

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/ono5/money-boy/api/datasources/mysql/products_db"
	"github.com/ono5/money-boy/api/utils/errors"
)

const (
	noSearchResult = "record not found"
)

// Get - product
func (p *Product) Get() *errors.ApiErr {
	if result := products_db.Client.Where("id = ?", p.Model.ID).Find(&p); result.Error != nil {
		errMsg := fmt.Sprintf("%s", result.Error)
		if strings.Contains(errMsg, noSearchResult) {
			return errors.NewNotFoundError(
				fmt.Sprintf("product %d not found", p.ID),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get product: %s", result.Error),
		)
	}
	return nil
}

// Save - product
func (p *Product) Save() *errors.ApiErr {
	// https://gorm.io/ja_JP/docs/error_handling.html
	if result := products_db.Client.Create(&p); result.Error != nil {
		sqlErr, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(
				fmt.Sprintf("error when trying to save product: %s", result.GetErrors()),
			)
		}

		// Check error number -> 1062
		fmt.Println("Mysql Error Key", sqlErr.Number)
		switch sqlErr.Number {
		case 1062: // duplicate key
			return errors.NewBadRequestError(fmt.Sprintf("name '%s' is already exists", p.Name))
		}

		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save product: %s", result.GetErrors()),
		)
	}
	return nil
}
