package handlers

import (
	"net/http"
	"stayawhile/microservices/bookingManagement/internal/handlers/dto"
	"stayawhile/microservices/bookingManagement/internal/models"
	"stayawhile/microservices/bookingManagement/internal/services"

	"github.com/gin-gonic/gin"
)

type bookingHandler struct {
	bookingService services.BookingService
}

func NewBookingHandler(bookingService services.BookingService) *bookingHandler {
	return &bookingHandler{
		bookingService: bookingService,
	}
}

func (b *bookingHandler) CreateBooking(c *gin.Context) {
	var request dto.CreateBookingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cliente := models.Cliente{
		Nombre:   request.ClienteNombre,
		Email:    request.ClienteEmail,
		Telefono: request.ClienteTelefono,
	}

	reserva := models.Reserva{
		HabitacionID:         request.HabitacionID,
		FechaEntrada:         request.FechaEntrada,
		FechaSalida:          request.FechaSalida,
		DesayunoIncluido:     request.DesayunoIncluido,
		CamaExtra:            request.CamaExtra,
		TransporteAeropuerto: request.TransporteAeropuerto,
	}

	// Aquí asumimos que CreateBooking manejará la lógica de buscar o crear el cliente
	err := b.bookingService.CreateBooking(&reserva, &cliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reserva)
}
