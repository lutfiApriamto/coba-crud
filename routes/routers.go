package routes

import (
	"database/sql"
	"net/http"

	"github.com/lutfiApriamto/CRUD-karyawan/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/karyawan", controller.NewIndexKaryawan(db))
	server.HandleFunc("/karyawan/create", controller.NewCreateKaryawanController(db))
	server.HandleFunc("/karyawan/update", controller.UpdateKaryawan(db))
	server.HandleFunc("/karyawan/delete", controller.DeleteEmployeeController(db))

}
