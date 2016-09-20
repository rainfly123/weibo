package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"menteslibres.net/gosexy/redis"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	UPLOAD_PATH string = "./upload/"
	ACCESS_URL  string = "http://192.168.1.251:8888/"
)

func compress(name string) {
	var quality int = 90
	fmt.Println(name)
	if !strings.HasSuffix(name, ".jpg") {
		fmt.Println("not jpg")
		return
	}
	fileinfo, _ := os.Stat(name)
	if fileinfo.Size()/1024 < 1024 {
		return
	}

	f1, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	m1, err := jpeg.Decode(f1)
	if err != nil {
		panic(err)
	}

	f1.Close()
	d := path.Dir(name)
	var temp string = path.Join(d, "_temp_")
RECOMPRESS:
	fmt.Println(name)
	quality -= 10
	f2, err := os.Create(temp)
	if err != nil {
		panic(err)
	}

	bounds := m1.Bounds()
	m := image.NewRGBA(bounds)
	draw.Draw(m, bounds, m1, bounds.Min, draw.Src)
	err = jpeg.Encode(f2, m, &jpeg.Options{quality})
	if err != nil {
		panic(err)
	}
	f2.Close()
	fileinfo, _ = os.Stat(temp)
	if fileinfo.Size()/1024 < 1024 {
		os.Rename(temp, name)
		return
	}
	goto RECOMPRESS

}

func getUUID() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
func getFileName(name string) string {

	var temp string = "error"
	i := strings.LastIndex(name, ".")
	if i > 0 {
		uuid := getUUID()
		temp = uuid + name[i:]
	}
	return temp

}

func uploadHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		io.WriteString(w, "<html><head><title>我的第一个页面</title></head><body><form action='upload?user=2' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file'  /><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		temp := getFileName(head.Filename)
		uuidFile := UPLOAD_PATH + temp
		fW, err := os.Create(uuidFile)
		if err != nil {
			fmt.Println("create file error")
			return
		}
		//defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("copy file error")
			return
		}
		fW.Close()
		compress(uuidFile)
		io.WriteString(w, (ACCESS_URL + temp))
		io.WriteString(w, r.FormValue("user"))
	}
}

func portraitHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		login_user := r.FormValue("login_user")
		io.WriteString(w, fmt.Sprintf("<html><head><title>我的第一个页面</title></head><body><form action='portrait?login_user=%s' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file'  /><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>", login_user))
	} else {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		temp := getFileName(head.Filename)
		uuidFile := UPLOAD_PATH + temp
		fW, err := os.Create(uuidFile)
		if err != nil {
			jsonres := JsonResponse{2, "system error"}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			jsonres := JsonResponse{2, "system error"}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		var ok bool
		var client *redis.Client
		client, ok = clients.Get()
		if ok != true {
			jsonres := JsonResponse{2, "system error"}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		login_user := r.FormValue("login_user")
		if len(login_user) < 1 {
			jsonres := JsonResponse{1, "argument error"}
			b, _ := json.Marshal(jsonres)
			io.WriteString(w, string(b))
			return
		}
		key := "user_" + login_user + "_profile"
		client.HMSet(key, "portrait", (ACCESS_URL + temp))
		client.Close()

		type MyResponse struct {
			JsonResponse
			Url string `json:"url"`
		}

		jsonres := MyResponse{}
		jsonres.Code = 0
		jsonres.Message = "Succeeded"
		jsonres.Url = (ACCESS_URL + temp)
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))

	}
}

func testHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		io.WriteString(w, "<html><head></head><body><form action='test?user=2' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"test\" name='file'  /><br/><label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		io.WriteString(w, r.FormValue("user"))
		io.WriteString(w, r.FormValue("file"))
	}
}
