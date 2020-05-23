// product_dao.go

package products

import (
	"fmt"
	"log"

	"github.com/ono5/money-boy/api/datasources/mysql/products_db"
	"github.com/ono5/money-boy/api/utils/errors"
)

var (
	productsDB = make(map[uint]*Product)
)

// Get - product
func (p *Product) Get() *errors.ApiErr {
	// https://gorm.io/ja_JP/docs/generic_interface.html
	if err := products_db.Client.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	result := productsDB[p.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("product %d not found", p.ID))
	}

	p.ID = result.ID
	p.Name = result.Name
	p.Detail = result.Detail
	p.Price = result.Price
	p.Img = result.Img
	p.CreatedAt = result.CreatedAt
	p.UpdatedAt = result.UpdatedAt
	p.DeletedAt = result.DeletedAt

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
