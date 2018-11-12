package main

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"log"
)
func send(ids string,fields string) string{
	token:=tokens()
	method:="users.get"
	url := "https://api.vk.com/method/"+method+"?&access_token="+token+"&v=5.87"
	rule := `{"user_ids":`+ids+`,"fields"=`+fields+`}`
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(rule)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
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
