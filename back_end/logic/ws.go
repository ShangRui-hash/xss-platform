package logic

import (
	"xss/dao/memory"
	"xss/pkg/ws"

	"github.com/gin-gonic/gin"
)

//Ws websocket 业务逻辑
func Ws(c *gin.Context, userID int64) error {
	//1.建立连接
	conn, err := ws.NewConnect(c)
	if err != nil {
		return err
	}
	//2.将连接添加到该用户维护的连接数组中
	memory.AddConn(userID, conn)
	//3.启动一个连接存活检测协程
	go ws.CheckAlive(userID, conn)
	//4.检查该用户是否已经有一个消息推送协程
	_, ok := memory.HasSendMessageGoroutineMap.Load(userID)
	if !ok {
		go ws.SendMessage(userID)
		memory.HasSendMessageGoroutineMap.Store(userID, true)
	}
	return nil
}
