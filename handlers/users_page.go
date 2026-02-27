package handlers

import (
	"net/http"
	"text/template"

	"github.com/Franklin-Salas-B/fgs-telecom-billing/db"
	"github.com/Franklin-Salas-B/fgs-telecom-billing/models"
)

func UsersPage(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
		users = append(users, u)
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/users.html",
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.ExecuteTemplate(w, "base", users)
}
