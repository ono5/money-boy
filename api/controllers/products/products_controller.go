// controller/products/products_controller.go

package products

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ono5/money-boy/api/domain/products"
	"github.com/ono5/money-boy/api/services"
	"github.com/ono5/money-boy/api/utils/errors"
)

// UpdateProduct - Update product
func UpdateProduct(c *gin.Context) {
	productID, productErr := strconv.ParseUint(c.Param("product_id"), 10, 64)
	if productErr != nil {
		err := errors.NewBadRequestError("product id should be a number")
		c.JSON(err.Status, err)
		return
	}

	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	product.ID = uint(productID)

	result, err := services.UpdateProduct(product)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

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

	product, getErr := services.GetProduct(uint(productID))
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, product)
}
