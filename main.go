package main
import (
	"io"
        "menteslibres.net/gosexy/redis"
        "log"
        "os"
        "strconv"
        "fmt"
	"net/http"
        "time"
        "sort"
        "strings"
        "encoding/json"
)

type WeiBo struct {
  Weiboid int `json:"Weiboid"`
  Msg string  `json:"Msg"`
  Author string  `json:"Author"`
  Creatime string `json:"Creatime"`
  Supports  int `json:Supports`
  Resent   int `json:Resent`
  Pictures []string `json:Pictures`
  Comments  int `json:Comments`
  Origin *WeiBo `json:Origin`
  Userinfo User `json:User`
}

type User struct {
  Userid string `json:"Userid"`
  Nickname string `json:"nickname"`
  Gender string  `json:"Gender"`
  Location string  `json:"Location"`
  Signature string `json:"Signature"`
  Portrait string `json:Portrait`
  Follower []User `json:Follower`
  Following []User `json:Following`
  Recommend []User`json:Recommend`
}

type  ALL_WeiBO []WeiBo
func (list ALL_WeiBO) Len() int {  
    return len(list)  
}  
func (list ALL_WeiBO) Less(i, j int) bool {  
    return list[i].Creatime > list[j].Creatime 
}  
func (list ALL_WeiBO) Swap(i, j int) {  
    list[i], list[j] = list[j], list[i]
}  

var logger  *log.Logger
var host = "127.0.0.1"
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

    strID, _:= client.Get("globalID")
    key := "weibo_" + strID
    now := time.Now().Format("2006-01-02 15:04:05")
    client.HMSet(key, "weiboid", strID, "msg", msg, "author", author, "creatime", now, "supports", 0, "resent", 0,
                "pictures", pic, "comments", 0)
    client.LPush("weibo_message", strID)
    user := "user_" + author + "_weibo"
    client.LPush(user, strID)
    client.Incr("globalID")

    fmt.Fprintf(w,"%s",key)
    io.WriteString(w, "ok!\n")

    client.Close() 
}

func commentHandle(w http.ResponseWriter, req *http.Request) {
    author := req.FormValue("login_user")
    if len(author) < 1 {
    }
    comment := req.FormValue("comment")
    if len(comment) < 3 {
    }
    weiboid := req.FormValue("weiboid")
    if len(weiboid) < 3 {
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
    io.WriteString(w, "ok!\n")
    client.Close() 
}
func checkcommentHandle(w http.ResponseWriter, req *http.Request) {

    type Comment struct {
        Author string `json:Author`
        Comment string `json:Comment`
    }

    weiboid := req.FormValue("weiboid")
    if len(weiboid) < 3 {
    }

    all :=  make([]Comment, 0, 1000)
    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    key := "weibo_" + weiboid + "_comments"
    comments , _ := client.LRange(key, 0, -1)
    for _, v := range comments {
        var comment Comment
        temp := strings.Split(v, ":")
        comment.Author = temp[0]
        comment.Comment = temp[1]
        all = append(all, comment)
    } 
   
    b, _ := json.Marshal(all)
    io.WriteString(w, string(b))
    client.Close() 
}

func supportHandle(w http.ResponseWriter, req *http.Request) {
    author := req.FormValue("login_user")
    if len(author) < 1 {
    }
    weiboid := req.FormValue("weiboid")
    if len(weiboid) < 3 {
    }

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    keys := "weibo_" + weiboid + "_supports"
    client.RPush(keys, author)

    keyv := "weibo_" + weiboid 
    client.HIncrBy(keyv, "supports", 1)
    io.WriteString(w, "ok!\n")
    client.Close() 
}
func checksupportHandle(w http.ResponseWriter, req *http.Request) {

    weiboid := req.FormValue("weiboid")
    if len(weiboid) < 3 {
    }

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    key := "weibo_" + weiboid + "_supports"
    supports , _ := client.LRange(key, 0, -1)
    b, _ := json.Marshal(supports)
    io.WriteString(w, string(b))
    client.Close() 
}

func concernHandle(w http.ResponseWriter, req *http.Request) {
    login_user := req.FormValue("login_user")
    if len(login_user) < 1 {
    }
    concern := req.FormValue("concern")
    if len(concern) < 1 {
    }

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    key := "user_" + login_user + "_following"
    client.SAdd(key, concern)

    key = "user_" + concern + "_follower"
    client.SAdd(key, login_user)

    io.WriteString(w, "ok!\n")

    client.Close() 
}

func cancelconcernHandle(w http.ResponseWriter, req *http.Request) {
    login_user := req.FormValue("login_user")
    if len(login_user) < 1 {
    }
    cancel := req.FormValue("cancel")
    if len(cancel) < 1 {
    }

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
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

    io.WriteString(w, "ok!\n")

    client.Close() 
}

func getWeibo(weiboid string, client *redis.Client) *WeiBo{
    ls, err := client.HGetAll("weibo_" + weiboid)
        if err != nil {
           return nil 
        }
        var weibo WeiBo
        for i, v := range(ls) {
            switch i {
                case 1:
                     weibo.Weiboid, _ = strconv.Atoi(v)
                case 3:
                     weibo.Msg = v
                case 5:
                     weibo.Author = v
                case 7:
                     weibo.Creatime= v
                case 9:
                     weibo.Supports, _ = strconv.Atoi(v)
                case 11:
                     weibo.Resent, _ = strconv.Atoi(v)
                case 13:
                     temp := strings.Split(v, ",")
                     weibo.Pictures  = append(weibo.Pictures, temp[:len(temp)]...)
                case 15:
                     weibo.Comments, _ = strconv.Atoi(v)
            }
        }
        weibo.Userinfo = getUserinfo(weibo.Author, client, false)
      return &weibo
}

func checkHandle(w http.ResponseWriter, req *http.Request) {
    login_user := req.FormValue("login_user")
    if len(login_user) < 1 {
    }

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    all := make([]string, 0 , 200)
    key := "user_" + login_user + "_following"
    users, _ := client.SMembers(key)

    for _, v := range users {
       k := "user_" + v + "_weibo"
       weibos, _ := client.LRange(k, 0, -1)
       if len(weibos) >= 1 {
           all = append(all, weibos[:len(weibos)]...)
       }
    }

    fmt.Println("weibo IDS ", all)  //all is weiboid list
    allweibo := make(ALL_WeiBO, 0 , 5000)
    for _, vv := range all {
        ls, err := client.HGetAll("weibo_" + vv)
        if err != nil {
            continue
        }
        var weibo WeiBo
        for i, v := range(ls) {
            switch i {
                case 1:
                     weibo.Weiboid, _ = strconv.Atoi(v)
                case 3:
                     weibo.Msg = v
                case 5:
                     weibo.Author = v
                case 7:
                     weibo.Creatime= v
                case 9:
                     weibo.Supports, _ = strconv.Atoi(v)
                case 11:
                     weibo.Resent, _ = strconv.Atoi(v)
                case 13:
                     temp := strings.Split(v, ",")
                     weibo.Pictures  = append(weibo.Pictures, temp[:len(temp)]...)
                case 15:
                     weibo.Comments, _ = strconv.Atoi(v)
                case 17:
                     weibo.Origin = getWeibo(v, client)
            }
        }
        weibo.Userinfo = getUserinfo(weibo.Author, client, false)
        allweibo = append(allweibo, weibo)
    }
    sort.Sort(allweibo)
    b, _ := json.Marshal(allweibo)
    io.WriteString(w, string(b))
    client.Close() 
}

func checkmyHandle(w http.ResponseWriter, req *http.Request) {
    login_user := req.FormValue("login_user")
    if len(login_user) < 1 {
    }
    key := "user_" + login_user + "_weibo"

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    weibos, _ := client.LRange(key, 0, -1)

    allweibo := make(ALL_WeiBO, 0 , 500)
    for _, vv := range weibos{
        ls, err := client.HGetAll("weibo_" + vv)
        if err != nil {
            continue
        }
        var weibo WeiBo
        for i, v := range(ls) {
            switch i {
                case 1:
                     weibo.Weiboid, _ = strconv.Atoi(v)
                case 3:
                     weibo.Msg = v
                case 5:
                     weibo.Author = v
                case 7:
                     weibo.Creatime= v

                case 9:
                     weibo.Supports, _ = strconv.Atoi(v)
                case 11:
                     weibo.Resent, _ = strconv.Atoi(v)
                case 13:
                     temp := strings.Split(v, ",")
                     weibo.Pictures  = append(weibo.Pictures, temp[:len(temp)]...)
                case 15:
                     weibo.Comments, _ = strconv.Atoi(v)
                case 17:
                     weibo.Origin = getWeibo(v, client)

            }
        }
        weibo.Userinfo = getUserinfo(weibo.Author, client, false)
        allweibo = append(allweibo, weibo)
    }
    //sort.Sort(allweibo)
    b, _ := json.Marshal(allweibo)
    io.WriteString(w, string(b))
    client.Close() 
}

func profileHandle(w http.ResponseWriter, req *http.Request) {
    login_user := req.FormValue("login_user")
    nickname := req.FormValue("nickname")
    gender := req.FormValue("gender")
    location := req.FormValue("location")
    signature := req.FormValue("signature")

    if len(login_user) < 1 || len(nickname) < 1 || len(gender) < 1 || len(location) < 1 || len(signature) < 1 {
        io.WriteString(w, "error")
        return
    }

    

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    portrait := "http://7xvsyw.com1.z0.glb.clouddn.com/a.jpeg"
    key := "user_" + login_user + "_profile"
    client.HMSet(key, "nickname", nickname, "gender", gender, "location", location, "signature", signature, "portrait", portrait)

    io.WriteString(w, "ok")
    client.Close() 
}

func recommend(client *redis.Client) []string {

    var users []string
    weibos, _ := client.LRange("weibo_message", 0, 100)
    for _, v := range weibos{
        key := "weibo_" + v
        ls, err := client.HGet(key, "author")
        if err != nil {
            continue
        }
        users = append(users, ls)
   }
   return users
}

func getUserinfo(userid string, client *redis.Client, detail bool ) User {

    var user User
    key := "user_" + userid + "_profile"
    ls, err := client.HGetAll(key)
    if err != nil {
       return  user
    }
    user.Userid = userid
    for i, v := range(ls) {
        switch i {
            case 1:
            user.Nickname  = v
            case 3:
            user.Gender  = v
            case 5:
            user.Location  = v
            case 7:
            user.Signature  = v
            case 9:
            user.Portrait  = v
        }
    }
    if detail {
        var temp User
        keyer := "user_" + userid + "_follower" 
        follower , _ := client.SMembers(keyer)
        for _, n := range(follower) {
            temp = getUserinfo(n, client, false)
            user.Follower = append(user.Follower, temp)
        }

        keying := "user_" + userid + "_following" 
        following , _ := client.SMembers(keying)
        for _, k := range(following) {
            temp = getUserinfo(k, client, false)
           user.Following = append(user.Following, temp)
        }

       // s, _ := client.SRandMember("all_users", 10)
        var recuser []string
        s := recommend(client)
        for _, us := range(s) {
            if has(following, us) || strings.EqualFold(us, userid) {
                continue
            }else {
                if !has(recuser, us) {
                    recuser = append(recuser, us)
                }
            }
        }
        
        for _, p := range(recuser) {
           temp = getUserinfo(p, client, false)
           user.Recommend = append(user.Recommend, temp)
        }
    }
    return user
}

func has(sl []string, el string) bool{
   for _, i := range(sl) {
       if strings.EqualFold(i, el) {
           return true;
       }
   }
   return false;
}

func userInfo(w http.ResponseWriter, req *http.Request) {
    userid := req.FormValue("userid")

    if len(userid) < 1 {
        io.WriteString(w, "error")
        return
    }

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    b, _ := json.Marshal(getUserinfo(userid, client, true))
    io.WriteString(w, string(b))
    client.Close() 
}

func forwardHandle(w http.ResponseWriter, req *http.Request) {
    author := req.FormValue("author")
    if len(author) < 1 {
    }
    msg := req.FormValue("msg")
    if len(msg) < 3 {
    }
    origin := req.FormValue("origin")
    if len(origin) < 3 {
    }

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    strID, _:= client.Get("globalID")
    key := "weibo_" + strID
    now := time.Now().Format("2006-01-02 15:04:05")
    client.HMSet(key, "weiboid", strID, "msg", msg, "author", author, "creatime", now, "supports", 0, "resent", 0,
                "pictures", "", "comments", 0, "origin", origin)
    client.LPush("weibo_message", strID)
    user := "user_" + author + "_weibo"
    client.LPush(user, strID)
    client.Incr("globalID")

    fmt.Fprintf(w,"%s",key)
    io.WriteString(w, "ok!\n")

    keyv := "weibo_" + origin
    client.HIncrBy(keyv, "resent", 1)

    client.Close() 
}
func squareHandle(w http.ResponseWriter, req *http.Request) {
    author := req.FormValue("login_user")
    if len(author) < 1 {
    }

    var ok bool 
    var client *redis.Client
    client, ok = clients.Get()
    if ok != true {
       io.WriteString(w, "error!\n")
    }

    client.SAdd("all_users", author)  //new user inter weibo system

    weibos, _ := client.LRange("weibo_message", 0, 50)
    allweibo := make(ALL_WeiBO, 0 , 50)
    for _, vv := range weibos{
        ls, err := client.HGetAll("weibo_" + vv)
        if err != nil {
            continue
        }
        var weibo WeiBo
        for i, v := range(ls) {
            switch i {
                case 1:
                     weibo.Weiboid, _ = strconv.Atoi(v)
                case 3:
                     weibo.Msg = v
                case 5:
                     weibo.Author = v
                case 7:
                     weibo.Creatime= v

                case 9:
                     weibo.Supports, _ = strconv.Atoi(v)
                case 11:
                     weibo.Resent, _ = strconv.Atoi(v)
                case 13:
                     temp := strings.Split(v, ",")
                     weibo.Pictures  = append(weibo.Pictures, temp[:len(temp)]...)
                case 15:
                     weibo.Comments, _ = strconv.Atoi(v)
                case 17:
                     weibo.Origin = getWeibo(v, client)

            }
        }
        weibo.Userinfo = getUserinfo(weibo.Author, client, false)
        allweibo = append(allweibo, weibo)
    }
    //sort.Sort(allweibo)
    b, _ := json.Marshal(allweibo)
    io.WriteString(w, string(b))


    client.Close() 
}




func main() {
        logfile, _:= os.OpenFile("./weibo.log", os.O_RDWR|os.O_CREATE,0)
        logger = log.New(logfile,"\n", log.Ldate|log.Ltime|log.Lshortfile);

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
