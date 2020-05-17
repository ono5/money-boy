// products_controller.go

package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ono5/money-boy/api/domain/products"
	"github.com/ono5/money-boy/api/services"
	"github.com/ono5/money-boy/api/utils/errors"
)

// CreateProduct - Create product
func CreateProduct(c *gin.Context) {
	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}
	newProduct, saveErr := services.CreateProduct(product)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

// GetProduct - Get product by product id
func GetProduct(c *gin.Context) {
	productID, productErr := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if productErr != nil {
		err := errors.NewBadRequestError("product id should be a number")
		c.JSON(err.Status, err)
		return
	}

	product, getErr := services.GetProduct(productID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, product)
}
