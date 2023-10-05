package api_body

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status       string      `json:"status"`
	ErrorMessage interface{} `json:"errorMessage"`
}
