package system

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maoqide/melody"
	"servermanager/globel"
)

type ApiWs struct {
}

func init() {
	globel.Melody = melody.New()

	globel.Melody.HandleMessage(func(s *melody.Session, msg []byte) {
		//globel.Melody.BroadcastFilter(msg, func(q *melody.Session) bool {
		//	return q.Get("id") == s.Get("id")
		//})

		globel.Melody.Broadcast(msg)
	})
}

func (p *ApiWs) WsConnect(gc *gin.Context) {
	globel.Melody.HandleRequest(gc.Writer, gc.Request)
}

type message struct {
	Type string
	Msg  string
}

func (p *ApiWs) WebSokectTest(gc *gin.Context) {

	msgtest := message{}
	msgtest.Type = "sms"
	msgtest.Msg = "消息测试"
	ss, _ := json.Marshal(msgtest)
	globel.Melody.Broadcast(ss)

	gc.JSON(200, msgtest)
}
