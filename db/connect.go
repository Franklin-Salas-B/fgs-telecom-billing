package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	// DSN (cadena de conexión)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Abrir conexión
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al abrir la BD:", err)
	}

	// Verificar conexión
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error al hacer ping a la BD:", err)
	}

	fmt.Println("✅ Conexión exitosa a MySQL")
}
