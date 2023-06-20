package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/gometri/config"
	"github.com/rulanugrh/gometri/entities/domain"
	"github.com/rulanugrh/gometri/handlers"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Run(ticket handlers.TicketHandlers) {
	// PrintLog
	log.Println("Server running browh")

	// Migrate Database
	db := config.GetConnection()
	db.AutoMigrate(&domain.Ticket{})

	// Route Handle
	name := "TicketServ"
	api := mux.NewRouter().StrictSlash(true)

	api.Use(otelmux.Middleware(name))
	api.Handle("/ticket", otelhttp.WithRouteTag("/api/v1/ticket", otelhttp.NewHandler(http.HandlerFunc(ticket.CreateTicket), name))).Methods("POST")
	api.Handle("/ticket/{id}", otelhttp.WithRouteTag("/api/v1/ticket/{id}", otelhttp.NewHandler(http.HandlerFunc(ticket.FindById), name))).Methods("GET")
	api.Handle("/ticket", otelhttp.WithRouteTag("/api/v1/ticket", otelhttp.NewHandler(http.HandlerFunc(ticket.FindTicket), name))).Methods("GET")

	// Config Server
	conf := config.GetConfig()

	host := fmt.Sprintf("%s:%s", conf.Runn.Host, conf.Runn.Port)
	server := http.Server{
		Addr:    host,
		Handler: api,
	}

	server.ListenAndServe()
}
