package main

import (
	"encoding/json"
	"log"
)

//"io/ioutil"

func main() {
	type GroupStruct struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		ScreenName   string `json:"screen_name"`
		IsClosed     int    `json:"is_closed"`
		Type         string `json:"type"`
		IsAdmin      int    `json:"is_admin"`
		IsMember     int    `json:"is_member"`
		IsAdvertiser int    `json:"is_advertiser"`
		Photo50      string `json:"photo_50"`
		Photo100     string `json:"photo_100"`
		Photo200     string `json:"photo_200"`
	}

	b := []byte(`{
	"id":147845620,
	"name":"VK Music",
	"screen_name":"vkmusic",
	"is_closed":0,
	"type":"page",
	"is_admin":0,
	"is_member":0,
	"is_advertiser":0,
	"photo_50":"https:\/\/pp.userapi.com\/c836138\/v836138505\/48277\/hs8Xd9zAYGo.jpg?ava=1",
	"photo_100":"https:\/\/pp.userapi.com\/c836138\/v836138505\/48276\/3yaDr9kg1Ac.jpg?ava=1",
	"photo_200":"https:\/\/pp.userapi.com\/c836138\/v836138505\/48274\/zpW_0KOAvvI.jpg?ava=1"}`)
	var result GroupStruct
	err, _ := json.Unmarshal(b, &result)
	log.Printf("%#v", result)
	//возвращает ошибку ./test.go:38:17: error: number of results does not match number of values
	//err, _ := json.Unmarshal(b, &result)
	//               ^
}
