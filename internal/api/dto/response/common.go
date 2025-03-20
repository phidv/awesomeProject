package response

type CommonResponse struct {
	Data  interface{}    `json:"data,omitempty"`
	Error *ErrorResponse `json:"error,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int          `json:"code"`
	Errors  []FieldError `json:"errors,omitempty"`
	Message string       `json:"message,omitempty"`
}
