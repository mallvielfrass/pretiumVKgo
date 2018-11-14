package pretiumvkgo

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"log"
)

type API struct {
	httpClient *http.Client

	Key string
}

func NewAPI(key string) *API {
	return &API{
		httpClient: &http.Client{},
		Key:        key}
}

func (api *API) Hello() {
	fmt.Println("Hello world!")
	//httpClient.Get(url)
}
func conventus (cmeth string,crule string) string{
	url := "https://api.vk.com/method/"cmeth+"?"+"&access_token="+api.Key+"&v=5.87"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(crule)))
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
	result:=string(bodyBuf)
	return result
}
func (api *API) Users_get(ids string, fields string) string{
	fmt.Println("Hello world!")
	//httpClient.Get(url)
	method:="users.get"
	rule := `{"user_ids":`+ids+`,"fields"=`+fields+`}`
	retres:= conventus(method,rule)
	return retres
}
