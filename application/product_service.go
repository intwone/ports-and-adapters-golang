package application

type ProdutServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (ps *ProductService) Get(id string) (ProductInterface, error) {
	result, err := ps.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ps *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()

	if err != nil {
		return nil, err
	}

	result, err := ps.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ps *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()

	if err != nil {
		return nil, err
	}

	result, err := ps.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ps *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()

	if err != nil {
		return nil, err
	}

	result, err := ps.Persistence.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}
