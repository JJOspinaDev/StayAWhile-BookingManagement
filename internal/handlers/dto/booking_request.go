package dto

type CreateBookingRequest struct {
	ClienteNombre        string    `json:"clienteNombre"`
	ClienteEmail         string    `json:"clienteEmail"`
	ClienteTelefono      string    `json:"clienteTelefono"`
	HabitacionID         int64     `json:"habitacionId"`
	FechaEntrada         string    `json:"fechaEntrada"`
	FechaSalida          string    `json:"fechaSalida"`
	DesayunoIncluido     bool      `json:"desayunoIncluido"`
	CamaExtra            bool      `json:"camaExtra"`
	TransporteAeropuerto bool      `json:"transporteAeropuerto"`
}
