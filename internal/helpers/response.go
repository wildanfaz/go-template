package helpers

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewResponse() *Response {
	return &Response{
		Error:   false,
		Message: "Unknown",
	}
}

func (r *Response) AsError() *Response {
	r.Error = true

	return r
}

func (r *Response) WithMessage(message string) *Response {
	r.Message = message

	return r
}

func (r *Response) WithData(data any) *Response {
	r.Data = data

	return r
}
