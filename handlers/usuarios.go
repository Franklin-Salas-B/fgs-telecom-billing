package handlers

import (
    "database/sql"
    "fmt"
    "net/http"
    "fgs_telecom_billing/models"
)

type UsuarioForm struct {
    Nombre string
    Rol    string
    Estado string
}

// Handler para mostrar la lista de usuarios
func UsuariosHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        usuarios, err := models.ListarUsuarios(db)
        if err != nil {
            http.Error(w, "Error al obtener usuarios", http.StatusInternalServerError)
            return
        }
        RenderTemplate(w, "usuarios.html", usuarios)
    }
}

// Handler para crear nuevo usuario
func NuevoUsuarioHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
            return
        }

        nombre := r.FormValue("nombre")
        rol := r.FormValue("rol")
        estado := r.FormValue("estado")

        if nombre == "" || rol == "" || estado == "" {
            http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
            return
        }

        err := models.CrearUsuario(db, nombre, rol, estado)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error al crear usuario: %v", err), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
    }
}
