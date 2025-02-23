package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html", g.Map{
			"header":    "This is header",
			"container": "This is container",
			"footer":    "This is footer",
		})
	})
	s.SetPort(8199)
	s.Run()
}
