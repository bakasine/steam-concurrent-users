package handle

import (
	"context"
	"fmt"
	"log"
	"time"

	"steam-concurrent-users/config"
	"steam-concurrent-users/steam"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
	"github.com/tencent-connect/botgo/openapi"
)

var (
	Ctx context.Context
	Api openapi.OpenAPI
)

const (
	CmdConcurrentUsers = "/game"
)

func Init() {
	Api = botgo.NewOpenAPI(config.UToken).WithTimeout(3 * time.Second)
	Ctx = context.Background()
}

func MessageEventHandler(event *dto.WSPayload, data *dto.WSATMessageData) error {

	res := message.ParseCommand(data.Content) // 提取查询信息

	log.Printf("receive msg : %v", res)

	cmd := res.Cmd         // 暂只支持game查询
	content := res.Content // 游戏名
	switch cmd {
	case CmdConcurrentUsers: // steam当前在线人数

		cnt, err := steam.FindConcurrentUsersByName(content)
		var rsl string
		if err != nil {
			log.Printf("%v", err)
			rsl = fmt.Sprintf("查询出错：%v", err)
		} else {
			rsl = fmt.Sprintf("%v ：现在在线人数为 ：%v", content, cnt)
		}


		 Api.PostMessage(Ctx, data.ChannelID, &dto.MessageToCreate{
			MsgID:   data.ID, //如果未空则表示主动消息
			Content: rsl,

		})
		if err != nil {
			log.Printf("调用 PostMessage 接口失败, err = %v", err)
		}
	}
	return nil
}
