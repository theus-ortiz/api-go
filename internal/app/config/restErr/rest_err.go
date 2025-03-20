package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string
	Message string
}

// Error é um método que retorna a mensagem do struct RestErr
func (r *RestErr) Error() string {
	return r.Message
}

// Constantes para mensagens de erro
const (
	ErrBadRequest          = "bad_request"
	ErrInternalServerError = "internal_server_error"
	ErrNotFound            = "not_found"
	ErrForbidden           = "forbidden"
	ErrEmailAlreadyExists  = "email_already_exists"
	ErrUnauthorized        = "unauthorized"
)

// Funções para criar novos erros
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return NewRestErr(message, ErrBadRequest, http.StatusBadRequest, nil)
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, ErrBadRequest, http.StatusBadRequest, causes)
}

func NewInternalServerError(message string) *RestErr {
	return NewRestErr(message, ErrInternalServerError, http.StatusInternalServerError, nil)
}

func NewNotFoundError(message string) *RestErr {
	return NewRestErr(message, ErrNotFound, http.StatusNotFound, nil)
}

func NewForbiddenError(message string) *RestErr {
	return NewRestErr(message, ErrForbidden, http.StatusForbidden, nil)
}

func NewEmailAlreadyExistsError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, ErrEmailAlreadyExists, http.StatusConflict, causes)
}

func NewUnauthorizedError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, ErrUnauthorized, http.StatusUnauthorized, causes)
}
