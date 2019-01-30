package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	pretiumvkgo "github.com/mallvielfrass/pretiumVKgo"
)

type Result struct {
	Response Response `json:"response"`
}

type Response struct {
	Count int    `json:"count"`
	Items []Item `json:"items"`
}
type Item struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ScreenName   string `json:"screen_name"`
	IsClosed     int    `json:"is_closed"`
	Type         string `json:"page"`
	IsAdmin      int    `json:"is_admin"`
	IsMember     int    `json:"is_member"`
	IsAdvertiser int    `json:"is_advertiser"`
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`
}

func main() {
	// Считываем ключ из файла
	ready, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatalln("Не удалось считать ключ из файла:", err)
	}
	key := (strings.Split(string(ready), "\n"))[0]
	api := pretiumvkgo.NewAPI(key)
	//id := "177388243"
	//id := "oldlentach"

	//S	fields := "members_count"
	//slice := api.Groups_getById(id, fields)
	//	fmt.Println(slice)
	r := api.GroupsSearch("typical weekday", 0, 10)
	//get := api.WallGet("-177388243", 0, 2)
	//fmt.Println(r)
	//	fmt.Println(get)
	//	idm := "-177388243"
	//comment := api.WallCreateComment(idm, 1, "hello go")
	//fmt.Println(comment)
	bx := []byte(r)
	var result Result

	err = json.Unmarshal(bx, &result)
	if err != nil {
		log.Fatalln(err)
	}

	R := result.Response.Items
	RV := 0
	count := result.Response.Count
	//fmt.Println(count)

	for i := 0; i < count; i++ {
		id := R[RV].ID
		name := R[RV].Name
		fmt.Println("Result number:", i, "|", "Name:", name, "|", "Id:", id)
		RV = RV + 1
	}

}
