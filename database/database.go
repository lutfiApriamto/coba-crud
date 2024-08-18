package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/karyawan")
	if err != nil {
		log.Fatal(err)
	}

	// Coba tes koneksi untuk memastikan database terbuka
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
