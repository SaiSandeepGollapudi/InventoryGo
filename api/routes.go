package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	h := NewHandler(db)
	http.HandleFunc("/create", h.CreateHandler())
}
