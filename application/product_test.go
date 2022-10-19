package application_test

import (
	"github/arlenmendes/hexagonal-arq-studies/application"
	"testing"

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
