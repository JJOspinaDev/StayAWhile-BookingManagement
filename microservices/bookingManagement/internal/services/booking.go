package services

import (
	"stayawhile/microservices/bookingManagement/internal/models"
	"stayawhile/microservices/bookingManagement/internal/repository"
)

// BookingService define las operaciones disponibles para los bookingos.
type BookingService interface {
	CreateBooking(booking *models.Reserva) error
}

type bookingService struct {
	repo repository.BookingRepository
}

// NewBookingService crea una nueva instancia de BookingService.
func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{
		repo: repo,
	}
}

func (s *bookingService) CreateBooking(booking *models.Reserva) error {
	// Iniciar transacción
	tx, err := s.repo.Begin()
	if err != nil {
		return err
	}

	// Asegurarse de hacer rollback en caso de error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Usar tx en las operaciones del repositorio
	available, err := s.repo.CheckRoomAvailability(tx, booking.HabitacionID, booking.FechaEntrada, booking.FechaSalida)
	if err != nil || !available {
		return err
	}

	if err = s.repo.Save(tx, booking); err != nil {
		return err
	}

	// Si todo va bien, hacer commit de la transacción
	return tx.Commit()
}
