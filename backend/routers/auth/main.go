package auth

import (
	"cutiecat6778/dont-trust-your-friend/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
		Password string `json:"password" validate:"required,min=8,max=20,ascii"`
		Name     string `json:"name" validate:"required,min=3,max=50,ascii"`
	}
)

func InitRoutes(app *gin.Engine) {
	app.POST("/login", login)
	app.POST("/register", register)
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errs := handlers.Validator.Validate(req); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "register",
	})
}
