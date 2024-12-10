package main

import (
	"cutiecat6778/dont-trust-your-friend/handlers"
	"cutiecat6778/dont-trust-your-friend/lib"
	"cutiecat6778/dont-trust-your-friend/routers/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	handlers.NewValidator()
	handlers.InitDB()

	auth.InitRoutes(app)

	app.Run(":3000")

	lib.Log(-1, "Your app is running on port 3000")
}
