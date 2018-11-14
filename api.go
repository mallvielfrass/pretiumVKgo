package pretiumvkgo

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
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
func (api *API) users_get(ids string, fields string) string{
	fmt.Println("Hello world!")
	//httpClient.Get(url)

	url := "https://api.vk.com/method/users.get?"+"&access_token="+api.Key+"&v=5.87"
	rule := `{"user_ids":`+ids+`,"fields"=`+fields+`}`
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(rule)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
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
