package helper

import "github.com/go-playground/validator/v10"

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type ResponseCreate struct {
	Message string `json:"message"`
}

func APIResponseCreate(message string) ResponseCreate {

	jsonResponse := ResponseCreate{
		Message: message,
	}

	return jsonResponse
}

type ResponseError struct {
	Message      string      `json:"message"`
	ErrorKey     string      `json:"error_key"`
	ErrorMessage string      `json:"error_message"`
	ErrorData    interface{} `json:"error_data"`
}

func APIResponseError(message string, errorKey string, errorMessage string, errorData interface{}) ResponseError {

	jsonResponse := ResponseError{
		Message:      message,
		ErrorKey:     errorKey,
		ErrorMessage: errorMessage,
		ErrorData:    errorData,
	}

	return jsonResponse
}

type ResponseGet struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func APIResponseGet(message string, data interface{}) ResponseGet {

	jsonResponse := ResponseGet{
		Message: message,
		Data:    data,
	}

	return jsonResponse
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
