package main

import (
	"stayawhile/microservices/bookingManagement/api"
	"database/sql"
	"log"
)

func main() {
	// Configurar la conexión a la base de datos
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
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
