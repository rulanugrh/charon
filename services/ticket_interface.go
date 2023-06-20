package services

import (
	"context"

	"github.com/rulanugrh/gometri/entities/domain"
	"github.com/rulanugrh/gometri/entities/web"
)

type TicketService interface {
	CreateTicket(ctx context.Context, ticket domain.Ticket) (*web.TicketResponse, error)
	FindTicket(ctx context.Context) (*[]domain.Ticket, error)
	FindById(ctx context.Context, id uint) (*web.TicketResponse, error)
}
