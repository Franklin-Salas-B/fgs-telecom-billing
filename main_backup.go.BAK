package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// Configuración de la base de datos
const (
	dbUser     = "ecom_user"
	dbPassword = "P1ch1nch4"
	dbName     = "ecommerce_db"
)

// Variables globales
var db *sql.DB
var templates *template.Template

func main() {
	// Conectar a MySQL
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPassword, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a MySQL:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("No se puede conectar a la base de datos:", err)
	}
	fmt.Println("Conexión MySQL exitosa")

	// Cargar templates HTML
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Servir archivos estáticos (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas HTML
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/usuarios", usuariosHandler)
	http.HandleFunc("/productos", productosHandler)
	http.HandleFunc("/factura", facturaHandler)

	// Rutas API (temporal)
	http.HandleFunc("/api/carrito", carritoAPIHandler)
	http.HandleFunc("/api/checkout", checkoutAPIHandler)

	fmt.Println("Servidor iniciado en http://0.0.0.0:8081")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}

// Función para renderizar templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handlers HTML
func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func usuariosHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "usuarios.html", nil)
}

func productosHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "productos.html", nil)
}

func facturaHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "factura.html", nil)
}

// Handlers API (temporal)
func carritoAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: conectar a la tabla carrito en MySQL
	w.Write([]byte(`[]`))
}

func checkoutAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: procesar el checkout, generar orden y detalle
	w.Write([]byte(`{"status":"ok"}`))
}
