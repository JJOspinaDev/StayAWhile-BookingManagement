package repository

import (
	"database/sql"
	"fmt"
	"stayawhile/microservices/bookingManagement/internal/models"
	"time"
)

type BookingRepository interface {
	Begin() (*sql.Tx, error)
	Save(tx *sql.Tx, booking *models.Reserva) error
	CheckRoomAvailability(tx *sql.Tx, habitacionID int64, fechaEntrada, fechaSalida time.Time) (bool, error)
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
        clienteId, 
        habitacionId, 
        fechaEntrada, 
        fechaSalida, 
        numeroNoches, 
        costoTotal, 
        estado, 
        desayunoIncluido, 
        camaExtra, 
        transporteAeropuerto, 
        notas
    ) 
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	fmt.Println("HOLAA")
	_, err := tx.Exec(query, booking.ClienteID, booking.HabitacionID, booking.FechaEntrada, booking.FechaSalida, booking.NumeroNoches, booking.CostoTotal, booking.Estado, booking.DesayunoIncluido, booking.CamaExtra, booking.TransporteAeropuerto, booking.Notas)
	if err != nil {
		return err
	}

	return nil
}

func (r *bookingRepository) CheckRoomAvailability(tx *sql.Tx, habitacionID int64, fechaEntrada, fechaSalida time.Time) (bool, error) {
	var count int
	query := `
        SELECT COUNT(*)
        FROM reservas
        WHERE habitacionId = ?
        AND fechaSalida > ?
        AND fechaEntrada < ?
    `
	err := tx.QueryRow(query, habitacionID, fechaEntrada, fechaSalida).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
