package services

import (
	"context"
	"errors"

	"github.com/rulanugrh/gometri/entities/domain"
	"github.com/rulanugrh/gometri/entities/web"
	"github.com/rulanugrh/gometri/repository"
	"github.com/rulanugrh/gometri/utils"
)

type ticketservice struct {
	ticket repository.TicketInterface
}

func NewTicketServices(ticket repository.TicketInterface) TicketService {
	return &ticketservice{
		ticket: ticket,
	}
}

func (tick *ticketservice) CreateTicket(ctx context.Context, ticket domain.Ticket) (*web.TicketResponse, error) {
	// Create Span Parrent
	ctx, parentSpan := utils.Tracer.Start(ctx, "ticketApi")
	defer parentSpan.End()

	// Create Event for Ticket
	parentSpan.AddEvent("create new ticket id")

	// Create Ticket
	data, err := tick.ticket.CreateTicket(ctx, ticket)
	if err != nil {
		parentSpan.RecordError(err)
		return nil, err
	}

	// Create response for api
	response := web.TicketResponse{
		Name:     data.Name,
		Price:    data.Capacity,
		Capacity: data.Capacity,
	}

	return &response, nil
}

func (tick *ticketservice) FindTicket(ctx context.Context) (*[]domain.Ticket, error) {
	// Create Span Parrent
	ctx, parentSpan := utils.Tracer.Start(ctx, "findTicket")
	defer parentSpan.End()

	data, err := tick.ticket.FindTicket(ctx)
	if err != nil {
		// record Error
		parentSpan.RecordError(err)
		return nil, errors.New("cant find ticket")
	}

	return data, nil
}

func (tick *ticketservice) FindById(ctx context.Context, id uint) (*web.TicketResponse, error) {
	// Create Span Parrent
	ctx, parentSpan := utils.Tracer.Start(ctx, "findId")
	defer parentSpan.End()

	data, err := tick.ticket.FindById(ctx, id)
	if err != nil {
		return nil, errors.New("cannot find by id")
	}

	// Set attribute keyvalue for metric
	// attrs := []attribute.KeyValue{
	//	utils.TicketStatus.String(string(data.Name)),
	// }

	// Ticketcounter for metric
	utils.TicketCounter.Add(ctx, 1, nil)

	// Create Response API
	response := web.TicketResponse{
		Name:     data.Name,
		Price:    data.Price,
		Capacity: data.Capacity,
	}

	return &response, nil
}
