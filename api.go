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

func (api *API) GroupsSearch(q string, offset int, count int) string {
	method := "groups.search"
	type Request struct {
		Q      string `url:"q"`                // текст поискового запроса.
		Offset int    `url:"offset,omitempty"` // смещение, необходимое для выборки определённого подмножества результатов поиска
		Count  int    `url:"count,omitempty"`  // количество результатов поиска, которое необходимо вернуть
	}
	rule := Request{
		Q:      q,
		Offset: offset,
		Count:  count}
	ruleValues, _ := query.Values(rule)
	retres := api.get(method, ruleValues)
	return retres
}
func (api *API) WallGet(owner_id string, offset int, count int) string {
	method := "wall.get"
	type Request struct {
		OwnerID string `url:"owner_id"`         // текст поискового запроса.
		Offset  int    `url:"offset,omitempty"` // смещение, необходимое для выборки определённого подмножества результатов поиска
		Count   int    `url:"count,omitempty"`  // количество результатов поиска, которое необходимо вернуть
	}
	rule := Request{
		OwnerID: owner_id,
		Offset:  offset,
		Count:   count}
	ruleValues, _ := query.Values(rule)
	retres := api.get(method, ruleValues)
	return retres
}

func (api *API) WallCreateComment(owner_id string, post_id int, message string) string {

	method := "wall.createComment"
	type CreateComment struct {
		OwnerID string `url:"owner_id"`
		PostID  int    `url:"post_id"`
		Message string `url:"message"`
	}
	rule := CreateComment{
		OwnerID: owner_id,
		PostID:  post_id,
		Message: message}

	ruleValues, _ := query.Values(rule)
	retres := api.get(method, ruleValues)
	return retres
}
func (api *API) AudioSearch(q string, count int, offset int) string {

	method := "audio.search"
	type CreateComment struct {
		Q      string `url:"q"`
		Count  int    `url:"count"`
		Offset int    `url:"offset"`
	}
	rule := CreateComment{
		Q:      q,
		Count:  count,
		Offset: offset}

	ruleValues, _ := query.Values(rule)
	retres := api.get(method, ruleValues)
	return retres
}
