package transport

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"` //! flexibel data
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Count struct {
	Count int64 `json:"count"`
}

type ResponseBookList struct {
	Meta Meta        `json:"meta"`
	Count Count `json:"count"`
	Data interface{} `json:"data"` //! flexibel data
}

func ApiResponse(message string, code int, status string, data interface{}) Response {
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

func ApiResponseGetListBook(message string, code int, status string, data interface{}, books int64) ResponseBookList {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	count := Count{
		Count: books,
	}

	jsonResponse := ResponseBookList{
		Meta: meta,
		Count: count,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) { //!mnegubah terlebih dahulu menjadi validation error
		errors = append(errors, e.Error()) //! menambahkan nilai errornya
	}

	return errors
}