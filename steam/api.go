package steam

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	concurrentUserUrl = "https://api.steampowered.com/ISteamUserStats/GetNumberOfCurrentPlayers/v1?appid=%d"
)

type concurrentUserRes struct {
	Data concurrentUser `json:"response"`
}

type concurrentUser struct {
	PlayerCnt int64 `json:"player_count"`
	result    int64 `json:"result""`
}

func FindConcurrentUsersByName(name string) (string, error) {
	code := getCodeForName(name)
	if -1 == code {
		return "0", errors.New("没找到对应的游戏")
	}
	url := fmt.Sprintf(concurrentUserUrl, code)
	data, err := http.Get(url)
	if err != nil {
		log.Printf("steam接口异常 : %v", err)
		return "0", errors.New("steam接口异常")
	}

	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Printf("steam接口异常 : %v", err)
		return "0", errors.New("steam接口异常")
	}

	defer data.Body.Close()

	var tmp concurrentUserRes

	err = json.Unmarshal(body, &tmp)

	if err != nil {
		log.Printf("数据解析异常 : %v", err)
		return "0", errors.New("数据解析异常")
	}

	return fmt.Sprintf("%d", tmp.Data.PlayerCnt), nil

}
