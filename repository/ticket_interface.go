package repository

import (
	"context"

	"github.com/rulanugrh/gometri/entities/domain"
)

type TicketInterface interface {
	CreateTicket(ctx context.Context, ticket domain.Ticket) (*domain.Ticket, error)
	FindTicket(ctx context.Context) (*[]domain.Ticket, error)
	FindById(ctx context.Context, id uint) (*domain.Ticket, error)
}
