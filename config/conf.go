/*
 * @Author: UerAx
 * @Date: 2022-06-13 17:21:27
 * @FilePath: \steam-concurrent-users\config\conf.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package config

import (
	"github.com/tencent-connect/botgo/token"
)

var UToken *token.Token

func Init() error {
	UToken = token.New(token.TypeBot)
	return UToken.LoadFromConfig("./etc/config.yaml")
}
