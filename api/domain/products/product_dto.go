// product_dto.go

package products

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/ono5/money-boy/api/utils/errors"
)

// Product - defines product info uploaded by user
type Product struct {
	gorm.Model
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Price  uint64 `json:"price"`
	Img    []byte `json:"img"`
}

// Validate - check parameters user inputs
func (p *Product) Validate() *errors.ApiErr {
	p.Name = strings.TrimSpace(strings.ToLower(p.Name))
	if p.Name == "" {
		return errors.NewBadRequestError("invalid product name")
	}
	return nil
}
