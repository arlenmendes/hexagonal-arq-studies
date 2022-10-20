package application_test

import (
	"github/arlenmendes/hexagonal-arq-studies/application"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.DISABLED

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.ENABLED

	err := product.Disable()
	require.Nil(t, err)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Price = 10
	product.Status = application.ENABLED

	isValid, err := product.IsValid()
	require.Nil(t, err)
	require.True(t, isValid)
}
