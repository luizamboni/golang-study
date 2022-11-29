package repo

import "go-study.com/m/core"

type ProductRepo struct {
}

func (r RulesRepo) GetProducts() []core.Product {

	return []core.Product{
		{
			Name: "carro veloz",
			Ean:  "Não tem",
		},
	}
}

// or panic
func NewProductRepo() *ProductRepo {

	return &ProductRepo{}
}
