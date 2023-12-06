package main

import "github.com/beego/beego/v2/server/web"

type HelloController struct {
    web.Controller
}

func (c *HelloController) Get() {
    c.Data["json"] = map[string]string{"message": "Hello World"}
    c.ServeJSON()
}

func main() {
	// 設置端口為 8000
    web.BConfig.Listen.HTTPPort = 8000
	
	web.Router("/", &HelloController{})
	web.Run()
}