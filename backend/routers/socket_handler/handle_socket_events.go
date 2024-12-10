package socket_handler

import (
	"cutiecat6778/dont-trust-your-friend/lib"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zishang520/socket.io/v2/socket"
)

func HandleSocketEvents(s socket.Server) {
	s.On("connection", OnConnection)
}

func OnConnection(clients ...interface{}) {
	client := clients[0].(*socket.Socket)

	claims := _authSocketRequest(client)
	if claims == nil {
		return
	}

	client.On("disconnect", func(clients ...interface{}) {
		client.Disconnect(true)
	})

	client.On("error", func(clients ...interface{}) {
		client.Disconnect(true)
	})

	client.On("message", func(clients ...interface{}) {
		client.Send("message", "Hello, World!")
	})
}

func _authSocketRequest(client *socket.Socket) *jwt.MapClaims {
	token, found := client.Request().Query().Get("token")
	if !found {
		client.Emit("error", "Token not found")
		client.Disconnect(true)
		return nil
	}

	claims, err := lib.DecodeJWT(token)
	if err != nil {
		client.Emit("error", err.Message)
		client.Disconnect(true)
		return nil
	}

	return &claims
}
