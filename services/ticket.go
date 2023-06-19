package services

import (
	"errors"

	"github.com/rulanugrh/gometri/entities/domain"
	"github.com/rulanugrh/gometri/entities/web"
	"github.com/rulanugrh/gometri/repository"
)

type ticketservice struct {
	ticket repository.TicketInterface
}

func NewTicketServices(ticket repository.TicketInterface) TicketService {
	return &ticketservice{
		ticket: ticket,
	}
}

func (tick *ticketservice) CreateTicket(ticket domain.Ticket) (*web.TicketResponse, error) {
	data, err := tick.ticket.CreateTicket(ticket)
	if err != nil {
		return nil, errors.New("cant create ticket")
	}

	response := web.TicketResponse{
		Name:     data.Name,
		Price:    data.Capacity,
		Capacity: data.Capacity,
	}

	return &response, nil
}

func (tick *ticketservice) FindTicket() (*[]domain.Ticket, error) {
	data, err := tick.ticket.FindTicket()
	if err != nil {
		return nil, errors.New("cant find ticket")
	}

	return data, nil
}

func (tick *ticketservice) FindById(id uint) (*web.TicketResponse, error) {
	data, err := tick.ticket.FindById(id)
	if err != nil {
		return nil, errors.New("cannot find by id")
	}

	response := web.TicketResponse{
		Name:     data.Name,
		Price:    data.Price,
		Capacity: data.Capacity,
	}

	return &response, nil
}
