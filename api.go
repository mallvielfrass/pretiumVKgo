package pretiumvkgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
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
	r := string(parametrs)
	//fmt.Println(r)
	url := "https://api.vk.com/method/" + cmeth + "?" + "&access_token=" + api.Key + "&v=5.58"
	bt := bytes.NewBuffer([]byte(r))
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

func (api *API) get(method string, params url.Values) string {
	// Добавляем токен и версию в список параметров, т.к. они обязательны
	params.Set("access_token", api.Key)
	params.Set("v", "5.58")

	// Формируем ссылку на сайт API. В Go ссылки имеют тип url.URL.
	apiURL := &url.URL{
		Scheme:   "https",            // протокол доступа
		Host:     "api.vk.com",       // адрес сайта
		Path:     "method/" + method, // путь на сайте
		RawQuery: params.Encode()}    // параметры запроса

	resp, err := http.Get(apiURL.String()) // выполняем запрос
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body) // считываем результат
	if err != nil {
		log.Fatalln(err)
	}

	return string(b) // PROFIT
}
func (api *API) Groups_getById(group_id string, fields string) string {
	fmt.Println("groups.getById")

	method := "groups.getById"
	type Request struct {
		GroupID string `url:"group_id"`
		Fields  string `url:"fields"`
	}
	rule := Request{
		GroupID: group_id,
		Fields:  fields}
	ruleValues, _ := query.Values(rule)
	retres := api.get(method, ruleValues)
	return retres
}
