package main

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"
    "strconv"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

// -------------------- Conexión DB --------------------
var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("mysql", "fgs:fgs1234@tcp(localhost:3306)/fgs_billing")
    if err != nil {
        log.Fatal("Error al conectar a MySQL:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("Error al hacer ping a DB:", err)
    }

    // -------------------- Rutas --------------------
    // Usuarios
    http.HandleFunc("/usuarios", usuariosHandler)
    http.HandleFunc("/usuarios/agregar", agregarUsuarioHandler)

    // Productos
    http.HandleFunc("/productos", productosHandler)
    http.HandleFunc("/productos/agregar", agregarProductoHandler)

    // Facturas
    http.HandleFunc("/facturas", facturasHandler)
    http.HandleFunc("/facturas/agregar", agregarFacturaHandler)

    // Archivos estáticos
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Dashboard (raíz) al final
    http.HandleFunc("/", dashboardHandler)

    log.Println("Servidor iniciado en http://localhost:8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}

// -------------------- Handlers --------------------

// ---------------- Dashboard ----------------
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/dashboard.html")
    if err != nil {
        log.Println("Error dashboard:", err)
        http.Error(w, "Error cargando dashboard", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// ---------------- Usuarios ----------------
type Usuario struct {
    ID       int
    Nombre   string
    Correo   string
    Password string
    CreadoEn string
}

func usuariosHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, nombre, correo, password, creado_en FROM usuarios")
    if err != nil {
        http.Error(w, "Error al obtener usuarios: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var usuarios []Usuario
    for rows.Next() {
        var u Usuario
        if err := rows.Scan(&u.ID, &u.Nombre, &u.Correo, &u.Password, &u.CreadoEn); err != nil {
            http.Error(w, "Error al leer usuarios: "+err.Error(), http.StatusInternalServerError)
            return
        }
        usuarios = append(usuarios, u)
    }

    tmpl, err := template.ParseFiles("templates/usuarios.html")
    if err != nil {
        log.Println("Error usuarios.html:", err)
        http.Error(w, "Error cargando plantilla", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, usuarios)
}

func agregarUsuarioHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }
    nombre := r.FormValue("nombre")
    correo := r.FormValue("correo")
    password := r.FormValue("password")
    creadoEn := time.Now().Format("2006-01-02 15:04:05")

    _, err := db.Exec("INSERT INTO usuarios (nombre, correo, password, creado_en) VALUES (?, ?, ?, ?)",
        nombre, correo, password, creadoEn)
    if err != nil {
        http.Error(w, "Error al insertar usuario: "+err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/usuarios", http.StatusSeeOther)
}

// ---------------- Productos ----------------
type Producto struct {
    ID       int
    Nombre   string
    Precio   float64
    CreadoEn string
}

func productosHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, nombre, precio, creado_en FROM productos")
    if err != nil {
        http.Error(w, "Error al obtener productos: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var productos []Producto
    for rows.Next() {
        var p Producto
        if err := rows.Scan(&p.ID, &p.Nombre, &p.Precio, &p.CreadoEn); err != nil {
            http.Error(w, "Error al leer productos: "+err.Error(), http.StatusInternalServerError)
            return
        }
        productos = append(productos, p)
    }

    tmpl, err := template.ParseFiles("templates/productos.html")
    if err != nil {
        log.Println("Error productos.html:", err)
        http.Error(w, "Error cargando plantilla", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, productos)
}

func agregarProductoHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }
    nombre := r.FormValue("nombre")
    precio, err := strconv.ParseFloat(r.FormValue("precio"), 64)
    if err != nil {
        http.Error(w, "Precio inválido", http.StatusBadRequest)
        return
    }
    creadoEn := time.Now().Format("2006-01-02 15:04:05")

    _, err = db.Exec("INSERT INTO productos (nombre, precio, creado_en) VALUES (?, ?, ?)", nombre, precio, creadoEn)
    if err != nil {
        http.Error(w, "Error al insertar producto: "+err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/productos", http.StatusSeeOther)
}

// ---------------- Facturas ----------------
type Factura struct {
    ID        int
    UsuarioID int
    Total     float64
    Fecha     string
}

type FacturasPage struct {
    Facturas []Factura
    Usuarios []Usuario
}

func facturasHandler(w http.ResponseWriter, r *http.Request) {
    // Facturas
    rows, err := db.Query("SELECT id, usuario_id, total, fecha FROM facturas")
    if err != nil {
        http.Error(w, "Error al obtener facturas: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var facturas []Factura
    for rows.Next() {
        var f Factura
        if err := rows.Scan(&f.ID, &f.UsuarioID, &f.Total, &f.Fecha); err != nil {
            http.Error(w, "Error al leer facturas: "+err.Error(), http.StatusInternalServerError)
            return
        }
        facturas = append(facturas, f)
    }

    // Usuarios para dropdown
    rowsU, err := db.Query("SELECT id, nombre, correo, password, creado_en FROM usuarios")
    if err != nil {
        http.Error(w, "Error al obtener usuarios: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rowsU.Close()

    var usuarios []Usuario
    for rowsU.Next() {
        var u Usuario
        if err := rowsU.Scan(&u.ID, &u.Nombre, &u.Correo, &u.Password, &u.CreadoEn); err != nil {
            http.Error(w, "Error al leer usuarios: "+err.Error(), http.StatusInternalServerError)
            return
        }
        usuarios = append(usuarios, u)
    }

    data := FacturasPage{
        Facturas: facturas,
        Usuarios: usuarios,
    }

    tmpl, err := template.ParseFiles("templates/facturas.html")
    if err != nil {
        log.Println("Error facturas.html:", err)
        http.Error(w, "Error cargando plantilla", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, data)
}

func agregarFacturaHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    usuarioID, err := strconv.Atoi(r.FormValue("usuario_id"))
    if err != nil {
        http.Error(w, "Usuario inválido", http.StatusBadRequest)
        return
    }

    total, err := strconv.ParseFloat(r.FormValue("total"), 64)
    if err != nil {
        http.Error(w, "Total inválido", http.StatusBadRequest)
        return
    }

    fecha := time.Now().Format("2006-01-02 15:04:05")
    _, err = db.Exec("INSERT INTO facturas (usuario_id, total, fecha) VALUES (?, ?, ?)", usuarioID, total, fecha)
    if err != nil {
        http.Error(w, "Error al insertar factura: "+err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/facturas", http.StatusSeeOther)
}
