/*
 * @Author: ww
 * @Date: 2022-06-13 03:18:01
 * @Description:
 * @FilePath: \steam-concurrent-users\main.go
 */
package main

import (
	"log"
	"steam-concurrent-users/config"
	"steam-concurrent-users/handle"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/websocket"
)

func Init() {
	err := config.Init()
	if err != nil {
		log.Printf("config read err : %v", err)
	}
	handle.Init()

}



func main() {

	Init()


	ws, err := handle.Api.WS(handle.Ctx, nil, "")
	if err != nil {
		log.Printf("%+v, err:%v", ws, err)
		return
	}

	var msg event.ATMessageEventHandler = handle.MessageEventHandler

	handler := websocket.RegisterHandlers(msg)

	botgo.NewSessionManager().Start(ws, config.UToken, &handler)


}
