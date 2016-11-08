package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"menteslibres.net/gosexy/redis"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type VideoType struct {
	State    int    `json:"state"`
	Snapshot string `json:"snapshot"`
	Type     string `json:"type"`
	Url      string `json:"url"`
}
type WeiBo struct {
	Weiboid  int       `json:"weiboid"`
	Type     string    `json:"type"`
	Concern  bool      `json:"concerned"`
	Support  bool      `json:"supported"`
	Msg      string    `json:"msg"`
	Author   string    `json:"author"`
	Creatime string    `json:"creatime"`
	Supports int       `json:"supports"`
	Resent   int       `json:"resent"`
	Pictures []string  `json:"pictures"`
	Comments int       `json:"comments"`
	Origin   *WeiBo    `json:"origin"`
	Flag     string    `json:"class"`
	Video    VideoType `json:"video"`
	Userinfo User      `json:"user"`
	Redevpid string    `json:"redpacketid"`
}

type User struct {
	Userid    string `json:"userid"`
	Concern   bool   `json:"concerned"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	Location  string `json:"location"`
	Signature string `json:"signature"`
	Portrait  string `json:"portrait"`
	Follower  []User `json:"follower"`
	Following []User `json:"following"`
	Recommend []User `json:"recommend"`
}

type ALL_WeiBO []WeiBo

func (list ALL_WeiBO) Len() int {
	return len(list)
}
func (list ALL_WeiBO) Less(i, j int) bool {
	return list[i].Creatime > list[j].Creatime
}
func (list ALL_WeiBO) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

var logger *log.Logger
var host = "192.168.1.251"
var port = uint(6379)
var clients redisPool

func writeHandle(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("author")
	if len(author) < 1 {
	}
	msg := req.FormValue("msg")
	if len(msg) < 3 {
	}
	pic := req.FormValue("pic")
	if len(pic) < 3 {
	}

	var ok bool
	var client *redis.Client
	client, ok = clients.Get()
	if ok != true {
		io.WriteString(w, "error!\n")
	}

	strID, _ := client.Get("globalID")
	key := "weibo_" + strID
	now := time.Now().Format("2006-01-02 15:04:05")
	client.HMSet(key, "weiboid", strID, "msg", msg, "author", author, "creatime", now, "supports", 0, "resent", 0,
		"pictures", pic, "comments", 0)
	client.LPush("weibo_message", strID)
	user := "user_" + author + "_weibo"
	client.LPush(user, strID)
	client.Incr("globalID")

	fmt.Fprintf(w, "%s", key)
	io.WriteString(w, "ok!\n")

	client.Close()
}
func receiveFile(w http.ResponseWriter, req *http.Request, name string) string {
	file, head, err := req.FormFile(name)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	temp := getFileName(head.Filename)
	uuidFile := UPLOAD_PATH + temp
	fW, err := os.Create(uuidFile)
	if err != nil {
		fmt.Println("create file error")
		return ""
	}
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("copy file error")
		return ""
	}
	fW.Close()
	//compress(uuidFile)
	Resize(uuidFile)
	return (ACCESS_URL + temp)
}

func writev2Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		io.WriteString(w, fmt.Sprintf("<html><head><title>我的第一个页面</title></head><body><form action=\"writev2?author=%s&msg=%s\" method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file0'/><br/><input type=\"file\" name='file1'/><br/><input type=\"file\" name='file2'/><br/><<input type=\"file\" name='file3'/><br/><<input type=\"file\" name='file4'/><br/><<input type=\"file\" name='file5'/><br/><<input type=\"file\" name='file6'/><br/><<input type=\"file\" name='file7'/><br/><<input type=\"file\" name='file8'/><br/><<label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>", req.FormValue("author"), req.FormValue("msg")))
	} else {
		var pictures []string
		for i := 0; i < 9; i++ {
			var filenametemp string
			filenametemp = receiveFile(w, req, fmt.Sprintf("file%d", i))
			if len(filenametemp) <= 1 {
				break
			}
			pictures = append(pictures, filenametemp)
		}

		pic := strings.Join(pictures, ",")

		author := req.FormValue("author")
		msg := req.FormValue("msg")
		if len(author) < 1 || len(msg) < 3 {
			jsonres := JsonResponse{1, "argument error"}
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

		strID, _ := client.Get("globalID")
		key := "weibo_" + strID
		now := time.Now().Format("2006-01-02 15:04:05")
		client.HMSet(key, "weiboid", strID, "msg", msg, "author", author, "creatime", now, "supports", 0, "resent", 0,
			"pictures", pic, "comments", 0)
		client.LPush("weibo_message", strID)
		user := "user_" + author + "_weibo"
		client.LPush(user, strID)
		client.Incr("globalID")

		type MyResponse struct {
			JsonResponse
			Weiboid  string `json:"weiboid"`
			Pictures string `json:"pictures"`
		}
		jsonres := MyResponse{}
		jsonres.Code = 0
		jsonres.Message = "Succeeded"
		jsonres.Weiboid = strID
		jsonres.Pictures = pic
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))

		client.Close()
	}
}

func writev3Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		io.WriteString(w, fmt.Sprintf("<html><head><title>我的第一个页面</title></head><body><form action=\"writev3?author=%s&liveid=%s&msg=%s\" method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file'/><br/><label><input type=\"submit\" value=\"上传视频\"/></label></form></body></html>", req.FormValue("author"), req.FormValue("liveid"), req.FormValue("msg")))
	} else {
		vtype := "live"
		var access string
		author := req.FormValue("author")
		msg := req.FormValue("msg")
		if len(author) < 1 || len(msg) < 3 {
			jsonres := JsonResponse{1, "argument error"}
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

		liveid := req.FormValue("liveid")
		if len(liveid) < 1 {
			vtype = "video"
			file, head, err := req.FormFile("file")
			if err != nil {
				jsonres := JsonResponse{1, "argument error"}
				b, _ := json.Marshal(jsonres)
				io.WriteString(w, string(b))
				return
			}
			defer file.Close()

			temp := getFileName(head.Filename)
			uuidFile := UPLOAD_VIDEO_PATH + temp
			liveid = uuidFile
			fW, err := os.Create(uuidFile)
			if err != nil {
				jsonres := JsonResponse{2, "system error"}
				b, _ := json.Marshal(jsonres)
				io.WriteString(w, string(b))
				return
			}
			_, err = io.Copy(fW, file)
			if err != nil {
				jsonres := JsonResponse{2, "system error"}
				b, _ := json.Marshal(jsonres)
				io.WriteString(w, string(b))
				return
			}
			fW.Close()
			access = ACCESS_VIDEO_URL + temp
		}

		strID, _ := client.Get("globalID")
		key := "weibo_" + strID
		key_video := key + "_video"
		now := time.Now().Format("2006-01-02 15:04:05")
		client.HMSet(key, "weiboid", strID, "msg", msg, "author", author, "creatime", now, "supports", 0, "resent", 0, "video", key_video, "comments", 0)
		client.LPush("weibo_message", strID)
		user := "user_" + author + "_weibo"
		client.LPush(user, strID)
		client.Incr("globalID")

		snapshot := "http://7xvsyw.com1.z0.glb.clouddn.com/b.jpg"
		client.HMSet(key_video, "state", 1, "snapshot", snapshot, "type", vtype, "url", "abcdefg")
		Channel <- key_video + "@" + liveid
		if strings.Contains(liveid, "/") {
			liveid = access
		}

		type MyResponse struct {
			JsonResponse
			Weiboid string `json:"weiboid"`
			Video   string `json:"Video"`
		}
		jsonres := MyResponse{}
		jsonres.Code = 0
		jsonres.Message = "Succeeded"
		jsonres.Weiboid = strID
		jsonres.Video = liveid
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))

		client.Close()
	}
}

func writev4Handle(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("author")
	msg := req.FormValue("msg")
	redpacketid := req.FormValue("redpacketid")
	if len(author) < 1 || len(msg) < 3 || len(redpacketid) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	strID, _ := client.Get("globalID")
	key := "weibo_" + strID
	now := time.Now().Format("2006-01-02 15:04:05")
	client.HMSet(key, "weiboid", strID, "msg", msg, "author", author, "creatime", now, "supports", 0, "resent", 0, "redpacketid", redpacketid, "comments", 0, "flag", "红包")
	client.LPush("weibo_message", strID)
	user := "user_" + author + "_weibo"
	client.LPush(user, strID)
	client.Incr("globalID")

	type MyResponse struct {
		JsonResponse
		Weiboid  string `json:"weiboid"`
		Redevpid string `json:"redpacketid"`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Weiboid = strID
	jsonres.Redevpid = redpacketid
	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
	client.Close()
}

func commentHandle(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("login_user")
	comment := req.FormValue("comment")
	weiboid := req.FormValue("weiboid")

	if len(author) < 1 || len(comment) < 3 || len(weiboid) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	now := time.Now().Format("2006-01-02 15:04:05")
	keyc := "weibo_" + weiboid + "_comments"
	idx, _ := client.LLen(keyc)
	value := author + "$" + comment + "$" + now + "$0"
	client.RPush(keyc, value)
	keyv := "weibo_" + weiboid
	client.HIncrBy(keyv, "comments", 1)
	tempp := strconv.Itoa(int(idx))
	keycomsup := "weibo_" + weiboid + "_comment_" + tempp
	client.SAdd(keycomsup, "0")
	client.Close()
	jsonres := JsonResponse{0, "Succeeded"}
	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}
func supportcommentHandle(w http.ResponseWriter, req *http.Request) {
	weiboid := req.FormValue("weiboid")
	login_user := req.FormValue("login_user")
	commentid := req.FormValue("commentid")
	if len(weiboid) < 1 || len(commentid) < 1 || len(login_user) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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
	key := "weibo_" + weiboid + "_comments"
	id, _ := strconv.Atoi(commentid)
	kid := int64(id)
	comment, _ := client.LRange(key, kid, kid)
	temp := strings.Split(comment[0], "$")
	author := temp[0]
	con := temp[1]
	creatime := temp[2]
	supports, _ := strconv.Atoi(temp[3])
	supports += 1
	value := author + "$" + con + "$" + creatime + "$" + strconv.Itoa(supports)
	client.LSet(key, kid, value)

	jsonres := JsonResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))

	keycomsup := "weibo_" + weiboid + "_comment_" + commentid
	client.SAdd(keycomsup, login_user)

	client.Close()

}

func checkcommentHandle(w http.ResponseWriter, req *http.Request) {

	type Comment struct {
		Id        int    `json:"commentid"`
		Comment   string `json:"comment"`
		Creatime  string `json:"creatime"`
		Supports  int    `json:"supports"`
		Supported bool   `json:"supported"`
		Author    User   `json:"author"`
	}

	weiboid := req.FormValue("weiboid")
	login_user := req.FormValue("login_user")
	if len(weiboid) < 1 || len(login_user) < 1 {
		jsonres := JsonResponse{1, "argument error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	all := make([]Comment, 0, 1000)
	var ok bool
	var client *redis.Client
	client, ok = clients.Get()
	if ok != true {
		jsonres := JsonResponse{2, "system error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	key := "weibo_" + weiboid + "_comments"
	comments, _ := client.LRange(key, 0, -1)
	for ix, v := range comments {
		var comment Comment
		comment.Id = ix
		temp := strings.Split(v, "$")
		comment.Author = getUserinfo(temp[0], client, false)
		comment.Comment = temp[1]
		comment.Creatime = temp[2]
		comment.Supports, _ = strconv.Atoi(temp[3])
		index := strconv.Itoa(ix)
		keycomsup := "weibo_" + weiboid + "_comment_" + index
		comment.Supported, _ = client.SIsMember(keycomsup, login_user)
		all = append(all, comment)
	}

	type MyResponse struct {
		JsonResponse
		Data []Comment `json:"data"`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Data = all

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
	client.Close()
}

func supportHandle(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("login_user")
	weiboid := req.FormValue("weiboid")
	if len(author) < 1 || len(weiboid) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	keys := "weibo_" + weiboid + "_supports"
	client.RPush(keys, author)

	keyv := "weibo_" + weiboid
	client.HIncrBy(keyv, "supports", 1)
	client.Close()

	jsonres := JsonResponse{0, "Succeeded"}
	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}
func unsupportHandle(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("login_user")
	weiboid := req.FormValue("weiboid")
	if len(author) < 1 || len(weiboid) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	keys := "weibo_" + weiboid + "_supports"
	client.LRem(keys, 0, author)

	keyv := "weibo_" + weiboid
	client.HIncrBy(keyv, "supports", -1)
	client.Close()

	jsonres := JsonResponse{0, "Succeeded"}
	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}

func checksupportHandle(w http.ResponseWriter, req *http.Request) {

	weiboid := req.FormValue("weiboid")
	if len(weiboid) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	key := "weibo_" + weiboid + "_supports"
	supports, _ := client.LRange(key, 0, -1)
	ALL_USERS := make([]User, 0)
	for _, v := range supports {
		user := getUserinfo(v, client, false)
		ALL_USERS = append(ALL_USERS, user)
	}

	type MyResponse struct {
		JsonResponse
		Data []User `json:"data"`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Data = ALL_USERS

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
	client.Close()
}

func concernHandle(w http.ResponseWriter, req *http.Request) {
	login_user := req.FormValue("login_user")
	concern := req.FormValue("concern")
	if len(login_user) < 1 || len(concern) < 1 || strings.EqualFold(login_user, concern) {
		jsonres := JsonResponse{1, "argument error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	var ok bool
	var client *redis.Client
	client, ok = clients.Get()
	if ok != true {
		jsonres := JsonResponse{2, "sytem error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	key := "user_" + login_user + "_following"
	client.SAdd(key, concern)

	key = "user_" + concern + "_follower"
	client.SAdd(key, login_user)
	client.Close()

	jsonres := JsonResponse{0, "Succeeded"}
	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}

func cancelconcernHandle(w http.ResponseWriter, req *http.Request) {
	login_user := req.FormValue("login_user")
	cancel := req.FormValue("cancel")
	if len(login_user) < 1 || len(cancel) < 1 {
		jsonres := JsonResponse{1, "argument error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	var ok bool
	var result bool
	var client *redis.Client
	client, ok = clients.Get()
	if ok != true {
		jsonres := JsonResponse{2, "sytem error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	key := "user_" + login_user + "_following"
	is, _ := client.SMove(key, "__trash__", cancel)
	result = is

	key = "user_" + cancel + "_follower"
	im, _ := client.SMove(key, "__trash__", login_user)
	result = result && im
	client.Close()
	if result != true {
		jsonres := JsonResponse{2, "system error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}
	jsonres := JsonResponse{0, "Succeeded"}
	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}

func getWeibo(weiboid string, client *redis.Client) *WeiBo {
	if len(weiboid) <= 1 {
		return nil
	}
	//	ls, err := client.HGetAll("weibo_" + weiboid)
	ls, err := client.HMGet("weibo_"+weiboid, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin", "flag", "video", "redpacketid")
	if err != nil {
		return nil
	}
	var weibo WeiBo
	for i, v := range ls {
		switch i {
		case 0:
			weibo.Weiboid, _ = strconv.Atoi(v)
		case 1:
			weibo.Msg = v
		case 2:
			weibo.Author = v
		case 3:
			weibo.Creatime = v
		case 4:
			weibo.Supports, _ = strconv.Atoi(v)
		case 5:
			weibo.Resent, _ = strconv.Atoi(v)
		case 6:
			if len(v) >= 3 {
				temp := strings.Split(v, ",")
				weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
			}
		case 7:
			weibo.Comments, _ = strconv.Atoi(v)
		case 9:
			weibo.Flag = v
		case 10:
			if len(v) >= 1 {
				weibo.Video = getVideoinfo(v, client)
				weibo.Type = "video"
			} else {
				weibo.Type = "text"
			}

		case 11:
			if len(v) >= 1 {
				weibo.Redevpid = v
				weibo.Type = "redpacket"
			}
		}
	}
	weibo.Userinfo = getUserinfo(weibo.Author, client, false)
	return &weibo
}

func deleteHandle(w http.ResponseWriter, req *http.Request) {
	login_user := req.FormValue("login_user")
	weiboid := req.FormValue("weiboid")
	if len(login_user) < 1 || len(weiboid) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	key := "weibo_message"
	client.LRem(key, 0, weiboid)
	wkey := "weibo_" + weiboid
	if login_user == "1" {
		ls, err := client.HMGet(wkey, "author")
		if err == nil {
			author := ls[0]
			key = "user_" + author + "_weibo"
			client.LRem(key, 0, weiboid)
		}
	} else {
		key = "user_" + login_user + "_weibo"
		client.LRem(key, 0, weiboid)
	}
	client.Del(wkey)

	jsonres := JsonResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
	client.Close()
}

func checkHandle(w http.ResponseWriter, req *http.Request) {
	login_user := req.FormValue("login_user")
	if len(login_user) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	all := make([]string, 0, 200)
	key := "user_" + login_user + "_following"
	users, _ := client.SMembers(key)
	users = append(users, login_user)

	for _, v := range users {
		k := "user_" + v + "_weibo"
		weibos, _ := client.LRange(k, 0, -1)
		if len(weibos) >= 1 {
			all = append(all, weibos[:len(weibos)]...)
		}
	}

	//fmt.Println("weibo IDS ", all) //all is weiboid list
	allweibo := make(ALL_WeiBO, 0, 5000)
	for _, vv := range all {
		//ls, err := client.HGetAll("weibo_" + vv)
		ls, err := client.HMGet("weibo_"+vv, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin", "flag", "video", "redpacketid")
		if err != nil {
			continue
		}
		var weibo WeiBo
		for i, v := range ls {
			switch i {
			case 0:
				weibo.Weiboid, _ = strconv.Atoi(v)
				weibo.Support = alreadSupport(weibo.Weiboid, login_user, client)
			case 1:
				weibo.Msg = v
			case 2:
				weibo.Author = v
				weibo.Concern = true
			case 3:
				weibo.Creatime = v
			case 4:
				weibo.Supports, _ = strconv.Atoi(v)
			case 5:
				weibo.Resent, _ = strconv.Atoi(v)
			case 6:
				if len(v) >= 3 {
					temp := strings.Split(v, ",")
					weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
				}
			case 7:
				weibo.Comments, _ = strconv.Atoi(v)
			case 8:
				weibo.Origin = getWeibo(v, client)
			case 9:
				weibo.Flag = v
			case 10:
				if len(v) >= 1 {
					weibo.Video = getVideoinfo(v, client)
					weibo.Type = "video"
				} else {
					weibo.Type = "text"
				}

			case 11:
				if len(v) >= 1 {
					weibo.Redevpid = v
					weibo.Type = "redpacket"
				}
			}
		}
		if weibo.Type == "video" {
			if strings.Contains(weibo.Video.Url, "abcdefg") || len(weibo.Video.Url) < 5 {
				continue
			}
		}
		weibo.Userinfo = getUserinfo(weibo.Author, client, false)
		if weibo.Origin != nil {
			weibo.Type = weibo.Origin.Type
		}
		allweibo = append(allweibo, weibo)
	}
	sort.Sort(allweibo)

	type MyResponse struct {
		JsonResponse
		Data ALL_WeiBO `json:"data"`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Data = allweibo

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
	client.Close()
}

func checkmyHandle(w http.ResponseWriter, req *http.Request) {
	login_user := req.FormValue("login_user")
	if len(login_user) < 1 {
		jsonres := JsonResponse{1, "argument error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}
	key := "user_" + login_user + "_weibo"

	var ok bool
	var client *redis.Client
	client, ok = clients.Get()
	if ok != true {
		jsonres := JsonResponse{2, "sytem error"}
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	weibos, _ := client.LRange(key, 0, -1)

	allweibo := make(ALL_WeiBO, 0, 500)
	for _, vv := range weibos {
		//ls, err := client.HGetAll("weibo_" + vv)
		ls, err := client.HMGet("weibo_"+vv, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin", "flag", "video", "redpacketid")
		if err != nil {
			continue
		}
		var weibo WeiBo
		for i, v := range ls {
			switch i {
			case 0:
				weibo.Weiboid, _ = strconv.Atoi(v)
			case 1:
				weibo.Msg = v
			case 2:
				weibo.Author = v
			case 3:
				weibo.Creatime = v

			case 4:
				weibo.Supports, _ = strconv.Atoi(v)
			case 5:
				weibo.Resent, _ = strconv.Atoi(v)
			case 6:
				if len(v) >= 3 {
					temp := strings.Split(v, ",")
					weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
				}
			case 7:
				weibo.Comments, _ = strconv.Atoi(v)
			case 8:
				weibo.Origin = getWeibo(v, client)
			case 9:
				weibo.Flag = v
			case 10:
				if len(v) >= 1 {
					weibo.Video = getVideoinfo(v, client)
					weibo.Type = "video"
				} else {
					weibo.Type = "text"
				}
			case 11:
				if len(v) >= 1 {
					weibo.Redevpid = v
					weibo.Type = "redpacket"
				}
			}
		}

		weibo.Userinfo = getUserinfo(weibo.Author, client, false)
		if weibo.Origin != nil {
			weibo.Type = weibo.Origin.Type
		}
		allweibo = append(allweibo, weibo)
	}
	//sort.Sort(allweibo)

	type MyResponse struct {
		JsonResponse
		Data ALL_WeiBO `json:"data"`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Data = allweibo
	b, _ := json.Marshal(jsonres)

	io.WriteString(w, string(b))
	client.Close()
}

func profileHandle(w http.ResponseWriter, req *http.Request) {
	login_user := req.FormValue("login_user")
	nickname := req.FormValue("nickname")
	gender := req.FormValue("gender")
	location := req.FormValue("location")
	signature := req.FormValue("signature")
	jsonres := JsonResponse{1, "argument error"}

	if len(login_user) < 1 {
		//if len(login_user) < 1 || len(nickname) < 1 || len(gender) < 1 || len(location) < 1 || len(signature) < 1 {
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	var ok bool
	var client *redis.Client
	client, ok = clients.Get()
	if ok != true {
		b, _ := json.Marshal(jsonres)
		io.WriteString(w, string(b))
		return
	}

	//	portrait := "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg"
	key := "user_" + login_user + "_profile"
	client.HMSet(key, "nickname", nickname, "gender", gender, "location", location, "signature", signature) //, "portrait", portrait)

	b, _ := json.Marshal(JsonResponse{0, "Succeeded"})
	io.WriteString(w, string(b))
	client.Close()
}

func recommend(client *redis.Client) []string {

	var users []string
	weibos, _ := client.LRange("weibo_message", 0, 100)
	for _, v := range weibos {
		key := "weibo_" + v
		ls, err := client.HGet(key, "author")
		if err != nil || len(ls) < 1 {
			continue
		}
		users = append(users, ls)
	}
	return users
}

func getVideoinfo(key string, client *redis.Client) VideoType {
	var video VideoType
	ls, err := client.HMGet(key, "state", "snapshot", "type", "url")
	if err != nil {
		return video
	}
	for i, v := range ls {
		switch i {
		case 0:
			video.State, _ = strconv.Atoi(v)
		case 1:
			video.Snapshot = v
		case 2:
			video.Type = v
		case 3:
			video.Url = v
		}
	}
	return video
}
func getUserinfo(userid string, client *redis.Client, detail bool) User {

	var user User
	key := "user_" + userid + "_profile"
	ls, err := client.HMGet(key, "nickname", "gender", "location", "signature", "portrait")
	if err != nil {
		return user
	}
	user.Userid = userid
	for i, v := range ls {
		switch i {
		case 0:
			user.Nickname = v
			if v == "" {
				end := len(userid)
				start := end - 4
				if start < 0 {
					start = 0
				}
				user.Nickname = "游客_" + userid[start:end]
			}
		case 1:
			user.Gender = v
		case 2:
			user.Location = v
		case 3:
			user.Signature = v
		case 4:
			if v == "" {
				v = "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg"
			}
			user.Portrait = v
		}
	}
	if detail {
		var temp User
		keyer := "user_" + userid + "_follower"
		follower, _ := client.SMembers(keyer)
		for _, n := range follower {
			temp = getUserinfo(n, client, false)
			user.Follower = append(user.Follower, temp)
		}

		keying := "user_" + userid + "_following"
		following, _ := client.SMembers(keying)
		for _, k := range following {
			temp = getUserinfo(k, client, false)
			user.Following = append(user.Following, temp)
		}

		// s, _ := client.SRandMember("all_users", 10)
		var recuser []string
		s := recommend(client)
		for _, us := range s {
			if len(recuser) >= 5 {
				break
			}
			if has(following, us) || strings.EqualFold(us, userid) {
				continue
			} else {
				if !has(recuser, us) {
					recuser = append(recuser, us)
				}
			}
		}

		for _, p := range recuser {
			temp = getUserinfo(p, client, false)
			user.Recommend = append(user.Recommend, temp)
		}
	}
	return user
}

func whetherConcerned(login_user, userid string, client *redis.Client) bool {
	key := "user_" + login_user + "_following"
	following, _ := client.SMembers(key)

	total := len(following)
	sort.Strings(following)

	has := sort.SearchStrings(following, userid)
	if has < total {
		if strings.EqualFold(userid, following[has]) {
			return true
		}
	}
	return false
}

func getUserinfoLoginuser(login_user, userid string, client *redis.Client, detail bool) User {

	var user User
	key := "user_" + userid + "_profile"
	ls, err := client.HMGet(key, "nickname", "gender", "location", "signature", "portrait")
	if err != nil {
		return user
	}
	if strings.EqualFold(userid, login_user) {
		user.Concern = true
	} else {
		user.Concern = whetherConcerned(login_user, userid, client)
	}
	user.Userid = userid
	for i, v := range ls {
		switch i {
		case 0:
			user.Nickname = v
			if v == "" {
				end := len(userid)
				start := end - 4
				if start < 0 {
					start = 0
				}
				user.Nickname = "游客_" + userid[start:end]
			}
		case 1:
			user.Gender = v
		case 2:
			user.Location = v
		case 3:
			user.Signature = v
		case 4:
			if v == "" {
				v = "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg"
			}
			user.Portrait = v
		}
	}
	if detail {
		var temp User
		keyer := "user_" + userid + "_follower"
		follower, _ := client.SMembers(keyer)
		for _, n := range follower {
			temp = getUserinfoLoginuser(login_user, n, client, false)
			user.Follower = append(user.Follower, temp)
		}

		keying := "user_" + userid + "_following"
		following, _ := client.SMembers(keying)
		for _, k := range following {
			temp = getUserinfoLoginuser(login_user, k, client, false)
			user.Following = append(user.Following, temp)
		}

		// s, _ := client.SRandMember("all_users", 10)
		var recuser []string
		s := recommend(client)
		for _, us := range s {
			if len(recuser) >= 5 {
				break
			}
			if has(following, us) || strings.EqualFold(us, userid) {
				continue
			} else {
				if !has(recuser, us) {
					recuser = append(recuser, us)
				}
			}
		}

		for _, p := range recuser {
			temp = getUserinfoLoginuser(login_user, p, client, false)
			user.Recommend = append(user.Recommend, temp)
		}
	}
	return user
}

func has(sl []string, el string) bool {
	for _, i := range sl {
		if strings.EqualFold(i, el) {
			return true
		}
	}
	return false
}

func userInfo(w http.ResponseWriter, req *http.Request) {
	userid := req.FormValue("userid")
	login_user := req.FormValue("login_user")

	if len(userid) < 1 || len(login_user) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	type MyResponse struct {
		JsonResponse
		Data User `json:"data"`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Data = getUserinfoLoginuser(login_user, userid, client, true)
	client.Close()

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}

func forwardHandle(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("login_user")
	msg := req.FormValue("msg")
	origin := req.FormValue("origin")
	if len(author) < 1 || len(msg) < 3 || len(origin) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	strID, _ := client.Get("globalID")
	key := "weibo_" + strID
	now := time.Now().Format("2006-01-02 15:04:05")
	client.HMSet(key, "weiboid", strID, "msg", msg, "author", author, "creatime", now, "supports", 0, "resent", 0,
		"pictures", "", "comments", 0, "origin", origin)
	client.LPush("weibo_message", strID)
	user := "user_" + author + "_weibo"
	client.LPush(user, strID)
	client.Incr("globalID")

	keyv := "weibo_" + origin
	client.HIncrBy(keyv, "resent", 1)

	client.Close()

	b, _ := json.Marshal(JsonResponse{0, "Succeeded"})
	io.WriteString(w, string(b))
}
func alreadSupport(weiboid int, login_user string, client *redis.Client) bool {

	key := fmt.Sprintf("weibo_%d_supports", weiboid)
	supports, _ := client.LRange(key, 0, -1)
	total := len(supports)
	sort.Strings(supports)
	has := sort.SearchStrings(supports, login_user)
	if has < total {
		if strings.EqualFold(login_user, supports[has]) {
			return true
		}
	}
	return false
}
func squareHandle(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("login_user")
	if len(author) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	key := "user_" + author + "_following"
	fansofwho, _ := client.SMembers(key)
	total := len(fansofwho)
	sort.Strings(fansofwho)
	//fmt.Println(fansofwho)

	client.SAdd("all_users", author) //new user inter weibo system

	weibos, _ := client.LRange("weibo_message", 0, 50)
	allweibo := make(ALL_WeiBO, 0, 50)
	for _, vv := range weibos {
		//ls, err := client.HGetAll("weibo_" + vv)
		ls, err := client.HMGet("weibo_"+vv, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin", "flag", "video", "redpacketid")
		if err != nil {
			continue
		}
		var weibo WeiBo
		for i, v := range ls {
			switch i {
			case 0:
				weibo.Weiboid, _ = strconv.Atoi(v)
				weibo.Support = alreadSupport(weibo.Weiboid, author, client)
			case 1:
				weibo.Msg = v
			case 2:
				weibo.Author = v
				if strings.EqualFold(weibo.Author, author) {
					weibo.Concern = true
				} else {
					has := sort.SearchStrings(fansofwho, weibo.Author)
					if has < total {
						if strings.EqualFold(weibo.Author, fansofwho[has]) {
							weibo.Concern = true
						}
					}
				}
			case 3:
				weibo.Creatime = v

			case 4:
				weibo.Supports, _ = strconv.Atoi(v)
			case 5:
				weibo.Resent, _ = strconv.Atoi(v)
			case 6:
				if len(v) >= 3 {
					temp := strings.Split(v, ",")
					weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
				}
			case 7:
				weibo.Comments, _ = strconv.Atoi(v)
			case 8:
				weibo.Origin = getWeibo(v, client)
			case 9:
				weibo.Flag = v
			case 10:
				if len(v) >= 1 {
					weibo.Video = getVideoinfo(v, client)
					weibo.Type = "video"
				} else {
					weibo.Type = "text"
				}
			case 11:
				if len(v) >= 1 {
					weibo.Redevpid = v
					weibo.Type = "redpacket"
				}
			}
		}
		if weibo.Type == "video" {
			if strings.Contains(weibo.Video.Url, "abcdefg") || len(weibo.Video.Url) < 5 {
				continue
			}
		}
		weibo.Userinfo = getUserinfo(weibo.Author, client, false)
		if weibo.Origin != nil {
			weibo.Type = weibo.Origin.Type
		}
		allweibo = append(allweibo, weibo)
	}
	//sort.Sort(allweibo)
	type MyResponse struct {
		JsonResponse
		Data ALL_WeiBO `json:"data"`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Data = allweibo

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))

	client.Close()
}

func filterHandle(w http.ResponseWriter, req *http.Request) {
	class := req.FormValue("class")
	author := req.FormValue("login_user")
	if len(author) < 1 || len(class) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	key := "user_" + author + "_following"
	fansofwho, _ := client.SMembers(key)
	total := len(fansofwho)
	sort.Strings(fansofwho)
	//fmt.Println(fansofwho)

	client.SAdd("all_users", author) //new user inter weibo system

	weibos, _ := client.LRange("weibo_message", 0, 50)
	allweibo := make(ALL_WeiBO, 0, 100)
	for _, vv := range weibos {
		//ls, err := client.HGetAll("weibo_" + vv)
		ls, err := client.HMGet("weibo_"+vv, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin", "flag", "video", "redpacketid")
		if err != nil {
			continue
		}
		if class != ls[9] {
			continue
		}
		var weibo WeiBo
		for i, v := range ls {
			switch i {
			case 0:
				weibo.Weiboid, _ = strconv.Atoi(v)
				weibo.Support = alreadSupport(weibo.Weiboid, author, client)
			case 1:
				weibo.Msg = v
			case 2:
				weibo.Author = v
				if strings.EqualFold(weibo.Author, author) {
					weibo.Concern = true
				} else {
					has := sort.SearchStrings(fansofwho, weibo.Author)
					if has < total {
						if strings.EqualFold(weibo.Author, fansofwho[has]) {
							weibo.Concern = true
						}
					}
				}
			case 3:
				weibo.Creatime = v

			case 4:
				weibo.Supports, _ = strconv.Atoi(v)
			case 5:
				weibo.Resent, _ = strconv.Atoi(v)
			case 6:
				if len(v) >= 3 {
					temp := strings.Split(v, ",")
					weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
				}
			case 7:
				weibo.Comments, _ = strconv.Atoi(v)
			case 8:
				weibo.Origin = getWeibo(v, client)
			case 9:
				weibo.Flag = v
			case 10:
				if len(v) >= 1 {
					weibo.Video = getVideoinfo(v, client)
					weibo.Type = "video"
				} else {
					weibo.Type = "text"
				}
			case 11:
				if len(v) >= 1 {
					weibo.Redevpid = v
					weibo.Type = "redpacket"
				}
			}
		}
		if weibo.Type == "video" {
			if strings.Contains(weibo.Video.Url, "abcdefg") || len(weibo.Video.Url) < 5 {
				continue
			}
		}
		weibo.Userinfo = getUserinfo(weibo.Author, client, false)
		allweibo = append(allweibo, weibo)
	}
	//sort.Sort(allweibo)
	type MyResponse struct {
		JsonResponse
		Data ALL_WeiBO `json:"data"`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Data = allweibo

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))

	client.Close()
}

func flagHandle(w http.ResponseWriter, req *http.Request) {
	weiboid := req.FormValue("weiboid")
	class := req.FormValue("class")
	if len(weiboid) < 1 || len(class) < 1 {
		jsonres := JsonResponse{1, "argument error"}
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

	key := "weibo_" + weiboid
	client.HSet(key, "flag", class)

	jsonres := JsonResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
	client.Close()
}

func classHandle(w http.ResponseWriter, req *http.Request) {
	var classname = [...]string{"红包", "商家", "政治", "军事", "财经", "社会", "文学", "名人", "电影", "旅游"}
	type MyResponse struct {
		JsonResponse
		Class []string `json:"data"`
		Total int      `json:"total"`
	}
	jsonres := MyResponse{}

	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Class = classname[:]
	jsonres.Total = len(classname)

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}

func main() {

	logfile, _ := os.OpenFile("./weibo.log", os.O_RDWR|os.O_CREATE, 0)
	logger = log.New(logfile, "\n", log.Ldate|log.Ltime|log.Lshortfile)

	clients.connFn = newcon
	Channel = make(chan string, 100)
	go Check_thread()

	var ok bool
	var client *redis.Client
	client, ok = clients.Get()
	if ok != true {
		return
	}

	ID, err := client.Get("globalID")
	if err != nil {
		client.Set("globalID", 1)
	}
	fmt.Println(ID)
	client.Close()

	http.HandleFunc("/support", supportHandle)
	http.HandleFunc("/unsupport", unsupportHandle)
	http.HandleFunc("/checksupport", checksupportHandle)
	http.HandleFunc("/write", writeHandle)
	http.HandleFunc("/writev2", writev2Handle)
	http.HandleFunc("/writev3", writev3Handle)
	http.HandleFunc("/writev4", writev4Handle)
	http.HandleFunc("/comment", commentHandle)
	http.HandleFunc("/checkcomment", checkcommentHandle)
	http.HandleFunc("/supportcomment", supportcommentHandle)
	http.HandleFunc("/upload", uploadHandle)
	http.HandleFunc("/concern", concernHandle)
	http.HandleFunc("/cancelconcern", cancelconcernHandle)
	http.HandleFunc("/check", checkHandle)
	http.HandleFunc("/checkmy", checkmyHandle)
	http.HandleFunc("/profile", profileHandle)
	http.HandleFunc("/portrait", portraitHandle)
	http.HandleFunc("/forward", forwardHandle)
	http.HandleFunc("/userinfo", userInfo)
	http.HandleFunc("/square", squareHandle)
	http.HandleFunc("/squarefilter", filterHandle)
	http.HandleFunc("/delete", deleteHandle)
	http.HandleFunc("/flag", flagHandle)
	http.HandleFunc("/queryclass", classHandle)
	http.HandleFunc("/test", testHandle)

	http.Handle("/", http.FileServer(http.Dir("./upload")))

	if err := http.ListenAndServe(":8888", nil); err != nil {
	}
}
