package models

import "database/sql"

// Estructura del usuario
type Usuario struct {
    ID     int
    Nombre string
    Rol    string
    Estado string
}

// CrearUsuario inserta un nuevo usuario en la base de datos
func CrearUsuario(db *sql.DB, nombre, rol, estado string) error {
    query := `INSERT INTO usuarios (nombre, rol, estado) VALUES (?, ?, ?)`
    _, err := db.Exec(query, nombre, rol, estado)
    return err
}

// ListarUsuarios obtiene todos los usuarios de la base de datos
func ListarUsuarios(db *sql.DB) ([]Usuario, error) {
    rows, err := db.Query(`SELECT id, nombre, rol, estado FROM usuarios ORDER BY id DESC`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        if err := rows.Scan(&u.ID, &u.Nombre, &u.Rol, &u.Estado); err != nil {
            return nil, err
        }
        usuarios = append(usuarios, u)
    }
    return usuarios, nil
}
