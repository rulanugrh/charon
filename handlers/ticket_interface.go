package handlers

import "net/http"

type TicketHandlers interface {
	CreateTicket(w http.ResponseWriter, r *http.Request)
	FindTicket(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}
