package api

import (
	"INVENTORY/dataservice"
	"INVENTORY/model"
	"database/sql"
	"fmt"
)

type IBizLogic interface {
	CreateProductLogic(Product model.Product) error
		UpdateProductLogic(Product model.Product) error

}

type BizLogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{DB: db}
}

func (bl *BizLogic) CreateProductLogic(Product model.Product) error {

	if Product.Name == "" {
		return fmt.Errorf("name should be present")
	}
	if err := dataservice.CreateProduct(bl.DB, Product); err != nil {
		return err
	}

	return nil
}

func (bl *BizLogic) UpdateProductLogic(Product model.Product) error {
	if err := dataservice.UpdateProduct(bl.DB, Product); err != nil {
		return err
	}
	return nil
}
