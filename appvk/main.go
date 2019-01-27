package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	pretiumvkgo "github.com/mallvielfrass/pretiumVKgo"
)

func main() {
	// Считываем ключ из файла
	ready, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatalln("Не удалось считать ключ из файла:", err)
	}
	key := (strings.Split(string(ready), "\n"))[0]
	api := pretiumvkgo.NewAPI(key)
	//id := "29534144"
	id := "oldlentach"
	fields := "members_count"
	slice := api.Groups_getById(id, fields)
	fmt.Println(slice)

}
