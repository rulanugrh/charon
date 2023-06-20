package main

import (
	"github.com/rulanugrh/gometri/handlers"
	"github.com/rulanugrh/gometri/repository"
	"github.com/rulanugrh/gometri/routes"
	"github.com/rulanugrh/gometri/services"
)

func main() {
	ticketRepo := repository.TicketRepository()
	ticketServ := services.NewTicketServices(ticketRepo)
	ticketHandle := handlers.NewTicketHandlers(ticketServ)

	routes.Run(ticketHandle)
}
