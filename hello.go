package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int `form:"id"`
	CreatedAt string
	CreatedBy string `form:"createdby"`
	Content   string `form:"content"`
	Status    int    `form:"status"`
}

var todo []Todo
var idMax = 1

func Saiban() int {
	idMax = idMax + 1
	return idMax
}

func GetDataTodo(c *gin.Context) {
	var b Todo
	c.Bind(&b)
	b.ID = Saiban()
	b.Status = 0
	b.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	todo = append(todo, b)
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}
func GetDoneTodo(c *gin.Context) {
	var b Todo
	c.Bind(&b)

	if b.Status == 0 {
		b.Status = 1
	} else {
		b.Status = 0
	}
	for idx, t := range todo {
		if t.ID == b.ID {
			todo[idx].Status = b.Status
		}
	}
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todo": todo,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./tmpl/index.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todo": todo,
		})
	})
	r.GET("/yaru", GetDataTodo)
	r.GET("/done", GetDoneTodo)
	r.Run(":80") // listen and serve on 0.0.0.0:80 (for windows "localhost:80")
}
