package main

import (
	"cutiecat6778/dont-trust-your-friend/routers/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	auth.InitRoutes(app)

	app.Run(":3000")
}
