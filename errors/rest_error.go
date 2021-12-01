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

func (e *RestErr) NewStatusNoContent(message string) *RestErr {
	e.Message = message
	e.Status = http.StatusNoContent
	e.Error = "empty_content"
	return e
}

func (e *RestErr) NewStatusInternalServerError(message string) *RestErr {
	e.Message = message
	e.Status = http.StatusInternalServerError
	e.Error = "internal_server_error"
	return e
}

func (e *RestErr) NewStatusMethodNotAllowed(message string) *RestErr {
	e.Message = message
	e.Status = http.StatusMethodNotAllowed
	e.Error = "method_not_allowed"
	return e
}
