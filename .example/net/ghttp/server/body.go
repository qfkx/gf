package main

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"io/ioutil"
)

func main() {
	s := g.Server()
	s.SetIndexFolder(true)
	s.BindHandler("/", func(r *ghttp.Request) {
		body1 := r.GetBody()
		body2, _ := ioutil.ReadAll(r.Body)
		fmt.Println(body1)
		fmt.Println(body2)
		r.Response.Write("hello world")
	})
	s.SetPort(8999)
	s.Run()
}
