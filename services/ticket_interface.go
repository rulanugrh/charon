package services

import (
	"github.com/rulanugrh/gometri/entities/domain"
	"github.com/rulanugrh/gometri/entities/web"
)

type TicketService interface {
	CreateTicket(ticket domain.Ticket) (*web.TicketResponse, error)
	FindTicket() (*[]domain.Ticket, error)
	FindById(id uint) (*web.TicketResponse, error)
}
