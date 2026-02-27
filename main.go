package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"

    _ "github.com/go-sql-driver/mysql"
)

type Producto struct {
    ID          int     `json:"id"`
    Nombre      string  `json:"nombre"`
    Descripcion string  `json:"descripcion"`
    Precio      float64 `json:"precio"`
    CreadoEn    string  `json:"creado_en"`
}

var db *sql.DB

func main() {
    var err error
    dsn := "ecom_user:P1ch1nch4@tcp(127.0.0.1:3306)/ecommerce_db"
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Conectado a MySQL correctamente")

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS productos (
            id INT AUTO_INCREMENT PRIMARY KEY,
            nombre VARCHAR(255) NOT NULL,
            descripcion TEXT,
            precio DECIMAL(10,2) NOT NULL,
            creado_en VARCHAR(50) NOT NULL
        );
    `)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/productos", productosHandler)
    http.HandleFunc("/productos/", productoHandler)

    fmt.Println("Servidor corriendo en :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func productosHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        listarProductos(w)
    case "POST":
        crearProducto(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func productoHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/productos/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    switch r.Method {
    case "PUT":
        actualizarProducto(w, r, id)
    case "DELETE":
        eliminarProducto(w, id)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func listarProductos(w http.ResponseWriter) {
    rows, err := db.Query("SELECT id, nombre, descripcion, precio, creado_en FROM productos")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    productos := []Producto{}
    for rows.Next() {
        var p Producto
        if err := rows.Scan(&p.ID, &p.Nombre, &p.Descripcion, &p.Precio, &p.CreadoEn); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        productos = append(productos, p)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(productos)
}

func crearProducto(w http.ResponseWriter, r *http.Request) {
    var p Producto
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    res, err := db.Exec(
        "INSERT INTO productos(nombre, descripcion, precio, creado_en) VALUES (?, ?, ?, ?)",
        p.Nombre, p.Descripcion, p.Precio, p.CreadoEn,
    )
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    id, _ := res.LastInsertId()
    p.ID = int(id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
}

func actualizarProducto(w http.ResponseWriter, r *http.Request, id int) {
    var p Producto
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    _, err := db.Exec(
        "UPDATE productos SET nombre=?, descripcion=?, precio=?, creado_en=? WHERE id=?",
        p.Nombre, p.Descripcion, p.Precio, p.CreadoEn, id,
    )
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    p.ID = id
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(p)
}

func eliminarProducto(w http.ResponseWriter, id int) {
    _, err := db.Exec("DELETE FROM productos WHERE id=?", id)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}
