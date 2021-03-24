package controller

import (
	"github.com/chanprogo/somemodule/app"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type MelodyController struct {
	app.WsController
}

func (t *MelodyController) Router(e *gin.Engine) {
	t.Melody = melody.New()
	group := e.Group("/api")

	// ws://192.168.56.2:8500/api/getExample
	group.GET("/getExample", func(ctx *gin.Context) {
		t.Melody.HandleRequest(ctx.Writer, ctx.Request)
	})

	t.Melody.HandleMessage(func(s *melody.Session, msg []byte) {
		ret := make(map[string]interface{})
		ret["state"] = 1
		t.Melody.BroadcastFilter(t.WsRespOK(ret), func(q *melody.Session) bool {
			return s == q
		})
	})

	t.Melody.HandleConnect(func(s *melody.Session) {
		ret := make(map[string]interface{})
		ret["state"] = 3
		t.Melody.BroadcastFilter(t.WsRespOK(ret), func(q *melody.Session) bool {
			return s == q
		})
	})

}
