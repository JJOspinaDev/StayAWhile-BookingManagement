package dto

import (
	"time"
)

type CreateBookingRequest struct {
	ClienteNombre        string    `json:"clienteNombre"`
	ClienteEmail         string    `json:"clienteEmail"`
	ClienteTelefono      string    `json:"clienteTelefono"`
	HabitacionID         int64     `json:"habitacionId"`
	FechaEntrada         time.Time `json:"fechaEntrada"`
	FechaSalida          time.Time `json:"fechaSalida"`
	DesayunoIncluido     bool      `json:"desayunoIncluido"`
	CamaExtra            bool      `json:"camaExtra"`
	TransporteAeropuerto bool      `json:"transporteAeropuerto"`
}
