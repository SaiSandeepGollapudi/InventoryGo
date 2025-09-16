package api

import (
	"INVENTORY/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	biz IBizLogic
}

func NewHandler(db *sql.DB) Handler {
	return Handler{biz: NewBizLogic(db)}
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


