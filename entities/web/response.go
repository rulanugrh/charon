package web

type TicketResponse struct {
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price"`
	Capacity int    `json:"capacity" form:"capacity"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FailureResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
