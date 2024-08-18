package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Karyawan struct {
	Id     string
	Name   string
	NPWP   string
	Addres string
}

func NewIndexKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, npwp, addres FROM karyawan ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var karyawans []Karyawan
		for rows.Next() {

			var karyawan Karyawan

			err = rows.Scan(
				&karyawan.Id,
				&karyawan.Name,
				&karyawan.NPWP,
				&karyawan.Addres,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			karyawans = append(karyawans, karyawan)

		}

		fp := filepath.Join("views", "index.html")
		tpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["karyawans"] = karyawans

		err = tpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
