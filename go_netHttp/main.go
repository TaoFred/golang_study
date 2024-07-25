package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://localhost:9280/main/home")
	if err != nil {
		fmt.Printf("get failed, err %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err: %v\n", err)
		return
	}
	fmt.Println(string(body))
}
