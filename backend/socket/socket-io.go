package socket

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/socket.io/v2/socket"
)

func InitSocketIo(e *gin.Engine) {
	c := socket.DefaultServerOptions()
	c.SetServeClient(true)
	c.SetPingInterval(1000)
	c.SetPingTimeout(5000)
	c.SetMaxHttpBufferSize(1000000)
	c.SetConnectTimeout(1000 * time.Millisecond)
	c.SetCors(&types.Cors{
		Origin:      []string{"http://localhost:5173"},
		Credentials: true,
	})

	socketio := socket.NewServer(nil, nil)

	e.GET("/socket", adapter.Wrap(func(h http.Handler) http.Handler {
		return socketio.ServeHandler(c)
	}))

	e.POST("/socket", adapter.Wrap(func(h http.Handler) http.Handler {
		return socketio.ServeHandler(c)
	}))
}
