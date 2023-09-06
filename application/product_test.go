package application_test

import (
	"testing"

	"github.com/intwone/ports-and-adapters-golang/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	t.Run("should enable product when price is greater than zero", func(t *testing.T) {
		product := application.Product{
			Name:   "any name",
			Status: string(application.Disabled),
			Price:  10,
		}

		err := product.Enable()

		require.Nil(t, err)
	})

	t.Run("should not enable product when price is smaller or equal than zero", func(t *testing.T) {
		product := application.Product{
			Name:   "any name",
			Status: string(application.Disabled),
			Price:  0,
		}

		err := product.Enable()

		require.NotNil(t, err)
		require.Equal(t, "the price must be greather than zero to enable the product", error.Error(err))
	})
}

func TestProduct_Disable(t *testing.T) {
	t.Run("should disable product when price is equal zero", func(t *testing.T) {
		product := application.Product{
			Name:   "any name",
			Status: string(application.Enabled),
			Price:  10,
		}

		product.Price = 0

		err := product.Disable()

		require.Nil(t, err)
	})

	t.Run("should not disable product when price is different than zero", func(t *testing.T) {
		product := application.Product{
			Name:   "any name",
			Status: string(application.Enabled),
			Price:  10,
		}

		err := product.Disable()

		require.NotNil(t, err)
		require.Equal(t, "the price must be zero to disable the product", error.Error(err))
	})
}

func TestProduct_IsValid(t *testing.T) {
	t.Run("should not return an error when product is valid", func(t *testing.T) {
		product := application.Product{
			Id:     uuid.NewV4().String(),
			Name:   "any name",
			Status: string(application.Disabled),
			Price:  10,
		}

		result, err := product.IsValid()

		require.Nil(t, err)
		require.True(t, result)
	})

	t.Run("should returnfalse when product price is smaller then zero", func(t *testing.T) {
		product := application.Product{
			Id:     uuid.NewV4().String(),
			Name:   "any name",
			Status: string(application.Disabled),
			Price:  10,
		}

		product.Price = -1

		result, err := product.IsValid()

		require.False(t, result)
		require.Equal(t, "the price must be greather than zero", error.Error(err))
	})

	t.Run("should return false when product validator return error", func(t *testing.T) {
		product := application.Product{
			Id:     uuid.NewV4().String(),
			Status: string(application.Disabled),
			Price:  10,
		}

		result, _ := product.IsValid()

		require.False(t, result)
	})
}
