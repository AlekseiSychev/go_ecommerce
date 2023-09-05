package controllers

type ErrorResponse struct {
	Error interface{} `json:"error"`
}

type ResponseCustom struct {
	Data interface{} `json:"data"`
}

type ResponseID struct {
	ID uint `json:"id"`
}