package models

import "time"

// Cliente representa la información del cliente que realiza la reserva.
type Cliente struct {
	ID       int64  `json:"id"`
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
}

// Habitacion representa los detalles de la habitación reservada.
type Habitacion struct {
	ID             int64   `json:"id"`
	Tipo           string  `json:"tipo"`
	Descripcion    string  `json:"descripcion"`
	PrecioPorNoche float64 `json:"precioPorNoche"`
}

// Reserva define la estructura para una reserva de hotel.
type Reserva struct {
	ID                   int64     `json:"id"`
	ClienteID            int64     `json:"clienteId"`    // Clave foránea que referencia a Cliente
	HabitacionID         int64     `json:"habitacionId"` // Clave foránea que referencia a Habitacion
	FechaEntrada         time.Time `json:"fechaEntrada"`
	FechaSalida          time.Time `json:"fechaSalida"`
	NumeroNoches         int       `json:"numeroNoches"`
	CostoTotal           float64   `json:"costoTotal"`
	Estado               string    `json:"estado"`
	DesayunoIncluido     bool      `json:"desayunoIncluido"`
	CamaExtra            bool      `json:"camaExtra"`
	TransporteAeropuerto bool      `json:"transporteAeropuerto"`
	Notas                string    `json:"notas"`
	FechaCreacion        time.Time `json:"fechaCreacion"`
	FechaActualizacion   time.Time `json:"fechaActualizacion"`
}
