package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	//	"time"

	pretiumvkgo "github.com/mallvielfrass/pretiumVKgo"
)

type Result struct {
	Response Response `json:"response"`
}

type Response struct {
	Count int    `json:"count"`
	Items []Item `json:"items"`
}
type GroupResult struct {
	Response GRoupResponse `json:"response"`
}

type GRoupResponse struct {
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
	//_________________________search________________________________________________
	fmt.Print("enter group`s name: ")
	bio := bufio.NewReader(os.Stdin)
	linex, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	line := string(linex)

	fmt.Print("enter count of group: ")
	offs := bufio.NewReader(os.Stdin)
	offsetx, _, err := offs.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	offset := string(offsetx)
	i, err := strconv.Atoi(offset)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	//"Typical weekday | Multifandom "
	r := api.GroupsSearch(line, 0, i)

	bx := []byte(r)
	var result Result

	err = json.Unmarshal(bx, &result)
	if err != nil {
		log.Fatalln(err)
	}

	R := result.Response.Items //_______________group`s list&info
	RV := 0
	count := len(R)
	fmt.Println("amount:", count)
	idLinks := make(map[int]int)
	for i := 0; i < count; i++ {
		id := R[RV].ID
		name := R[RV].Name
		idLinks[i] = R[RV].ID
		idz := "-" + strconv.Itoa(idLinks[i])
		get := api.WallGet(idz, 0, 1)
		wallbin := []byte(get)
		var wallres GroupResult
		err = json.Unmarshal(wallbin, &wallres)
		if err != nil {
			log.Fatalln(err)
		}
		f := wallres.Response.Count

		fmt.Println("Result number:", i, "|", "Name:", name, "|", "Id:", id, "|", "Number of post :", f)
		RV = RV + 1
	}

}
