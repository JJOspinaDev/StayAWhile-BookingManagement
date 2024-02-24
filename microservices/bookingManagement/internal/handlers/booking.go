package handlers

import (
	"net/http"
	"stayawhile/microservices/bookingManagement/internal/handlers/dto"
	"stayawhile/microservices/bookingManagement/internal/models"
	"stayawhile/microservices/bookingManagement/internal/services"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	DateFormat = "2006-01-02"
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

	// Parsear las cadenas de fecha a time.Time utilizando la constante DateFormat
	fechaEntrada, err := time.Parse(DateFormat, request.FechaEntrada)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear la fecha de entrada"})
		return
	}

	fechaSalida, err := time.Parse(DateFormat, request.FechaSalida)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear la fecha de salida"})
		return
	}

	reserva := models.Reserva{
		HabitacionID:         request.HabitacionID,
		FechaEntrada:         fechaEntrada,
		FechaSalida:          fechaSalida,
		DesayunoIncluido:     request.DesayunoIncluido,
		CamaExtra:            request.CamaExtra,
		TransporteAeropuerto: request.TransporteAeropuerto,
	}

	// Aquí asumimos que CreateBooking manejará la lógica de buscar o crear el cliente
	err = b.bookingService.CreateBooking(&reserva, &cliente)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reserva)
}
