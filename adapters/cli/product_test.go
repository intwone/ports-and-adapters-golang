package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/intwone/ports-and-adapters-golang/adapters/cli"
	mock_application "github.com/intwone/ports-and-adapters-golang/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "product1"
	productPrice := 1000
	productStatus := "enabled"
	productId := "1"

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProdutServiceInterface(ctrl)

	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	t.Run("should return correct message when create", func(t *testing.T) {
		resultExpected := fmt.Sprintf("product ID %s with the name %s has been created with price %d and status %s", productId, productName, productPrice, productStatus)
		result, err := cli.Run(service, "create", "", productName, int64(productPrice))

		require.Nil(t, err)
		require.Equal(t, resultExpected, result)
	})

	t.Run("should return correct message when enable", func(t *testing.T) {
		resultExpected := fmt.Sprintf("product %s has been enable", productName)
		result, err := cli.Run(service, "enable", "", productName, int64(productPrice))

		require.Nil(t, err)
		require.Equal(t, resultExpected, result)
	})
}
