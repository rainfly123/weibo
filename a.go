package main

import (
    "fmt"
)

func main() {
   var s = []string {"a","b","c"}
   fmt.Println(len(s))
   m := append(s[:1], s[2:]...) 
   fmt.Println(m)

}
