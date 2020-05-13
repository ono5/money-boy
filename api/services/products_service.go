// products_service.go

package services

import (
	"github.com/ono5/money-boy/api/domain/products"
	"github.com/ono5/money-boy/api/utils/errors"
)

func CreateProduct(product products.Product) (*products.Product, *errors.ApiErr) {
	return &product, nil
}
