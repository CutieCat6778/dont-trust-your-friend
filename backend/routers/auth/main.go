package auth

import (
	"cutiecat6778/dont-trust-your-friend/handlers"
	"cutiecat6778/dont-trust-your-friend/lib"
	"cutiecat6778/dont-trust-your-friend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	RegisterRequest struct {
		Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
		Password string `json:"password" validate:"required,min=8,max=20,ascii"`
		Name     string `json:"name" validate:"required,min=3,max=50,ascii"`
	}
	LoginRequest struct {
		Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
		Password string `json:"password" validate:"required,min=8,max=20,ascii"`
	}
)

func InitRoutes(app *gin.Engine) {
	auth := app.Group("/auth")
	auth.POST("/login", login)
	auth.PUT("/register", register)
}

func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if errs := handlers.VHandler.Validate(req); len(errs) > 0 && errs[0].Error {
		fmt.Println(errs)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs,
			"by":    "validator",
		})
		return
	}

	db := <-handlers.DB
	defer func() {
		handlers.DB <- db
	}()

	user, err := db.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	if !lib.CompareHashAndString(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password",
		})
		return
	}

	jwt, err := lib.SignJWT(user.ID)
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"tokens":  jwt,
		"user":    user,
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

	if errs := handlers.VHandler.Validate(req); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs,
		})
		return
	}

	db := <-handlers.DB
	defer func() {
		handlers.DB <- db
	}()

	hashedPassword := lib.HashString(req.Password)

	newUser := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		Name:     req.Name,
		Balance:  100,
	}

	if err := db.CreateUser(newUser); err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
