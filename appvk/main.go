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
	ID         int           `json:"id"`
	FromID     string        "from_id"
	OwnerID    string        "owner_id"
	Date       int           "date"
	Marked     int           "marked_as_ads"
	PostType   string        "post_type"
	Text       string        "text"
	CanEdit    bool          "can_edit"
	CreatedBy  int           "created_by"
	CanDelete  bool          "can_delete"
	CanPin     bool          "can_pin"
	PostSource []PostSourceX "post_source"
	Comments   []CommentsX   "comments"
	Likes      []LikesX      "likes"
}
type PostSourceX struct {
	Type string "type"
}
type CommentsX struct {
	Count         int    "count"
	CanPost       int    "can_post"
	GroupsCanPost string "groups_can_post"
	CanClose      bool   "can_close"
}
type LikesX struct {
	Count      int  "count"
	UserLikes  int  "user_likes"
	CanLike    bool "can_like"
	CanPublish bool "can_publish"
}

func main() {
	// Считываем ключ из файла
	ready, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatalln("Не удалось считать ключ из файла:", err)
	}
	key := (strings.Split(string(ready), "\n"))[0]
	api := pretiumvkgo.NewAPI(key)
	id := "177388243"
	//id := "oldlentach"
	fields := "members_count"
	slice := api.Groups_getById(id, fields)
	fmt.Println(slice)
	search := api.Groups_search("vielfrass", 0, 2)
	get := api.Wall_Get("-177388243", 0, 2)
	fmt.Println(search)
	fmt.Println(get)
	bx := []byte(get)
	var result Result
	err = json.Unmarshal(bx, &result)
	if err != nil {
		log.Fatalln(err)
	}

	R := result.Response.Items
	RV := 0
	id := R[RV].ID
	name := R[RV].Name
	fmt.Println(name, id)
}
