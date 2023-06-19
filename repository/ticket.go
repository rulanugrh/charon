package repository

import (
	"errors"

	"github.com/rulanugrh/gometri/config"
	"github.com/rulanugrh/gometri/entities/domain"
)

type ticketstruct struct{}

func TicketRepository() TicketInterface {
	return &ticketstruct{}
}

func (tick *ticketstruct) CreateTicket(ticket domain.Ticket) (*domain.Ticket, error) {
	err := config.DB.Create(&ticket).Error
	if err != nil {
		return nil, errors.New("cant create ticket")
	}

	return &ticket, nil
}

func (tick *ticketstruct) FindTicket() (*[]domain.Ticket, error) {
	var ticket []domain.Ticket

	err := config.DB.Find(&ticket).Error
	if err != nil {
		return nil, errors.New("cant find all ticket")
	}

	return &ticket, nil
}

func (tick *ticketstruct) FindById(id uint) (*domain.Ticket, error) {
	var ticket domain.Ticket
	err := config.DB.First(&ticket, id)
	if err != nil {
		return nil, errors.New("cant find ticket by id")
	}

	return &ticket, nil
}
