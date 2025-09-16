package api

import (
	"INVENTORY/model"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/IBM/sarama"
)

type Handler struct {
	biz IBizLogic
}

func NewHandler(db *sql.DB, producer sarama.SyncProducer) Handler {
	return Handler{biz: NewBizLogic(db, producer)}
}

func (h Handler) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var Product model.Product
		if err := json.NewDecoder(r.Body).Decode(&Product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.biz.CreateProductLogic(Product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (h Handler) UpdateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut && r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var Product model.Product
		if err := json.NewDecoder(r.Body).Decode(&Product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.biz.UpdateProductLogic(Product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
