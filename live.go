package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"menteslibres.net/gosexy/redis"
	"net/http"
	"strings"
	"time"
)

type LIVE struct {
	Code             int
	Liveid           string
	Title            string
	Message          string
	Publish_url      string
	Rtmp_live_url    string
	State            int
	Playback_hls_url string
	Snapshot         string
	Persons          int
	Hls_live_url     string
	Tojson           string
	Supports         int
	Startime         string
}

func Checkvideo(redis_key string, filepath string) {

}

func Checklive(redis_key string, liveid string) {
	URL := "http://live.66boss.com/api/detail?liveid="
	var client *redis.Client
	client, _ = clients.Get()

ReCheck:
	res, err := http.Get(URL + liveid)
	if err != nil {
		client.Close()
		return
	}
	detail, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		client.Close()
		return
	}

	var live LIVE
	err = json.Unmarshal(detail, &live)
	if err != nil {
		fmt.Println("error:", err)
	}
	switch live.State {
	case 2:
		client.HMSet(redis_key, "state", 2, "snapshot", live.Snapshot, "type", "live", "url", live.Playback_hls_url)
		fmt.Println(live.State, live.Playback_hls_url)
		client.Close()
		fmt.Println("return")
		return
	case 1:
		client.HMSet(redis_key, "state", 1, "snapshot", live.Snapshot, "type", "live", "url", live.Rtmp_live_url)
		fmt.Println(live.State, live.Rtmp_live_url)
	}
	time.Sleep(10 * time.Second)
	goto ReCheck
}

func Check_thread() {
	for i := range channel {
		temp := strings.Split(i, "@")
		redis_key := temp[0]
		liveid := temp[1]
		if len(liveid) <= 1 {
			go Checkvideo(redis_key, filepath)
		} else {
			go Checklive(redis_key, liveid)
		}
	}
}

var Channel chan string
