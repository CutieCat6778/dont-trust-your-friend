package lib

const (
	AuthService = iota
	UsersService
	JwtService
)

type CustomError struct {
	Message string
	Code    int
	By      int
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewError(message string, code int, by int) *CustomError {
	return &CustomError{
		Message: message,
		Code:    code,
		By:      by,
	}
}
