package main

import (
	"fmt"
	"net/http"
	"net/url"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

)

type SearchResult struct {
	Users []string 
	Weibos []string 
}
func getusers(key string ) []string{
	db, err := sql.Open("mysql", "root:123321@/weibo")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()
 
        words := fmt.Sprintf("select userid from user where name like '%%%s%%'", key)
	rows, err := db.Query(words)
	if err != nil {
		log.Println(err)
	}
 
	defer rows.Close()
	var id string 
	var users []string 
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
                users = append(users, id)
	}
 
       return users
}

func getweibos(key string) []string{
	db, err := sql.Open("mysql", "root:123321@/weibo")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()
 
        words := fmt.Sprintf("select weiboid from weibo where msg like '%%%s%%'", key)
	rows, err := db.Query(words)
	if err != nil {
		log.Println(err)
	}
 
	defer rows.Close()
	var id string 
	var weibos[]string 
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
                weibos = append(weibos, id)
	}
       return weibos
}

func search(key string) SearchResult {
	var result SearchResult
        result.Users = getusers(key)
        result.Weibos = getweibos(key)
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
    //syncweibo("3230")
    fmt.Println(search("毒"))
}*/
