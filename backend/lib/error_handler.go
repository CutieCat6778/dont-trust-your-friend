package lib

import "fmt"

const (
	AuthService = iota
	UsersService
	JwtService
	RedisService
	DatabaseService
	SnowflakeService
)

type CustomError struct {
	Message string
	Code    int
	By      int
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) ParseToString() string {
	return fmt.Sprintf("Error: %s, Code: %d, By: %d", e.Message, e.Code, e.By)
}

func NewError(message string, code int, by int) *CustomError {
	err := &CustomError{
		Message: message,
		Code:    code,
		By:      by,
	}
	fmt.Println(err.ParseToString())
	return err
}
