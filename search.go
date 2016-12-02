package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SearchResult struct {
	Users []string 
	Weibos []string 
}

func search(key string) SearchResult {
	var result SearchResult
	URL := "http://127.0.0.1:6666/search?"
        value := url.Values{}
        value.Set("key", key)
        url := URL + value.Encode()
	res, err := http.Get(url)
	if err != nil {
           return result
	}
	detail, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
            return result
	}

	err = json.Unmarshal(detail, &result)
	if err != nil {
		fmt.Println("error:", err)
	}
        return result
}
func syncweibo(weiboid string) {
	URL := "http://127.0.0.1:6666/syncweibo?"
        value := url.Values{}
        value.Set("weiboid", weiboid)
        url := URL + value.Encode()
	res, err := http.Get(url)
	if err != nil {
	}
	res.Body.Close()
}
/*
func main(){
    fmt.Println(search("测试"))
    syncweibo("3230")
    fmt.Println(search("好厉害"))
}
*/
