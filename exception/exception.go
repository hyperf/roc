package exception

type ExceptionInterface interface {
	error
	GetMessage() string
	GetCode() int
}
type Exception struct {
	Code    int
	Message string
}

func (e *Exception) Error() string {
	return e.Message
}

func (e *Exception) GetMessage() string {
	return e.Message
}

func (e *Exception) GetCode() int {
	return e.Code
}

func NewDefaultException(message string) ExceptionInterface {
	return &Exception{Code: SERVER_ERROR, Message: message}
}
