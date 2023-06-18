package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NameRequest struct {
	Name   string `form:"name"`
	Status bool   `form:"status"`
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")

	server.GET("/", test)

	server.Run()
}

func test(c *gin.Context) {
	type IndexData struct {
		Title   string
		Content string
		Name    string
	}
	var name NameRequest

	c.Bind(&name)

	fmt.Print(name)

	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一支 gin 專案"
	data.Name = name.Name

	c.HTML(http.StatusOK, "index.html", data)
}
