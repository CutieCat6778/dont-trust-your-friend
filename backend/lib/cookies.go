package lib

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"fmt"
)

func initCors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"GET", "POST", "PUT"},
			AllowHeaders:     []string{"Origin, Content-Type, Accept"},
			AllowCredentials: true,
			MaxAge:           86400,
		},
	)
}

func CreateCookie(c *gin.Context, jwt *Jwt) *CustomError {
	c.SetCookie("access_token", jwt.AccessToken, 54000, "/", "localhost", false, true)
	c.SetCookie("refresh_token", jwt.RefreshToken, 1512000, "/", "localhost", false, true)
	c.SetCookie("created_at", fmt.Sprintf("%d", jwt.CreatedAt.Unix()), 54000, "/", "localhost", false, true)
	return nil
}
