package repository

import "github.com/rulanugrh/gometri/entities/domain"

type TicketInterface interface {
	CreateTicket(ticket domain.Ticket) (*domain.Ticket, error)
	FindTicket() (*[]domain.Ticket, error)
	FindById(id uint) (*domain.Ticket, error)
}
