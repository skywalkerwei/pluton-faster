package resultx

const defaultCode = 400

type CodeError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CodeErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewCode(code int, message string, data interface{}) error {
	return &CodeError{Code: code, Message: message, Data: data}
}

func NewDefault(message string, data interface{}) error {
	return NewCode(defaultCode, message, data)
}

func Error(code int, message string, data interface{}) error {
	return NewCode(code, message, data)
}

func Success(data interface{}) error {
	return NewCode(200, "ok", data)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) DataInfo() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
		Data:    e.Data,
	}
}
