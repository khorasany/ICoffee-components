package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func (e *RestErr) NewStatusNotAcceptable(message string) *RestErr {
	e.Message = message
	e.Status = http.StatusNotAcceptable
	e.Error = "not_accepted"
	return e
}
