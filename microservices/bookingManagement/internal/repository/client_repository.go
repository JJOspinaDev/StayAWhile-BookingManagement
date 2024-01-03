package repository

import (
	"database/sql"
	"stayawhile/microservices/bookingManagement/internal/models"
)

type ClientRepository interface {
	FindClienteByEmail(tx *sql.Tx, email string) (*models.Cliente, error)
	CreateCliente(tx *sql.Tx, cliente *models.Cliente) error
}

type clientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) ClientRepository {
	return &clientRepository{db: db}
}

func (c *clientRepository) CreateCliente(tx *sql.Tx, cliente *models.Cliente) error {
	query := "INSERT INTO clientes (nombre, email, telefono) VALUES (?, ?, ?)"
	result, err := tx.Exec(query, cliente.Nombre, cliente.Email, cliente.Telefono)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // Asumiendo que 'id' es una columna autoincrementable
	if err != nil {
		return err
	}

	cliente.ID = id // Actualiza el ID del cliente con el valor generado por la base de datos
	return nil
}

func (c *clientRepository) FindClienteByEmail(tx *sql.Tx, email string) (*models.Cliente, error) {
	var cliente models.Cliente
	query := "SELECT id, nombre, email, telefono FROM clientes WHERE email = ?"
	err := tx.QueryRow(query, email).Scan(&cliente.ID, &cliente.Nombre, &cliente.Email, &cliente.Telefono)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No se encontró el cliente
		}
		return nil, err // Error de base de datos
	}

	return &cliente, nil
}
