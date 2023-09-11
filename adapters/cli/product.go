package cli

import (
	"fmt"

	"github.com/intwone/ports-and-adapters-golang/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice int64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("product ID %s with the name %s has been created with price %d and status %s", product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("product %s has been enable", res.GetName())

	case "disable":
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("product %s has been disable", res.GetName())

	default:
		product, err := service.Get(productId)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("product ID %s \nname %s \nprice %f \nstatus %s", product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
