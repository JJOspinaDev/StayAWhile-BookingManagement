package services

import (
	"database/sql"
	"fmt"
	"stayawhile/microservices/bookingManagement/internal/models"
	"stayawhile/microservices/bookingManagement/internal/repository"
)

// BookingService define las operaciones disponibles para los bookingos.
type BookingService interface {
	CreateBooking(booking *models.Reserva, clienteInfo *models.Cliente) error
}

type bookingService struct {
	bookingRepo repository.BookingRepository
	clientRepo  repository.ClientRepository
}

func NewBookingService(bookingRepo repository.BookingRepository, clientRepo repository.ClientRepository) *bookingService {
	return &bookingService{
		bookingRepo: bookingRepo,
		clientRepo:  clientRepo,
	}
}

func (s *bookingService) CreateBooking(booking *models.Reserva, clienteInfo *models.Cliente) error {
	// Iniciar transacción

	fmt.Println("Hjsjdj")
	tx, err := s.bookingRepo.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Buscar cliente por email o alguna otra identificación única
	cliente, err := s.clientRepo.FindClienteByEmail(tx, clienteInfo.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if cliente == nil {
		// El cliente no existe, crear uno nuevo
		err := s.clientRepo.CreateCliente(tx, clienteInfo)
		if err != nil {
			return err
		}
		booking.ClienteID = clienteInfo.ID
	} else {
		// El cliente ya existe
		booking.ClienteID = cliente.ID
	}

	// Usar tx en las operaciones del repositorio
	available, err := s.bookingRepo.CheckRoomAvailability(tx, booking.HabitacionID, booking.FechaEntrada, booking.FechaSalida)
	if err != nil || !available {
		return err
	}

	if err = s.bookingRepo.Save(tx, booking); err != nil {
		return err
	}

	// Si todo va bien, hacer commit de la transacción
	return tx.Commit()
}
