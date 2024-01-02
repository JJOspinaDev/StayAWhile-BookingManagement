package repository

import (
	"database/sql"
	"stayawhile/microservices/bookingManagement/internal/models"
	"time"
)

type BookingRepository interface {
	Begin() (*sql.Tx, error)
	Save(tx *sql.Tx, booking *models.Reserva) error
	CheckRoomAvailability(tx *sql.Tx, habitacionID string, fechaEntrada, fechaSalida time.Time) (bool, error)
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) Begin() (*sql.Tx, error) {
    return r.db.Begin()
}

func (r *bookingRepository) Save(tx *sql.Tx, booking *models.Reserva) error {
	query := `
    INSERT INTO reservas 
    (
        cliente_id, 
        habitacion_id, 
        fecha_entrada, 
        fecha_salida, 
        numero_noches, 
        costo_total, 
        estado, 
        desayuno_incluido, 
        cama_extra, 
        transporte_aeropuerto, 
        notas
    ) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := tx.Exec(query, booking.ClienteID, booking.HabitacionID, booking.FechaEntrada, booking.FechaSalida, booking.NumeroNoches, booking.CostoTotal, booking.Estado, booking.DesayunoIncluido, booking.CamaExtra, booking.TransporteAeropuerto, booking.Notas)
	if err != nil {
		return err
	}

	return nil
}

func (r *bookingRepository) CheckRoomAvailability(tx *sql.Tx, habitacionID string, fechaEntrada, fechaSalida time.Time) (bool, error) {
	var count int
	query := `
        SELECT COUNT(*)
        FROM reservas
        WHERE habitacion_id = ?
        AND fecha_salida > ?
        AND fecha_entrada < ?
    `
	err := tx.QueryRow(query, habitacionID, fechaEntrada, fechaSalida).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
