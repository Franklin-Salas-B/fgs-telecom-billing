package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

func main() {
	// Cargar templates HTML
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Servir archivos estáticos
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Rutas HTML
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/usuarios", usuariosHandler)
	http.HandleFunc("/productos", productosHandler)
	http.HandleFunc("/factura", facturaHandler)

	fmt.Println("Servidor corriendo en http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

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
