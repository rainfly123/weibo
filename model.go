package main

type WeiBo struct {
  msg string  `json:"Msg"`
  author string  `json:"Author"`
  creatime string `json:"Creatime"`
  supports string `json:Supports`
  resent  string `json:Resent`
  pictures []string `json:Pictures`
}
