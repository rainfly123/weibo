package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"menteslibres.net/gosexy/redis"
	"net/http"
	//	"os"
	"os/exec"
	"path"
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

func Checkvideo(redis_key string, origin string) {

	var snapshot string
	snapshot = "http://7xvsyw.com1.z0.glb.clouddn.com/c.jpg"
	index := strings.LastIndex(origin, ".")
	if index > 0 {
		dest := origin[0:index+1] + "jpg"
		var args = []string{"-i", origin, "-vframes", "1", "-vf", "crop=iw:iw*9/16", "-f", "image2", "-y", dest}
		cmd := exec.Command("ffmpeg", args[0:]...)
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err == nil {
			snapshot = ACCESS_VIDEO_URL + path.Base(dest)
		}
	}
	videoaccess := ACCESS_VIDEO_URL + path.Base(origin)
	var client *redis.Client
	client, _ = clients.Get()
	client.HMSet(redis_key, "state", 2, "snapshot", snapshot, "type", "video", "url", videoaccess)
	client.Close()
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
	for i := range Channel {
		temp := strings.Split(i, "@")
		redis_key := temp[0]
		liveid := temp[1]
		if strings.Contains(liveid, "/") {
			go Checkvideo(redis_key, liveid)
		} else {
			go Checklive(redis_key, liveid)
		}
	}
}

var Channel chan string
