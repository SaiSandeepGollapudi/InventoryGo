package api

import (
	"INVENTORY/dataservice"
	"INVENTORY/model"
	"INVENTORY/queue"
	"database/sql"
	"fmt"

	"github.com/IBM/sarama"
)

type IBizLogic interface {
	CreateProductLogic(Product model.Product) error
	UpdateProductLogic(Product model.Product) error
}

type BizLogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

func NewBizLogic(db *sql.DB, producer sarama.SyncProducer) *BizLogic {
	return &BizLogic{DB: db, Producer: producer}
}

func (bl *BizLogic) CreateProductLogic(Product model.Product) error {

	if Product.Name == "" {
		return fmt.Errorf("name should be present")
	}
	if err := dataservice.CreateProduct(bl.DB, Product); err != nil {
		return err
	}

	// produce the message to kafka
	message := fmt.Sprintf("Name: %s has Quantity : %d", Product.Name, Product.Quantity)
	err := queue.ProduceKafkaMessage("product_created_topic", message, bl.Producer)
	if err != nil {
		return fmt.Errorf("failed to produce kafka message: %v", err)
	}

	return nil
}

func (bl *BizLogic) UpdateProductLogic(Product model.Product) error {
	if err := dataservice.UpdateProduct(bl.DB, Product); err != nil {
		return err
	}
	return nil
}
