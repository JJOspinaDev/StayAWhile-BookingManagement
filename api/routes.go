package api

import (
	"database/sql"
	"stayawhile/microservices/bookingManagement/internal/handlers"
	"stayawhile/microservices/bookingManagement/internal/repository"
	"stayawhile/microservices/bookingManagement/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// Crear instancias del servicio y del handler
	bookingRepo := repository.NewBookingRepository(db) // Asumiendo que db es tu conexión a la base de datos
	clientRepo := repository.NewClientRepository(db)
	bookingService := services.NewBookingService(bookingRepo, clientRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	// Definir grupo de rutas para reservas
	app := router.Group("/booking")
	{
		// Ruta para crear una reserva
		app.POST("/", bookingHandler.CreateBooking)
		// Ruta para obtener detalles de una reserva específica por ID
		app.GET("/reservas", bookingHandler.GetAllBookings)
		// Ruta para actualizar una reserva existente por ID
		//app.PUT("/reservas/:id", bookingHandler.UpdateBookingByID)
		// Ruta para cancelar una reserva por ID
		//app.DELETE("/reservas/:id", bookingHandler.DeleteBookingByID)
	}
	return router

}

