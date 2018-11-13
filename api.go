package pretiumvkgo

import (
	"fmt"
	"net/http"
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
	fmt.Println("key =", api.Key)
}
