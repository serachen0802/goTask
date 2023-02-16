package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type data struct {
	ID     int
	Name   string
	Status bool
}

var Alldata []data

func main() {

	// http.HandleFunc("/", sayhi)

	// err := http.ListenAndServe(":9090", nil)
	// if err != nil {
	// 	log.Fatal("ListenAnndAerve:", err)
	// }
	server := set()
	server.Run(":9090")
}

func set() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", sayhi)
	router.GET("/tasks", getTask)
	router.POST("/task", task)
	router.PUT("/task/:id", updateTask)
	router.DELETE("/task/:id", delete)

	return router
}
func sayhi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hi"})
}

//取得全部Task
func getTask(c *gin.Context) {
	c.JSON(http.StatusOK, Alldata)
}

//新增Task
func task(c *gin.Context) {
	// input := context.Param("input")
	name := c.PostForm("name")
	num := len(Alldata)
	if num > 0 {
		num = Alldata[num-1].ID + 1
	} else {
		num = 1
	}
	stashTask := data{
		ID:     num, //取到現有最大ID
		Name:   name,
		Status: false,
	}

	Alldata = append(Alldata, stashTask)
	c.JSON(http.StatusOK, stashTask)
}

//更新Task狀態
func updateTask(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	name := c.PostForm("name")
	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		return
	}

	for i := 0; i < len(Alldata); i++ {
		if Alldata[i].ID == id {
			Alldata[i].Name = name
			Alldata[i].Status = status
		}
	}
	c.JSON(http.StatusOK, data{
		ID:     id,
		Name:   name,
		Status: status,
	})
}

//刪除Task
func delete(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	IdtoInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	for i := 0; i < len(Alldata); i++ {
		if Alldata[i].ID == IdtoInt {
			if i == 0 {
				Alldata = Alldata[1:]
			} else {
				Alldata = append(Alldata[:i], Alldata[i+1:]...)
			}
		}
	}
	c.JSON(http.StatusOK, "")
}
