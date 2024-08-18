package main

import (
	"net/http"

	"github.com/lutfiApriamto/CRUD-karyawan/database"
	"github.com/lutfiApriamto/CRUD-karyawan/routes"
)

func main() {
	db := database.InitDatabase()
	defer db.Close() // Pastikan koneksi hanya ditutup ketika aplikasi berhenti

	server := http.NewServeMux()
	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
