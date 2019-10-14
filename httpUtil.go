package main

import (
	"github.com/astaxie/beego/httplib"
	"log"
	"time"
)

func main()  {
	req := httplib.Get("http://beego.me/")

	req.SetTimeout(100 * time.Second, 30 * time.Second)
	req.Header("Accept-Encoding","gzip,deflate,sdch")
	req.Header("Host","beego.me")

	str , err := req.String()
	if err != nil{
		log.Fatal(err)
	}

	log.Println(str)
}
