package main

import (
	"encoding/json"
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
	slice := api.Groups_search("Music", 0, 2)
	//Groups_search(q string, offset int, count int)
	bx := []byte(slice)
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
