// products_controller_test.go

package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ono5/money-boy/api/domain/products"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductNoError(t *testing.T) {
	// Arrange ---
	p := products.Product{ID: 123, Name: "coca cola"}
	byteProduct, _ := json.Marshal(p)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/products",
		bytes.NewBuffer(byteProduct),
	)

	// Act ---
	CreateProduct(c)

	// Assert ---
	var product products.Product
	err := json.Unmarshal(response.Body.Bytes(), &product)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	fmt.Println(product)
	assert.EqualValues(t, uint64(123), product.ID)
}
