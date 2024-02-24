package main

import (
	"database/sql"
	"stayawhile/microservices/bookingManagement/api"

	_ "github.com/go-sql-driver/mysql"

	"log"
)

func main() {
	// Configurar la conexión a la base de datos
	db, err := sql.Open("mysql", "root:mipassword@tcp(localhost:3307)/miBaseDeDatos")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Verificar la conexión a la base de datos
	if err := db.Ping(); err != nil {
		log.Fatal("Error al realizar ping a la base de datos:", err)
	}

	// Configurar y arrancar el servidor
	router := api.SetupRouter(db)
	router.Run(":8080")
}
