package pretiumvkgo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
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
func (api *API) conventus(cmeth string, crule string) string {
	url := "https://api.vk.com/method/" + cmeth + "?" + "&access_token=" + api.Key + "&v=5.58"
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

func (api *API) Users_get(ids string, fields string) string {
	fmt.Println("Users_get")
	//httpClient.Get(url)
	method := "users.get"
	rule := `{"user_ids":` + ids + `,"fields"=` + fields + `}`
	retres := api.conventus(method, rule)
	return retres
}

func (api *API) Groups_search(q string, offset int, count int) string {
	method := "groups.search"

	// Описываем структуру запроса согласно документации ВКонтакта.
	// Это нужно для того, чтобы не ошибиться при ручном формировании запросов.
	type Request struct {
		Q      string `url:"q"`                // текст поискового запроса.
		Offset int    `url:"offset,omitempty"` // смещение, необходимое для выборки определённого подмножества результатов поиска
		Count  int    `url:"count,omitempty"`  // количество результатов поиска, которое необходимо вернуть
	}

	// Создаём свой запрос с нашими данными
	rule := Request{
		Q:      q,
		Offset: offset,
		Count:  count}

	// Преобразуем структура запроса в формат url.Values. Этот тип содержит
	// параметры для типичных get-запросов, у которых параметры запроса указываются
	// в адресе
	ruleValues, _ := query.Values(rule)

	// Выполняем запрос
	retres := api.get(method, ruleValues)
	return retres
}

func (api *API) GetProfileInfo() string {

	method := "account.getProfileInfo"
	rule := ``
	retres := api.conventus(method, rule)
	return retres
}
