package gateway

import (
	"github.com/C-Agudo/ecomerce-ApiRest/internal/database"
	"github.com/C-Agudo/ecomerce-ApiRest/internal/logs"
	"github.com/C-Agudo/ecomerce-ApiRest/products/domain"
)

type ProductStorageGateway interface {
	create(cmd *domain.CreateProductCMD) (*domain.Product, error)
}

type ProductStorage struct {
	*database.MySqlClient
}

func (s *ProductStorage) create(cmd *domain.CreateProductCMD) (*domain.Product, error) {
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into product (name, price) 
	values (?, ?, ?, ?)`, cmd.Name, cmd.Price)

	if err != nil {
		logs.Log().Error("cannot execute statement")
		_ = tx.Rollback()
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		logs.Log().Error("cannot fetch last id")
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return &domain.Product{
		Id:    id,
		Name:  cmd.Name,
		Price: cmd.Price,
	}, nil
}
