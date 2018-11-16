package note

import (
	"demo/models/note"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"time"
)
type record = note.Note
var Note = note.Note{}

func Insert(c *gin.Context)  {
	var note record
	err := c.BindJSON(&note)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "404"})
		fmt.Println(err.Error())
		return
	}
	note.Id = bson.NewObjectId()
	note.CreatedAt = time.Now()
	note.ModifiedAt = time.Now()
	err_ := Note.Insert(note)
	if err_ == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": note})
	} else {
		fmt.Println(err_.Error())
		c.JSON(http.StatusOK, gin.H{"code": 403, "data": note})
	}
}

func Get(c *gin.Context)  {
	id := c.Param("id")
	exist, err := Note.FindById(id)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": exist})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 404, "data": ""})
	}
}
func GetAll(c *gin.Context)  {
	exist, err := Note.FindAll()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": exist})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 404, "data": ""})
	}
}