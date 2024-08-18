package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewCreateKaryawanController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			addres := r.Form["addres"][0]
			_, err := db.Exec("INSERT INTO karyawan (name, npwp, addres) VALUES (?, ?, ?)", name, npwp, addres)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/karyawan", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create.html")
			tpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}
}
