package pretiumvkgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type API struct {
	httpClient *http.Client
	Key        string
}

func NewAPI(key string) *API {
	return &API{
		httpClient: &http.Client{},
		Key:        key}
}
func (api *API) conventus(cmeth string, a interface{}) string {
	parametrs, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return "conventus false"
	}
	fmt.Println(string(parametrs))
	url := "https://api.vk.com/method/" + cmeth + "?" + "&access_token=" + api.Key + "&v=5.58"
	bt := bytes.NewBuffer([]byte(string(parametrs)))
	req, err := http.NewRequest("POST", url, bt)
	req.Header.Set("Content-Type", "application/json")
	resp, err := api.httpClient.Do(req)
	if err != nil {
		log.Fatal("http request error:", err)
	}
	defer resp.Body.Close()
	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("response body read error:", err)
	}
	result := string(bodyBuf)
	return result
}

type GroupGetByID struct {
	ID     string `json:"group_id"`
	Fields string `json:"fields"`
}

func (api *API) Groups_getById(id string, fields string) string {
	fmt.Println("groups.getById")
	//httpClient.Get(url)
	method := "groups.getById"

	rule := GroupGetByID{
		ID:     id,
		Fields: fields}

	retres := api.conventus(method, rule)
	return retres
}
