package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	//	"time"

	pretiumvkgo "github.com/mallvielfrass/pretiumVKgo"
)

type MResult struct {
	Response MResponse `json:"response"`
}

type MResponse struct {
	Count int     `json:"count"`
	Items []MItem `json:"items"`
}
type MItem struct {
	ID       int    `json:"id"`
	OwnedID  int    `json:"owner_id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	Date     int    `json:"date"`
	URL      string `json:"url"`
}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
func main() {

	// Считываем ключ из файла
	ready, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatalln("Не удалось считать ключ из файла:", err)
	}
	key := (strings.Split(string(ready), "\n"))[0]
	api := pretiumvkgo.NewAPI(key)
	fmt.Println("enter music name")
	offs := bufio.NewReader(os.Stdin)
	offsetx, _, err := offs.ReadLine()
	if err != nil {
		fmt.Println(err)
	}
	mnamez := string(offsetx)
	res := api.AudioSearch(mnamez, 1, 0)
	//fmt.Println(x)
	bx := []byte(res)
	var result MResult

	err = json.Unmarshal(bx, &result)
	if err != nil {
		log.Fatalln(err)
	}

	R := result.Response.Items

	fmt.Println("title:", R[0].Title, "\n", "url: ", R[0].URL)
	filename := "audio/" + R[0].Title + ".mp3"
	downloadFile(filename, R[0].URL)
	//	exec.Command("/bin/bash", "-c", "command -v foo")
	fmt.Println("down")
	command := "ffmpeg -i " + "'" + filename + "'" + " -c:a libopus -b:a 48k -vbr on -compression_level 10 -frame_duration 60 -application voip " + "'" + R[0].Title + "'" + ".opus"
	fmt.Println(command)
	out, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
		panic("some error found")
	}
	fmt.Println("cout")
	fmt.Println(string(out))
}
