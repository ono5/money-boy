// domain/products/product_dao.go

package products

import (
	"github.com/ono5/money-boy/api/datasources/mysql/products_db"
	"github.com/ono5/money-boy/api/utils/errors"
	"github.com/ono5/money-boy/api/utils/mysqlutils"
)

// PartialUpdate - product
func (p *Product) PartialUpdate() *errors.ApiErr {
	if result := products_db.Client.
		Table("products").
		Where("id IN (?)", p.ID).
		Updates(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

// Update - product
func (p *Product) Update() *errors.ApiErr {
	if result := products_db.Client.Save(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

// Get - product
func (p *Product) Get() *errors.ApiErr {
	if result := products_db.Client.Where("id = ?", p.Model.ID).Find(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}

// Save - product
func (p *Product) Save() *errors.ApiErr {
	// https://gorm.io/ja_JP/docs/error_handling.html
	if result := products_db.Client.Create(&p); result.Error != nil {
		return mysqlutils.ParseError(result.Error)
	}
	return nil
}
