package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func UpdateKaryawan(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()

			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			addres := r.Form["addres"][0]
			_, err := db.Exec("UPDATE karyawan SET name=?, npwp=?, addres=? WHERE id=?", name, npwp, addres, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/karyawan", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			row := db.QueryRow("SELECT name, npwp, addres FROM karyawan WHERE id =?", id)
			if row.Err() != nil {
				http.Error(w, row.Err().Error(), http.StatusInternalServerError)
				return
			}

			var karyawan Karyawan
			err := row.Scan(
				&karyawan.Name,
				&karyawan.NPWP,
				&karyawan.Addres,
			)
			karyawan.Id = id
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "update.html")
			tpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["karyawan"] = karyawan

			err = tpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}
}
