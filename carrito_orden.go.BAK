package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

// --- Estructuras ---
type Carrito struct {
    ID         int    `json:"id"`
    UsuarioID  int    `json:"usuario_id"`
    ProductoID int    `json:"producto_id"`
    Cantidad   int    `json:"cantidad"`
    CreadoEn   string `json:"creado_en"`
}

type Orden struct {
    ID        int     `json:"id"`
    UsuarioID int     `json:"usuario_id"`
    Total     float64 `json:"total"`
    Estado    string  `json:"estado"`
    CreadoEn  string  `json:"creado_en"`
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
    fmt.Println("Conectado a MySQL correctamente (carrito y órdenes)")

    // Endpoints
    http.HandleFunc("/carrito", carritoHandler)
    http.HandleFunc("/ordenes", ordenHandler)
    http.HandleFunc("/checkout", checkoutHandler)

    fmt.Println("Servidor corriendo en :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}

// --- Handlers ---
func carritoHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        listarCarrito(w)
    case "POST":
        agregarCarrito(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func ordenHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        listarOrdenes(w)
    case "POST":
        crearOrden(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

// --- Funciones CRUD Carrito ---
func listarCarrito(w http.ResponseWriter) {
    rows, err := db.Query("SELECT id, usuario_id, producto_id, cantidad, creado_en FROM carrito")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    items := []Carrito{} // slice inicializado correctamente
    for rows.Next() {
        var c Carrito
        if err := rows.Scan(&c.ID, &c.UsuarioID, &c.ProductoID, &c.Cantidad, &c.CreadoEn); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        items = append(items, c)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func agregarCarrito(w http.ResponseWriter, r *http.Request) {
    var c Carrito
    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    res, err := db.Exec(
        "INSERT INTO carrito(usuario_id, producto_id, cantidad, creado_en) VALUES (?, ?, ?, ?)",
        c.UsuarioID, c.ProductoID, c.Cantidad, c.CreadoEn,
    )
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    id, _ := res.LastInsertId()
    c.ID = int(id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(c)
}

// --- Funciones CRUD Orden ---
func listarOrdenes(w http.ResponseWriter) {
    rows, err := db.Query("SELECT id, usuario_id, total, estado, creado_en FROM ordenes")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    ordenes := []Orden{} // slice inicializado correctamente
    for rows.Next() {
        var o Orden
        if err := rows.Scan(&o.ID, &o.UsuarioID, &o.Total, &o.Estado, &o.CreadoEn); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        ordenes = append(ordenes, o)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(ordenes)
}

func crearOrden(w http.ResponseWriter, r *http.Request) {
    var o Orden
    if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    res, err := db.Exec(
        "INSERT INTO ordenes(usuario_id, total, estado, creado_en) VALUES (?, ?, ?, ?)",
        o.UsuarioID, o.Total, o.Estado, o.CreadoEn,
    )
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    id, _ := res.LastInsertId()
    o.ID = int(id)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(o)
}

// --- Función Checkout ---
func checkoutHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    type CheckoutRequest struct {
        UsuarioID int `json:"usuario_id"`
    }

    var req CheckoutRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), 400)
        return
    }

    // Obtener items del carrito del usuario
    rows, err := db.Query("SELECT producto_id, cantidad FROM carrito WHERE usuario_id = ?", req.UsuarioID)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    type Item struct {
        ProductoID int
        Cantidad   int
        Precio     float64
    }

    var items []Item
    var total float64 = 0

    for rows.Next() {
        var it Item
        if err := rows.Scan(&it.ProductoID, &it.Cantidad); err != nil {
            http.Error(w, err.Error(), 500)
            return
        }

        // Obtener precio del producto
        err := db.QueryRow("SELECT precio FROM productos WHERE id = ?", it.ProductoID).Scan(&it.Precio)
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }

        total += it.Precio * float64(it.Cantidad)
        items = append(items, it)
    }

    if len(items) == 0 {
        http.Error(w, "El carrito está vacío", 400)
        return
    }

    // Crear orden
    res, err := db.Exec("INSERT INTO ordenes(usuario_id, total, estado, creado_en) VALUES (?, ?, ?, NOW())",
        req.UsuarioID, total, "pendiente")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    ordenID, _ := res.LastInsertId()

    // Insertar detalle_orden
    for _, it := range items {
        _, err := db.Exec("INSERT INTO detalle_orden(orden_id, producto_id, cantidad, precio) VALUES (?, ?, ?, ?)",
            ordenID, it.ProductoID, it.Cantidad, it.Precio)
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
    }

    // Vaciar carrito
    _, err = db.Exec("DELETE FROM carrito WHERE usuario_id = ?", req.UsuarioID)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "orden_id": ordenID,
        "total":    total,
        "estado":   "pendiente",
    })
}
