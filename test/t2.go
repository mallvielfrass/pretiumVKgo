package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Ответ VK API содержит один экземпляр "response"

func main() {
	b := []byte(`{"response":{"count":111339,"items":[{"id":147845620,"name":"VK Music","screen_name":"vkmusic","is_closed":0,"type":"page","is_admin":0,"is_member":0,"is_advertiser":0,"photo_50":"https:\/\/pp.userapi.com\/c836138\/v836138505\/48277\/hs8Xd9zAYGo.jpg?ava=1","photo_100":"https:\/\/pp.userapi.com\/c836138\/v836138505\/48276\/3yaDr9kg1Ac.jpg?ava=1","photo_200":"https:\/\/pp.userapi.com\/c836138\/v836138505\/48274\/zpW_0KOAvvI.jpg?ava=1"},{"id":339767,"name":"ТНТ MUSIC","screen_name":"tntmusic","is_closed":0,"type":"group","is_admin":0,"is_member":0,"is_advertiser":0,"photo_50":"https:\/\/sun9-11.userapi.com\/c846216\/v846216801\/164ba0\/CGwmZ4ycwJk.jpg?ava=1","photo_100":"https:\/\/sun9-7.userapi.com\/c846216\/v846216801\/164b9f\/ucFZdQdxgBk.jpg?ava=1","photo_200":"https:\/\/sun9-8.userapi.com\/c846216\/v846216801\/164b9e\/bMW8lu5tqDo.jpg?ava=1"}]}}`)
	var result Result
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Printf("%#v", result.Response.Items[1].ID)
	R := result.Response.Items
	x := R[1].ID
	fmt.Println(x)
}
