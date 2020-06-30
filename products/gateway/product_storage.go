package gateway

import (
	"log"

	"github.com/C-Agudo/ecomerce-ApiRest/internal/database"
	"github.com/C-Agudo/ecomerce-ApiRest/internal/logs"
	"github.com/C-Agudo/ecomerce-ApiRest/products/domain"
)

type ProductStorageGateway interface {
	Create(cmd *domain.CreateProductCMD) (*domain.Product, error)
	delete(productID int64) *domain.Product
	update(cmd *domain.UpdateProductCMD) *domain.Product
	list() []*domain.Product
}

type ProductStorage struct {
	*database.MySqlClient
}

func (s *ProductStorage) Create(cmd *domain.CreateProductCMD) (*domain.Product, error) {
	tx, err := s.MySqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot create transaction")
		return nil, err
	}

	res, err := tx.Exec(`insert into product (name, price) 
	values (?, ?)`, cmd.Name, cmd.Price)

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

func (s *ProductStorage) delete(productID int64) *domain.Product {
	tx, err := s.Begin()

	if err != nil {
		return nil
	}

	var product domain.Product

	err = tx.QueryRow(`select name, price from product where id = ?`, productID).
		Scan(&product.Name, &product.Price)

	if err != nil {
		_ = tx.Rollback()
		return nil
	}

	_, err = tx.Exec(`delete from product where id = ?`, productID)

	if err != nil {
		logs.Log().Error(err.Error())
		_ = tx.Rollback()
		return nil

	}

	_ = tx.Commit()
	return &domain.Product{
		Id:    productID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func (s *ProductStorage) update(cmd *domain.UpdateProductCMD) *domain.Product {
	tx, err := s.Begin()

	if err != nil {
		return nil
	}

	_, err = tx.Exec(`update product set name = ?, price = ? where id = ?`,
		cmd.Name, cmd.Price, cmd.Id)

	if err != nil {
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()

	return &domain.Product{
		Id:    cmd.Id,
		Name:  cmd.Name,
		Price: cmd.Price,
	}
}

func (s *ProductStorage) list() []*domain.Product {
	tx, err := s.Begin()

	if err != nil {
		return nil
	}

	rows, err := tx.Query(`select * from product`)

	var p []*domain.Product
	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			log.Println("cannot read current row")
			return nil
		}
		p = append(p, &product)
	}
	return p
}
