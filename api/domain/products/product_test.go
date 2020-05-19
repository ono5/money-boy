// product_test.go
package products

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestProductValiateNoError(t *testing.T) {
	// Arrange ---
	p := Product{Model: gorm.Model{ID: 123}, Name: "coca cola"}

	// Act ---
	err := p.Validate()

	// Assert ---
	assert.Nil(t, err)
}

func TestProductValiateBadRequestError(t *testing.T) {
	// Arrange ---
	p := Product{Model: gorm.Model{ID: 123}}

	// Act ---
	err := p.Validate()

	// Assert ---
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid product name", err.Message)
	assert.EqualValues(t, 400, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)

}
