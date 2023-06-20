package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rulanugrh/gometri/entities/domain"
	"github.com/rulanugrh/gometri/entities/web"
	"github.com/rulanugrh/gometri/services"
)

type tickethandlers struct {
	ticketServ services.TicketService
}

func NewTicketHandlers(ticket services.TicketService) TicketHandlers {
	return &tickethandlers{
		ticketServ: ticket,
	}
}

func (tick *tickethandlers) CreateTicket(w http.ResponseWriter, r *http.Request) {
	upData, _ := ioutil.ReadAll(r.Body)

	var ticket domain.Ticket
	ctx := context.Background()

	json.Unmarshal(upData, &ticket)
	data, err := tick.ticketServ.CreateTicket(ctx, ticket)
	if err != nil {
		_response := web.FailureResponse{
			Code:    400,
			Message: "cant create ticket",
		}

		response, _ := json.Marshal(_response)
		w.WriteHeader(400)
		w.Write(response)
	}

	_success := web.SuccessResponse{
		Code:    200,
		Message: "success create ticket",
		Data:    data,
	}

	response, errJson := json.Marshal(_success)
	if errJson != nil {
		log.Fatal(errJson)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func (tick *tickethandlers) FindTicket(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	data, err := tick.ticketServ.FindTicket(ctx)
	if err != nil {
		_failure := web.FailureResponse{
			Code:    400,
			Message: "cant find all ticket",
		}

		response, _ := json.Marshal(_failure)
		w.WriteHeader(400)
		w.Write(response)
	}

	_success := web.SuccessResponse{
		Code:    200,
		Message: "find all ticket",
		Data:    data,
	}
	response, errJson := json.Marshal(_success)
	if errJson != nil {
		log.Fatal(errJson)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func (tick *tickethandlers) FindById(w http.ResponseWriter, r *http.Request) {
	getId := mux.Vars(r)
	stringId := getId["id"]
	Id, _ := strconv.Atoi(stringId)

	ctx := context.Background()

	data, err := tick.ticketServ.FindById(ctx, uint(Id))
	if err != nil {
		_failure := web.FailureResponse{
			Code:    400,
			Message: "cant find ticket by this id",
		}

		response, _ := json.Marshal(_failure)
		w.WriteHeader(400)
		w.Write(response)
	}

	_success := web.SuccessResponse{
		Code:    200,
		Message: "find ticket",
		Data:    data,
	}

	response, errJson := json.Marshal(_success)
	if errJson != nil {
		log.Fatal(errJson)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}
