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

func getProductID(productIDParam string) (uint, *errors.ApiErr) {
	productID, productErr := strconv.ParseUint(productIDParam, 10, 64)
	if productErr != nil {
		return 0, errors.NewBadRequestError("product id should be a number")
	}
	return uint(productID), nil
}

// DeleteProduct - Delete product
func DeleteProduct(c *gin.Context) {
	productID, idErr := getProductID(c.Param("product_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteProduct(productID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// UpdateProduct - Update product
func UpdateProduct(c *gin.Context) {
	productID, idErr := getProductID(c.Param("product_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var product products.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status, apiErr)
		return
	}

	product.ID = productID

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateProduct(isPartial, product)
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
	productID, idErr := getProductID(c.Param("product_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	product, getErr := services.GetProduct(productID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, product)
}
