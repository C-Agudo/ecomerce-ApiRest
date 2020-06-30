package gateway

import (
	"github.com/C-Agudo/ecomerce-ApiRest/internal/database"
	"github.com/C-Agudo/ecomerce-ApiRest/products/domain"
)

type ProductGateway interface {
	Create(cmd *domain.CreateProductCMD) (*domain.Product, error)
	Delete(productId int64) *domain.Product
	Update(cmd *domain.UpdateProductCMD) *domain.Product
	List() []*domain.Product
}

type ProductGtw struct {
	ProductStorageGateway
}

func NewProductCreateGateway(db *database.MySqlClient) ProductGateway {
	return &ProductGtw{&ProductStorage{db}}
}

func (g *ProductGtw) Delete(productId int64) *domain.Product {
	return g.delete(productId)
}

func (g *ProductGtw) Update(cmd *domain.UpdateProductCMD) *domain.Product {
	return g.update(cmd)
}

func (g *ProductGtw) List() []*domain.Product {
	return g.list()
}
