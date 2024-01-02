package handlers

import (
	"encoding/json"
	"net/http"
	"stayawhile/microservices/bookingManagement/internal/handlers/dto"
	"stayawhile/microservices/bookingManagement/internal/models"
	"stayawhile/microservices/bookingManagement/internal/services"
)

type bookingHandler struct {
	bookingService services.BookingService
}

func NewBookingHandler(bookingService services.BookingService) *bookingHandler {
	return &bookingHandler{
		bookingService: bookingService,
	}
}

func (b *bookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {

	var request dto.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		// Manejar error de decodificación
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/*	ClienteNombre: request.ClienteNombre,
		ClienteEmail: request.ClienteEmail,
		ClienteTelefono: request.ClienteTelefono,
	*/

	reserva := models.Reserva{
		//ClienteID: cliente.ID, // Suponiendo que se obtuvo o creó el cliente

		HabitacionID:         request.HabitacionID,
		FechaEntrada:         request.FechaEntrada,
		FechaSalida:          request.FechaSalida,
		DesayunoIncluido:     request.DesayunoIncluido,
		CamaExtra:            request.CamaExtra,
		TransporteAeropuerto: request.TransporteAeropuerto,
	}

	err := b.bookingService.CreateBooking(&reserva)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
