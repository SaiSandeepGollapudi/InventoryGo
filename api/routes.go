package api

import (
	"database/sql"
	"net/http"

	"github.com/IBM/sarama"
)

func RegisterRoutes(db *sql.DB, producer sarama.SyncProducer) {
	h := NewHandler(db, producer)
	http.HandleFunc("/create", h.CreateHandler())
	http.HandleFunc("/update", h.UpdateHandler())

}
