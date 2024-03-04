package repository

import (
	"database/sql"
	"fmt"
	"stayawhile/microservices/bookingManagement/internal/models"
	"strings"
	"time"
)

type BookingRepository interface {
	Begin() (*sql.Tx, error)
	Save(tx *sql.Tx, booking *models.Reserva) error
	CheckRoomAvailability(tx *sql.Tx, habitacionID int64, fechaEntrada, fechaSalida time.Time) (bool, error)
	GetBookingsWithFilters(filtros map[string]interface{}) ([]models.Reserva, error)
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

// GetBookingsWithFilters busca reservas con filtros opcionales como clienteId, fechaInicio y fechaFin.
func (r *bookingRepository) GetBookingsWithFilters(filtros map[string]interface{}) ([]models.Reserva, error) {
	var reservas []models.Reserva

	// Comenzamos con la consulta base
	baseQuery := "SELECT id, clienteId, habitacionId, fechaEntrada, fechaSalida, numeroNoches, costoTotal, estado, desayunoIncluido, camaExtra, transporteAeropuerto, notas, fechaCreacion, fechaActualizacion FROM reservas WHERE 1=1"

	var args []interface{}
	var conditions []string

	// Agregamos condiciones según los filtros
	if clienteId, ok := filtros["clienteId"]; ok {
		conditions = append(conditions, "clienteId = ?")
		args = append(args, clienteId)
	}
	if fechaInicio, ok := filtros["fechaInicio"]; ok {
		conditions = append(conditions, "fechaEntrada >= ?")
		args = append(args, fechaInicio)
	}
	if fechaFin, ok := filtros["fechaFin"]; ok {
		conditions = append(conditions, "fechaSalida <= ?")
		args = append(args, fechaFin)
	}

	// Construimos la consulta final
	finalQuery := baseQuery
	if len(conditions) > 0 {
		finalQuery += " AND " + strings.Join(conditions, " AND ")
	}

	// Ejecutamos la consulta
	rows, err := r.db.Query(finalQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta de reservas: %w", err)
	}
	defer rows.Close()

	// Procesamos los resultados
	for rows.Next() {
		var reserva models.Reserva
		if err := rows.Scan(&reserva.ID, &reserva.ClienteID, &reserva.HabitacionID, &reserva.FechaEntrada, &reserva.FechaSalida, &reserva.NumeroNoches, &reserva.CostoTotal, &reserva.Estado, &reserva.DesayunoIncluido, &reserva.CamaExtra, &reserva.TransporteAeropuerto, &reserva.Notas, &reserva.FechaCreacion, &reserva.FechaActualizacion); err != nil {
			return nil, fmt.Errorf("error al leer fila de reservas: %w", err)
		}
		reservas = append(reservas, reserva)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error después de leer todas las filas de reservas: %w", err)
	}

	return reservas, nil
}