package main

import (
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/client/orm"
    _ "github.com/go-sql-driver/mysql" // 導入 MySQL 驅動
)

type User struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

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
    var o orm.Ormer
    o = orm.NewOrm() // 创建一个 Ormer

    id := u.Ctx.Input.Param(":id")

    uid, err := strconv.Atoi(id) // 將字符串轉換為 int
    if err != nil {
        u.Ctx.WriteString("Invalid ID")
        return
    }

    user := &User{ID: uid}
    err = o.Read(user)

    if err != nil {
        u.Ctx.WriteString("Err: ")
    }

    u.Data["json"] = map[string]string{"id": strconv.Itoa(user.ID) ,"name": user.Name}
    u.ServeJSON()
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(User))

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
}


func main() {
    orm.RunSyncdb("default", false, true)
    o := orm.NewOrm()

    // data
	user := new(User)
	user.Name = "mike"

	// insert data
	o.Insert(user)
	
	web.Router("/", &HelloController{})
    web.AutoRouter(&UserController{})
    web.CtrlGet("api/user/:id", UserController.GetUserById)

	web.Run()
}