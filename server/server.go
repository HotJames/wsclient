package server

import (
	"github.com/HotJames/wsclient/gindefault"
	"github.com/HotJames/wsclient/wsserver"
	"github.com/gin-gonic/gin"
)

func StartServer(addr string, initial func(engine *gin.Engine, ws *wsserver.WsEngine)) {
	gindefault.Run(addr, func(engine *gin.Engine) {

		engine.GET("/healthz", func(c *gin.Context) {

			c.String(200, "ok")

		})

		wsEngine := wsserver.NewServer()

		engine.GET("/ws", wsserver.WsReadHandler(wsEngine))

		initial(engine, wsEngine)

	})
}
