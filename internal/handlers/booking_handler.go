package handlers

import (
	"net/http"
	"stayawhile/internal/handlers/dto"
	"stayawhile/internal/models"
	"stayawhile/internal/services"
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

func (b *bookingHandler) GetAllBookings(c *gin.Context) {
	// Obtener parámetros de consulta opcionales
	clienteID := c.Query("clienteId")     // Ejemplo: /booking/reservas?clienteId=1
	fechaInicio := c.Query("fechaInicio") // Ejemplo: /booking/reservas?fechaInicio=2023-01-01
	fechaFin := c.Query("fechaFin")       // Ejemplo: /booking/reservas?fechaFin=2023-01-31

	// Utilizar el servicio para obtener reservas con los filtros aplicados
	filtros := make(map[string]interface{})
	if clienteID != "" {
		filtros["clienteId"] = clienteID
	}
	if fechaInicio != "" && fechaFin != "" {
		filtros["fechaInicio"] = fechaInicio
		filtros["fechaFin"] = fechaFin
	}

	reservas, err := b.bookingService.GetBookingsWithFilters(filtros)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la lista de reservas"})
		return
	}

	// Si se obtienen las reservas, devolverlas en la respuesta
	c.JSON(http.StatusOK, reservas)
}
