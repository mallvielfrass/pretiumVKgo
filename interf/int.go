package main

import (
	"encoding/json"
	"fmt"
)

func inter(g interface{}) {
	x, err := json.Marshal(g)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(x))

}
func dx() {

	type Group struct {
		ID   string `json:"user_ids"`
		Name string `json:"name"`
	}
	group := Group{
		ID:   "233234567",
		Name: "vk"}

	inter(group)
}
func main() {
	dx()

}
