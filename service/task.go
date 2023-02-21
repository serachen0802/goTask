package service

import (
	"fmt"
	taskModel "goTask/model"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Alldata map[int]taskModel.Task

func init() {
	Alldata = map[int]taskModel.Task{}
}

// 取得全部Task
func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, Alldata)
}

//新增Task
func Task(c *gin.Context) {
	name := c.PostForm("name")

	allIds := []int{}
	for key := range Alldata {
		allIds = append(allIds, key)
	}

	sort.Ints(allIds)
	maxid := 1
	if len(allIds) != 0 {
		maxid = allIds[len(allIds)-1] + 1
	}

	Alldata[maxid] = taskModel.Task{
		ID:     maxid,
		Name:   name,
		Status: false,
	}
	fmt.Println("alldata", Alldata)
	c.JSON(http.StatusOK, Alldata[maxid])

	// for key := range Alldata {
	// 	fmt.Println("key", key)
	// }
	// num := len(Alldata)
	// if num > 0 {
	// 	num = Alldata[num-1].ID + 1
	// } else {
	// 	num = 1
	// }
	// stashTask := taskModel.Task{
	// 	ID:     num, //取到現有最大ID
	// 	Name:   name,
	// 	Status: false,
	// }

	// c.JSON(http.StatusOK, stashTask)
}

//更新Task狀態
func UpdateTask(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	// 看這個ID存在嗎
	name := c.PostForm("name")
	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		return
	}

	_, ok := Alldata[id]
	if !ok {
		c.JSON(http.StatusOK, "ID不存在")
		return
	}

	Alldata[id] = taskModel.Task{
		ID:     id,
		Name:   name,
		Status: status,
	}
	// 	for i := 0; i < len(Alldata); i++ {
	// 		if Alldata[i].ID == id {
	// 			Alldata[i].Name = name
	// 			Alldata[i].Status = status
	// 		}
	// 	}
	c.JSON(http.StatusOK, Alldata[id])
}

//刪除Task
func Delete(c *gin.Context) {
	id := c.Param("id")

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"id": id,
	// 	})
	IdtoInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
		return
	}
	delete(Alldata, IdtoInt)
	c.Status(http.StatusOK)
	// 	for i := 0; i < len(Alldata); i++ {
	// 		if Alldata[i].ID == IdtoInt {
	// 			if i == 0 {
	// 				Alldata = Alldata[1:]
	// 			} else {
	// 				Alldata = append(Alldata[:i], Alldata[i+1:]...)
	// 			}
	// 		}
	// 	}

}
