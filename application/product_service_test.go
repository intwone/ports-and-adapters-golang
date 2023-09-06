package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/intwone/ports-and-adapters-golang/application"
	mock_application "github.com/intwone/ports-and-adapters-golang/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should get a product", func(t *testing.T) {
		product := mock_application.NewMockProductInterface(ctrl)
		persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
		persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistence,
		}

		result, err := service.Get("any_id")

		require.Nil(t, err)
		require.Equal(t, product, result)
	})
}

func TestProductService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should save a product", func(t *testing.T) {
		product := mock_application.NewMockProductInterface(ctrl)
		persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistence,
		}

		result, err := service.Create("any name", 100)

		require.Nil(t, err)
		require.Equal(t, product, result)
	})
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should enable a product", func(t *testing.T) {
		product := mock_application.NewMockProductInterface(ctrl)
		product.EXPECT().Enable().Return(nil)

		persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistence,
		}

		result, err := service.Enable(product)

		require.Nil(t, err)
		require.Equal(t, product, result)
	})
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should disable a product", func(t *testing.T) {
		product := mock_application.NewMockProductInterface(ctrl)
		product.EXPECT().Disable().Return(nil)

		persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

		service := application.ProductService{
			Persistence: persistence,
		}

		result, err := service.Disable(product)

		require.Nil(t, err)
		require.Equal(t, product, result)
	})
}
