package rest_err

import "net/http"

// RestErr representa um erro padronizado para respostas HTTP.
type RestErr struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes"`
}

// Cause representa detalhes adicionais sobre um erro.
type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Erro implementa a interface de erro para RestErr.
func (r *RestErr) Error() string {
	return r.Message
}

// NewRestErr cria uma nova instância de RestErr.
func NewRestErr(message, err string, code int, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// NewBadRequestError cria um erro de bad request.
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  nil,
	}
}

// NewBadRequestValidationError cria um erro de validação para bad request.
func NewBadRequestValidationError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// NewInternalServerError cria um erro de servidor interno.
func NewInternalServerError(message string, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  causes,
	}
}

// NewNotFoundError cria um erro para recurso não encontrado.
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

// NewForbiddenError cria um erro de acesso proibido.
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
