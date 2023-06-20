package repository

import (
	"context"
	"errors"

	"github.com/rulanugrh/gometri/config"
	"github.com/rulanugrh/gometri/entities/domain"
)

type ticketstruct struct{}

func TicketRepository() TicketInterface {
	return &ticketstruct{}
}

func (tick *ticketstruct) CreateTicket(ctx context.Context, ticket domain.Ticket) (*domain.Ticket, error) {
	err := config.DB.WithContext(ctx).Create(&ticket).Error
	if err != nil {
		return nil, errors.New("cant create ticket")
	}

	return &ticket, nil
}

func (tick *ticketstruct) FindTicket(ctx context.Context) (*[]domain.Ticket, error) {
	var ticket []domain.Ticket

	err := config.DB.WithContext(ctx).Find(&ticket).Error
	if err != nil {
		return nil, errors.New("cant find all ticket")
	}

	return &ticket, nil
}

func (tick *ticketstruct) FindById(ctx context.Context, id uint) (*domain.Ticket, error) {
	var ticket domain.Ticket
	err := config.DB.WithContext(ctx).First(&ticket, id)
	if err != nil {
		return nil, errors.New("cant find ticket by id")
	}

	return &ticket, nil
}
