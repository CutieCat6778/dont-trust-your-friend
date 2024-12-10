package lib

import (
	"fmt"
	"log"
)

// const (
// 	AuthService = iota
// 	UsersService
// 	JwtService
// 	RedisService
// 	DatabaseService
// 	SnowflakeService
// )

func Log(service int, message ...interface{}) {
	var serviceName string
	switch service {
	case 0:
		serviceName = "Auth Service"
	case 1:
		serviceName = "Users Service"
	case 2:
		serviceName = "Jwt Service"
	case 3:
		serviceName = "Redis Service"
	case 4:
		serviceName = "Database Service"
	case 5:
		serviceName = "Snowflake Service"
	default:
		serviceName = "System Service"
	}
	log.Printf(fmt.Sprintf("[%s] ", serviceName), message)
}
