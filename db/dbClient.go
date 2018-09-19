package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// DB database
var DB *sql.DB

// InitDbClient init ddatabase
func InitDbClient() {
	dbs, err := sql.Open("mysql", "root:123456@tcp(192.168.1.9:3306)/jun")
	if err != nil {
		log.Println("database connect failed")
		log.Println(err)
	}
	DB = dbs
}
