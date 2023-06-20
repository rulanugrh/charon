package utils

import (
	"strconv"

	"github.com/rulanugrh/gometri/entities/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var Tracer = otel.Tracer("github.com/rulanugrh/gometri/ticket")
var Meter = otel.Meter("github.com/rulanugrh/gometri/ticket")

var ticket domain.Ticket
var TicketCounter, _ = Meter.Int64Counter("ticket", metric.WithDescription("number ticketing with status"), metric.WithUnit(strconv.Itoa(int(ticket.ID))))

var TicketStatus = attribute.Key("status")
