package v1

import (
	"fmt"
	models2 "gin-go-bl/coveralls/models"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddArticle(c *gin.Context) {
	user, _ := c.Get("user")
	id := user.(models2.User).ID
	var data models2.Article
	data.UserId = id
	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	code := models2.CreatArticle(&data)
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetAllArticle(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	arts, code, atotal := models2.GetAllArticle(size, num)
	//cate, _ := models.GetCategory(100, 0)
	c.JSON(200, gin.H{
		"arts": arts,
		//"cate":    cate,
		"size":    atotal,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetAllArticles(c *gin.Context) {

	arts, code, atotal := models2.GetAllArticles()
	//cate, _ := models.GetCategory(100, 0)
	c.JSON(200, gin.H{
		"arts": arts,
		//"cate":    cate,
		"size":    atotal,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetUserArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	data, code, total := models2.GetUserArticle(id, size, num)
	c.JSON(200, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetMeArticle(c *gin.Context) {
	data, _ := c.Get("user")
	id := data.(models2.User).ID
	//personValue := reflect.ValueOf(data)
	//ID := personValue.FieldByName("ID").Uint()
	//id := middleware.Userid
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	data, code, total := models2.GetUserArticle(int(id), size, num)
	c.JSON(200, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetCaArticle(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("size"))
	num, _ := strconv.Atoi(c.Param("num"))
	id, _ := strconv.Atoi(c.Param("id"))
	if size == 0 {
		size = -1
	}
	if num == 0 {
		num = -1
	}

	data, code, total := models2.GetCatArticle(id, size, num)
	c.JSON(200, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := models2.GetArticle(id)
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func EditArticle(c *gin.Context) {
	var data models2.Article
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	code := models2.EditArticle(id, &data)
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models2.DeleteArticle(id)
	c.JSON(200, gin.H{
		"status":  data,
		"message": utils.GetErrMsg(data),
	})
}
