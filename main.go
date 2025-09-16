package main

import (
	
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:mysql@2979@tcp(127.0.0.1:3306)/inventory?parseTime=true"
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

	api.RegisterRoutes(db)

	//Start the HTTP server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
