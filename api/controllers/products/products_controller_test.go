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
	"github.com/ono5/money-boy/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

func requestHandler(p interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	byteProduct, _ := json.Marshal(p)
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/products",
		bytes.NewBuffer(byteProduct),
	)
	return c, response
}

func TestCreateProductNoError(t *testing.T) {
	// Arrange ---
	p := products.Product{ID: 123, Name: "coca cola"}
	c, response := requestHandler(p)
	// Act ---
	CreateProduct(c)

	// Assert ---
	var product products.Product
	err := json.Unmarshal(response.Body.Bytes(), &product)
	assert.EqualValues(t, http.StatusCreated, response.Code)
	assert.Nil(t, err)
	fmt.Println(product)
	assert.EqualValues(t, uint64(123), product.ID)
}

func TestCreateProductWith404Error(t *testing.T) {
	// Arrange ---
	type demiProduct struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	p := demiProduct{ID: "123", Name: "coca cola"}
	c, response := requestHandler(p)

	// Act ---
	CreateProduct(c)

	// Assert ---
	var apiErr errors.ApiErr
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.Nil(t, err)
	assert.EqualValues(t, "invalid json body", apiErr.Message)
	assert.EqualValues(t, 400, apiErr.Status)
	assert.EqualValues(t, "bad_request", apiErr.Error)
}
