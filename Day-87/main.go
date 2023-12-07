package main

import "github.com/beego/beego/v2/server/web"

type HelloController struct {
    web.Controller
}

type UserController struct {
	web.Controller
}

func (c *HelloController) Get() {
    c.Data["json"] = map[string]string{"message": "Hello World"}
    c.ServeJSON()
}

func (u *UserController) HelloWorld()  {
	u.Ctx.WriteString("hello, world")
}

func (u UserController) GetUserById() {
    id := u.Ctx.Input.Param(":id")
	u.Ctx.WriteString("GetUserById: "+id)
}

func main() {
	// 設置端口為 8000
    web.BConfig.Listen.HTTPPort = 8000
	
	web.Router("/", &HelloController{})
    web.AutoRouter(&UserController{})
    web.CtrlGet("api/user/:id", UserController.GetUserById)

	web.Run()
}