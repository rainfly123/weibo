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
	Code    int    `json:Code`
	Message string `json:Message`
}
type WeiBo struct {
	Weiboid  int      `json:"Weiboid"`
	Msg      string   `json:"Msg"`
	Author   string   `json:"Author"`
	Creatime string   `json:"Creatime"`
	Supports int      `json:Supports`
	Resent   int      `json:Resent`
	Pictures []string `json:Pictures`
	Comments int      `json:Comments`
	Origin   *WeiBo   `json:Origin`
	Userinfo User     `json:User`
}

type User struct {
	Userid    string `json:"Userid"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"Gender"`
	Location  string `json:"Location"`
	Signature string `json:"Signature"`
	Portrait  string `json:Portrait`
	Follower  []User `json:Follower`
	Following []User `json:Following`
	Recommend []User `json:Recommend`
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
	defer fW.Close()
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("copy file error")
		return ""
	}
	return (ACCESS_URL + temp)
}

func writev2Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		io.WriteString(w, fmt.Sprintf("<html><head><title>我的第一个页面</title></head><body><form action='writev2?author=%s&msg=%s' method=\"post\" enctype=\"multipart/form-data\"><label>上传图片</label><input type=\"file\" name='file0'/><br/><input type=\"file\" name='file1'/><br/><input type=\"file\" name='file2'/><br/><<input type=\"file\" name='file3'/><br/><<input type=\"file\" name='file4'/><br/><<input type=\"file\" name='file5'/><br/><<input type=\"file\" name='file6'/><br/><<input type=\"file\" name='file7'/><br/><<input type=\"file\" name='file8'/><br/><<label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>", req.FormValue("author"), req.FormValue("msg")))
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
			Weiboid  string `json:Weiboid`
			Pictures string `json:Pictures`
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
		io.WriteString(w, "error!\n")
	}

	keyc := "weibo_" + weiboid + "_comments"
	value := author + ":" + comment
	client.RPush(keyc, value)
	keyv := "weibo_" + weiboid
	client.HIncrBy(keyv, "comments", 1)
	client.Close()
	jsonres := JsonResponse{0, "Succeeded"}
	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}
func checkcommentHandle(w http.ResponseWriter, req *http.Request) {

	type Comment struct {
		Author  User   `json:Author`
		Comment string `json:Comment`
	}

	weiboid := req.FormValue("weiboid")
	if len(weiboid) < 1 {
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
	for _, v := range comments {
		var comment Comment
		temp := strings.Split(v, ":")
		comment.Author = getUserinfo(temp[0], client, false)
		comment.Comment = temp[1]
		all = append(all, comment)
	}

	b, _ := json.Marshal(all)
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
	b, _ := json.Marshal(ALL_USERS)
	io.WriteString(w, string(b))
	client.Close()
}

func concernHandle(w http.ResponseWriter, req *http.Request) {
	login_user := req.FormValue("login_user")
	concern := req.FormValue("concern")
	if len(login_user) < 1 || len(concern) < 1 {
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
	if is != true {
		io.WriteString(w, "error!\n")
	}

	key = "user_" + cancel + "_follower"
	im, _ := client.SMove(key, "__trash__", login_user)
	if im != true {
		io.WriteString(w, "error!\n")
	}

	client.Close()
	jsonres := JsonResponse{0, "Succeeded"}
	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))
}

func getWeibo(weiboid string, client *redis.Client) *WeiBo {
	if len(weiboid) <= 1 {
		return nil
	}
	//	ls, err := client.HGetAll("weibo_" + weiboid)
	ls, err := client.HMGet("weibo_"+weiboid, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin")
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
			temp := strings.Split(v, ",")
			weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
		case 7:
			weibo.Comments, _ = strconv.Atoi(v)
		}
	}
	weibo.Userinfo = getUserinfo(weibo.Author, client, false)
	return &weibo
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

	for _, v := range users {
		k := "user_" + v + "_weibo"
		weibos, _ := client.LRange(k, 0, -1)
		if len(weibos) >= 1 {
			all = append(all, weibos[:len(weibos)]...)
		}
	}

	fmt.Println("weibo IDS ", all) //all is weiboid list
	allweibo := make(ALL_WeiBO, 0, 5000)
	for _, vv := range all {
		//ls, err := client.HGetAll("weibo_" + vv)
		ls, err := client.HMGet("weibo_"+vv, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin")
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
				temp := strings.Split(v, ",")
				weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
			case 7:
				weibo.Comments, _ = strconv.Atoi(v)
			case 8:
				weibo.Origin = getWeibo(v, client)
			}
		}
		weibo.Userinfo = getUserinfo(weibo.Author, client, false)
		allweibo = append(allweibo, weibo)
	}
	sort.Sort(allweibo)

	type MyResponse struct {
		JsonResponse
		Data ALL_WeiBO `json:Data`
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
		ls, err := client.HMGet("weibo_"+vv, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin")
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
				temp := strings.Split(v, ",")
				weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
			case 7:
				weibo.Comments, _ = strconv.Atoi(v)
			case 8:
				weibo.Origin = getWeibo(v, client)

			}
		}
		weibo.Userinfo = getUserinfo(weibo.Author, client, false)
		allweibo = append(allweibo, weibo)
	}
	//sort.Sort(allweibo)

	type MyResponse struct {
		JsonResponse
		Data ALL_WeiBO `json:Data`
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

	if len(login_user) < 1 || len(nickname) < 1 || len(gender) < 1 || len(location) < 1 || len(signature) < 1 {
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
	}

	portrait := "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg"
	key := "user_" + login_user + "_profile"
	client.HMSet(key, "nickname", nickname, "gender", gender, "location", location, "signature", signature, "portrait", portrait)

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
		if err != nil {
			continue
		}
		users = append(users, ls)
	}
	return users
}

func getUserinfo(userid string, client *redis.Client, detail bool) User {

	var user User
	key := "user_" + userid + "_profile"
	ls, err := client.HGetAll(key)
	if err != nil {
		return user
	}
	user.Userid = userid
	for i, v := range ls {
		switch i {
		case 1:
			user.Nickname = v
		case 3:
			user.Gender = v
		case 5:
			user.Location = v
		case 7:
			user.Signature = v
		case 9:
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

	if len(userid) < 1 {
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

	b, _ := json.Marshal(getUserinfo(userid, client, true))
	io.WriteString(w, string(b))
	client.Close()
}

func forwardHandle(w http.ResponseWriter, req *http.Request) {
	author := req.FormValue("author")
	msg := req.FormValue("msg")
	origin := req.FormValue("origin")
	if len(author) < 1 || len(msg) < 3 || len(origin) < 3 {
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

	client.SAdd("all_users", author) //new user inter weibo system

	weibos, _ := client.LRange("weibo_message", 0, 50)
	allweibo := make(ALL_WeiBO, 0, 50)
	for _, vv := range weibos {
		//ls, err := client.HGetAll("weibo_" + vv)
		ls, err := client.HMGet("weibo_"+vv, "weiboid", "msg", "author", "creatime", "supports", "resent", "pictures", "comments", "origin")
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
				temp := strings.Split(v, ",")
				weibo.Pictures = append(weibo.Pictures, temp[:len(temp)]...)
			case 7:
				weibo.Comments, _ = strconv.Atoi(v)
			case 8:
				weibo.Origin = getWeibo(v, client)

			}
		}
		weibo.Userinfo = getUserinfo(weibo.Author, client, false)
		allweibo = append(allweibo, weibo)
	}
	//sort.Sort(allweibo)
	type MyResponse struct {
		JsonResponse
		Data ALL_WeiBO `json:Data`
	}
	jsonres := MyResponse{}
	jsonres.Code = 0
	jsonres.Message = "Succeeded"
	jsonres.Data = allweibo

	b, _ := json.Marshal(jsonres)
	io.WriteString(w, string(b))

	client.Close()
}

func main() {

	logfile, _ := os.OpenFile("./weibo.log", os.O_RDWR|os.O_CREATE, 0)
	logger = log.New(logfile, "\n", log.Ldate|log.Ltime|log.Lshortfile)

	clients.connFn = newcon

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
	http.HandleFunc("/checksupport", checksupportHandle)
	http.HandleFunc("/write", writeHandle)
	http.HandleFunc("/writev2", writev2Handle)
	http.HandleFunc("/comment", commentHandle)
	http.HandleFunc("/checkcomment", checkcommentHandle)
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

	http.HandleFunc("/test", testHandle)

	http.Handle("/", http.FileServer(http.Dir("./upload/")))

	if err := http.ListenAndServe(":8888", nil); err != nil {
	}
}
