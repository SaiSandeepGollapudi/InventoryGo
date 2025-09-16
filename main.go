package main

import (
	"INVENTORY/api"
	"INVENTORY/dataservice"
	"database/sql"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/inventory?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := dataservice.CreateTableIfNotExists(db); err != nil {
		log.Fatal("Table creation failed:", err)
	}

	producer, err := initKafkaProducer()
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
	defer producer.Close()

	api.RegisterRoutes(db, producer)

	//Start the HTTP server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	//Start the HTTP server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initKafkaProducer() (sarama.SyncProducer, error) {
	brokerList := []string{"localhost:9092"}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}
