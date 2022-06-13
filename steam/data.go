/*
 * @Author: UerAx
 * @Date: 2022-06-13 20:45:09
 * @FilePath: \steam-concurrent-users\steam\data.go
 * Copyright (c) 2022 by UerAx uerax@live.com, All Rights Reserved.
 */
package steam

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var hash map[string]int64

const (
	url = "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
)

type steamRes struct {
	Applist Applist `json:"applist"`
}

type Applist struct {
	Apps []Apps `json:"apps"`
}

type Apps struct {
	Name string `json:"name"`
	Appid int64 `json:"appid"`
}

func init() {
	getGameNameCodeMap()
}

func getGameNameCodeMap() {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("steam接口异常 err: %v", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("steam接口异常 err: %v", err)
		return
	}

	defer resp.Body.Close()

	var tmp steamRes
	err = json.Unmarshal(body, &tmp)
	if err != nil {
		log.Printf("数据解析异常 err: %v", err)
		return
	}

	tmpMap := make(map[string]int64, len(tmp.Applist.Apps))

	for _, v := range tmp.Applist.Apps {
		// steam下架游戏存在空
		if v.Name == "" {
			continue
		}
		tmpMap[v.Name] = v.Appid
	}

	hash = tmpMap
}

func getCodeForName(name string) int64 {
	if code, ok := hash[name]; ok {
		return code
	}
	return -1
}
