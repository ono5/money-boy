// product_dao.go

package products

import (
	"fmt"

	"github.com/ono5/money-boy/api/datasources/mysql/products_db"
	"github.com/ono5/money-boy/api/utils/errors"
)

// Get - product
func (p *Product) Get() *errors.ApiErr {
	if result := products_db.Client.Where("id = ?", p.Model.ID).Find(&p); result.Error != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get product: %s", result.GetErrors()),
		)
	}
	return nil
}

// Save - product
func (p *Product) Save() *errors.ApiErr {
	// https://gorm.io/ja_JP/docs/error_handling.html
	if result := products_db.Client.Create(&p); result.Error != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save product: %s", result.GetErrors()),
		)
	}
	return nil
}
